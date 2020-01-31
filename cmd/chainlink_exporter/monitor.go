package main

import (
	"chainlink_exporter/abi"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/atomic"
	"go.uber.org/zap"
	"math/big"
	"sync"
	"time"
)

const (
	PRECISION = 1000000000
)

type (
	Monitor struct {
		client      *ethclient.Client
		aggregators map[common.Address]*AggregatorMonitor

		addr            common.Address
		fulfillmentAddr common.Address
		oracle          *abi.Oracle
		linkContract    *abi.ERC

		lock sync.Mutex

		lastResTime *atomic.Uint64
		lastReqTime *atomic.Uint64

		lastResGauge          prometheus.Gauge
		lastReqGauge          prometheus.Gauge
		currentHeightGauge    prometheus.Gauge
		balanceGauge          prometheus.Gauge
		linkBalanceGauge      *prometheus.GaugeVec
		responseTimeHistogram *prometheus.HistogramVec

		revenueCounter     *prometheus.CounterVec
		fulfillmentCounter *prometheus.CounterVec
		missCounter        *prometheus.CounterVec
	}
)

func NewMonitor(addr common.Address, fulfillmentAddr common.Address, linkAddr common.Address, client *ethclient.Client) (*Monitor, error) {
	m := &Monitor{
		addr:            addr,
		fulfillmentAddr: fulfillmentAddr,
		client:          client,
		aggregators:     map[common.Address]*AggregatorMonitor{},
		lock:            sync.Mutex{},
		lastResTime:     atomic.NewUint64(0),
		lastReqTime:     atomic.NewUint64(0),
		lastResGauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "cl",
			Subsystem: "mon",
			Name:      "last_response",
			Help:      "Height of the last response",
		}),
		lastReqGauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "cl",
			Subsystem: "mon",
			Name:      "last_request",
			Help:      "Height of the last oracle request",
		}),
		currentHeightGauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "cl",
			Subsystem: "mon",
			Name:      "height",
			Help:      "Last synced height",
		}),
		responseTimeHistogram: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "cl",
			Subsystem: "mon",
			Name:      "response_time",
			Help:      "Average response time in blocks",
			Buckets:   []float64{1, 2, 3, 4, 5, 10, 15},
		}, []string{"spec_id"}),
		fulfillmentCounter: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "cl",
			Subsystem: "mon",
			Name:      "fulfilled",
			Help:      "Number of successfully fulfilled requests",
		}, []string{"spec_id", "requester"}),
		missCounter: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "cl",
			Subsystem: "mon",
			Name:      "missed",
			Help:      "Number of missed requests",
		}, []string{"spec_id", "requester"}),
		revenueCounter: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "cl",
			Subsystem: "mon",
			Name:      "revenue",
			Help:      "Number of LINK tokens earned",
		}, []string{"spec_id", "requester", "status"}),
		balanceGauge: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "cl",
			Subsystem: "mon",
			Name:      "eth_balance",
			Help:      "Balance of the oracle account",
		}),
		linkBalanceGauge: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "cl",
			Subsystem: "mon",
			Name:      "link_balance",
			Help:      "Link balance of the oracle",
		}, []string{"type"}),
	}

	prometheus.MustRegister(m.lastResGauge)
	prometheus.MustRegister(m.lastReqGauge)
	prometheus.MustRegister(m.currentHeightGauge)
	prometheus.MustRegister(m.responseTimeHistogram)
	prometheus.MustRegister(m.fulfillmentCounter)
	prometheus.MustRegister(m.revenueCounter)
	prometheus.MustRegister(m.missCounter)
	prometheus.MustRegister(m.balanceGauge)
	prometheus.MustRegister(m.linkBalanceGauge)

	oracle, err := abi.NewOracle(addr, m.client)
	if err != nil {
		return nil, err
	}
	m.oracle = oracle

	linkContract, err := abi.NewERC(linkAddr, m.client)
	if err != nil {
		return nil, err
	}
	m.linkContract = linkContract

	_, err = oracle.Owner(nil)
	if err != nil {
		return nil, fmt.Errorf("address does not belong to an oracle")
	}

	return m, nil
}

func (m *Monitor) Start() {
	go m.headRoutine()
	go m.requestRoutine()
	go m.metricRoutine()
}

func (m *Monitor) metricRoutine() {
	zap.L().Info("Starting metric routine")
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			m.lastResGauge.Set(float64(m.lastResTime.Load()))
			m.lastReqGauge.Set(float64(m.lastReqTime.Load()))
		}
	}
}

