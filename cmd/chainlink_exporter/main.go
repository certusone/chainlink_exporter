package main

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	"net/http"
	"os"
	"time"

	_ "net/http/pprof"
)

var (
	oracleAddr      = os.Getenv("ADDRESS")
	fulfillmentAddr = os.Getenv("NODE_ADDRESS")
	linkAddr        = os.Getenv("LINK_ADDRESS")
	rpcHost         = os.Getenv("RPC")
	lAddr           = os.Getenv("LADDR")
)

func main() {
	l, _ := zap.NewProduction()
	zap.ReplaceGlobals(l)

	if lAddr == "" {
		panic("LADDR must be set")
	}
	if rpcHost == "" {
		panic("RPC must be set")
	}
	if fulfillmentAddr == "" {
		panic("NODE_ADDRESS must be set")
	}
	if oracleAddr == "" {
		panic("ADDRESS must be set")
	}
	if linkAddr == "" {
		zap.L().Warn("LINK_ADDRESS isn't set. Falling back to mainnet default.")
		linkAddr = "0x514910771af9ca656af840dff83e8264ecf986ca"
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	c, err := ethclient.DialContext(ctx, rpcHost)
	if err != nil {
		panic(err)
	}

	mon, err := NewMonitor(common.HexToAddress(oracleAddr), common.HexToAddress(fulfillmentAddr), common.HexToAddress(linkAddr), c)
	if err != nil {
		panic(err)
	}
	mon.Start()

	http.Handle("/metrics", promhttp.Handler())

	panic(http.ListenAndServe(lAddr, nil))
}
