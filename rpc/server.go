package rpc

import (
	"github.com/ethereum/go-ethereum/rpc"
	"net/http"
)

var RPCServer = NewRPCServer()

type server struct {
	Port      uint16
	rpcServer *rpc.Server
	server    *http.Server
}

func NewRPCServer() *server {
	rpcServer := rpc.NewServer()
	service := SmartPoolService{}
	rpcServer.RegisterName("eth", service)
	return &server{uint16(1633), rpcServer, &http.Server{
		Addr:    ":1633",
		Handler: rpcServer,
	}}
}

func (s server) Start() {
	s.server.ListenAndServe()
}
