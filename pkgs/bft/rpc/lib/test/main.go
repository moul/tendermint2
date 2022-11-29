package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tendermint/tendermint2/pkgs/log"
	osm "github.com/tendermint/tendermint2/pkgs/os"

	rpcserver "github.com/tendermint/tendermint2/pkgs/bft/rpc/lib/server"
	rpctypes "github.com/tendermint/tendermint2/pkgs/bft/rpc/lib/types"
)

var routes = map[string]*rpcserver.RPCFunc{
	"hello_world": rpcserver.NewRPCFunc(HelloWorld, "name,num"),
}

func HelloWorld(ctx *rpctypes.Context, name string, num int) (Result, error) {
	return Result{fmt.Sprintf("hi %s %d", name, num)}, nil
}

type Result struct {
	Result string
}

func main() {
	var (
		mux    = http.NewServeMux()
		logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	)

	// Stop upon receiving SIGTERM or CTRL-C.
	osm.TrapSignal(func() {})

	rpcserver.RegisterRPCFuncs(mux, routes, logger)
	config := rpcserver.DefaultConfig()
	listener, err := rpcserver.Listen("0.0.0.0:8008", config)
	if err != nil {
		osm.Exit(err.Error())
	}
	rpcserver.StartHTTPServer(listener, mux, logger, config)
}