func (m *Monitor) updateBalances() {
	zap.L().Debug("fetching balances")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	balance, err := m.client.BalanceAt(ctx, m.fulfillmentAddr, nil)
	if err != nil {
		zap.L().Error("failed to fetch oracle balance", zap.Error(err))
		return
	}

	balance.Div(balance, big.NewInt(params.Ether/PRECISION))
	m.balanceGauge.Set(float64(balance.Uint64()) / PRECISION)

	owner, err := m.oracle.Owner(&bind.CallOpts{
		Context: ctx,
	})
	withdrawableLinkBalance, err := m.oracle.Withdrawable(&bind.CallOpts{
		From:    owner,
		Context: ctx,
	})
	if err != nil {
		zap.L().Error("failed to fetch withdrawable LINK balance", zap.Error(err))
		return
	}
	withdrawableLinkBalance.Div(withdrawableLinkBalance, big.NewInt(params.Ether/PRECISION))
	m.linkBalanceGauge.WithLabelValues("withdrawable").Set(float64(withdrawableLinkBalance.Uint64()) / PRECISION)

	linkBalance, err := m.linkContract.BalanceOf(&bind.CallOpts{Context: ctx}, m.addr)
	if err != nil {
		zap.L().Error("failed to fetch LINK balance", zap.Error(err))
		return
	}
	linkBalance.Div(linkBalance, big.NewInt(params.Ether/PRECISION))
	m.linkBalanceGauge.WithLabelValues("balance").Set(float64(linkBalance.Uint64()) / PRECISION)

	zap.L().Debug("fetched balances")
}

func (m *Monitor) headRoutine() {
	for {
		zap.L().Info("Starting head routine")
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			headChan := make(chan *types.Header, 100)
			sub, err := m.client.SubscribeNewHead(ctx, headChan)
			if err != nil {
				zap.L().Error("failed to subscribe to new heads", zap.Error(err))
				return
			}

			for {
				select {
				case err = <-sub.Err():
					zap.L().Error("head subscription errored", zap.Error(err))
					return
				case header, has := <-headChan:
					if !has {
						zap.L().Error("head subscription closed", zap.Error(err))
						return
					}

					// Update balances
					go m.updateBalances()

					// Update metrics and update aggregator monitors
					m.currentHeightGauge.Set(float64(header.Number.Uint64()))
					for _, monitor := range m.aggregators {
						monitor.HandleNewBlock(header.Number.Uint64())
					}
				}
			}
		}()

		zap.L().Warn("head routine died. restarting in 5sec")
		time.Sleep(5 * time.Second)
	}
}

func (m *Monitor) requestRoutine() {
	for {
		zap.L().Info("Starting request routine")
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			reqChan := make(chan *abi.OracleOracleRequest, 100)
			sub, err := m.oracle.WatchOracleRequest(&bind.WatchOpts{
				Context: ctx,
			}, reqChan, nil)
			if err != nil {
				zap.L().Error("failed to watch oracle requests", zap.Error(err))
				return
			}

			for {
				select {
				case err = <-sub.Err():
					zap.L().Error("oracle requests subscription errored", zap.Error(err))
					return

				case req, has := <-reqChan:
					if !has {
						zap.L().Error("request subscription closed", zap.Error(err))
						return
					}
					err := m.handleRequest(req)
					if err != nil {
						zap.L().Warn("failed to handle request", zap.Error(err))
						continue
					}
				}
			}
		}()

		zap.L().Warn("request routine died. restarting in 5sec")
		time.Sleep(5 * time.Second)
	}
}

func (m *Monitor) handleRequest(req *abi.OracleOracleRequest) error {
	logger := zap.L().With(zap.Uint64("height", req.Raw.BlockNumber),
		zap.String("requester", req.Requester.String()), zap.Binary("request_id", req.RequestId[:]),
		zap.ByteString("spec_id", req.SpecId[:]))
	logger.Info("received request")

	if old := m.lastReqTime.Load(); old < req.Raw.BlockNumber {
		m.lastReqTime.CAS(old, req.Raw.BlockNumber)
	}

	if agg, contains := m.aggregators[req.Requester]; contains {
		agg.handleRequest(req)
		return nil
	}

	logger.Debug("requester unknown; creating a new aggregator monitor")

	m.lock.Lock()
	defer m.lock.Unlock()
	agg, err := abi.NewAggregator(req.Requester, m.client)
	if err != nil {
		return err
	}

	_, err = agg.Owner(nil)
	if err != nil {
		logger.Warn("requester is not an aggregator")
		return nil
	}

	am := NewAggregatorMonitor(agg, req.Requester, m)
	am.handleRequest(req)

	go am.Monitor()

	m.aggregators[req.Requester] = am

	return nil
}

func (m *Monitor) HandleFulfillment(res *abi.AggregatorChainlinkFulfilled, req *abi.OracleOracleRequest) {
	if old := m.lastResTime.Load(); old < res.Raw.BlockNumber {
		m.lastResTime.CAS(old, res.Raw.BlockNumber)
	}

	deltaBlocks := res.Raw.BlockNumber - req.Raw.BlockNumber
	m.responseTimeHistogram.WithLabelValues(string(req.SpecId[:])).Observe(float64(deltaBlocks))

	m.fulfillmentCounter.WithLabelValues(string(req.SpecId[:]), req.Requester.String()).Inc()
	m.revenueCounter.WithLabelValues(string(req.SpecId[:]), req.Requester.String(), "fulfilled").Add(float64(req.Payment.Uint64()) / params.Ether)
}

func (m *Monitor) HandleMiss(req *abi.OracleOracleRequest) {
	m.missCounter.WithLabelValues(string(req.SpecId[:]), req.Requester.String()).Inc()
	m.revenueCounter.WithLabelValues(string(req.SpecId[:]), req.Requester.String(), "missed").Add(float64(req.Payment.Uint64()) / params.Ether)
}
