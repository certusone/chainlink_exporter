package main

import (
	"chainlink_exporter/abi"
	"context"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"sync"
	"time"
)

type (
	AggregatorMonitor struct {
		aggregator *abi.Aggregator
		address    common.Address

		pendingJobs map[string]*abi.OracleOracleRequest

		seenRequestIDs map[string]bool

		monitor *Monitor
		lock    sync.Mutex
	}
)

func NewAggregatorMonitor(agg *abi.Aggregator, addr common.Address, m *Monitor) *AggregatorMonitor {
	return &AggregatorMonitor{
		aggregator:     agg,
		pendingJobs:    map[string]*abi.OracleOracleRequest{},
		seenRequestIDs: map[string]bool{},
		monitor:        m,
		address:        addr,
	}
}

func (a *AggregatorMonitor) Monitor() {
	for {
		zap.L().Debug("Starting aggregator routine", zap.String("address", a.address.String()))
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			defer cancel()

			resChan := make(chan *abi.AggregatorChainlinkFulfilled)
			sub, err := a.aggregator.WatchChainlinkFulfilled(&bind.WatchOpts{Context: ctx}, resChan, nil)
			if err != nil {
				zap.L().Error("failed to watch aggregator fulfillment", zap.Error(err), zap.String("address", a.address.String()))
				return
			}

			for {
				select {
				case err = <-sub.Err():
					zap.L().Error("aggregator fulfillment subscription errored", zap.Error(err), zap.String("address", a.address.String()))
					return
				case res, has := <-resChan:
					if !has {
						zap.L().Error("head subscription closed", zap.Error(err), zap.String("address", a.address.String()))
						return
					}
					a.handleFulfillment(res)
				}
			}
		}()

		zap.L().Warn("aggregator routine died. restarting in 5sec", zap.String("address", a.address.String()))
		time.Sleep(5 * time.Second)
	}
}

func (a *AggregatorMonitor) HandleNewBlock(height uint64) {
	a.lock.Lock()
	defer a.lock.Unlock()

	for reqID, n := range a.pendingJobs {
		delta := int(height) - int(n.Raw.BlockNumber)
		// todo make dynamic
		if delta > 15 {
			zap.L().Info("job fulfillment slot missed", zap.Uint64("height", n.Raw.BlockNumber),
				zap.String("requester", n.Requester.String()), zap.Binary("request_id", n.RequestId[:]),
				zap.String("spec_id", sanitizeSpecID(n.SpecId)))
			delete(a.pendingJobs, reqID)
			a.monitor.HandleMiss(n)
		}
	}
}

func (a *AggregatorMonitor) handleRequest(res *abi.OracleOracleRequest) {
	a.lock.Lock()
	defer a.lock.Unlock()
	requestIDString := hex.EncodeToString(res.RequestId[:])
	if _, exists := a.seenRequestIDs[requestIDString]; exists {
		zap.L().Info("request dropped; already seen same reqID", zap.Uint64("height", res.Raw.BlockNumber),
			zap.String("requester", res.Requester.String()), zap.Binary("request_id", res.RequestId[:]),
			zap.String("spec_id", sanitizeSpecID(res.SpecId)), zap.Uint64("request_height", res.Raw.BlockNumber))
		return
	} else {
		a.seenRequestIDs[requestIDString] = true
	}

	a.pendingJobs[requestIDString] = res
}

func (a *AggregatorMonitor) handleFulfillment(res *abi.AggregatorChainlinkFulfilled) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if job, ok := a.pendingJobs[hex.EncodeToString(res.Id[:])]; ok {
		zap.L().Info("job fulfilled", zap.Uint64("height", res.Raw.BlockNumber),
			zap.String("requester", job.Requester.String()), zap.Binary("request_id", job.RequestId[:]),
			zap.String("spec_id", sanitizeSpecID(job.SpecId)), zap.Uint64("request_height", job.Raw.BlockNumber))
		delete(a.pendingJobs, hex.EncodeToString(res.Id[:]))

		a.monitor.HandleFulfillment(res, job)
	}
}
