package rpc

import (
	"github.com/ethereum/go-ethereum/rpc"
	"net/http"
)

var RPCServer = NewRPCServer()

type Server struct {
	Port      uint16
	rpcServer *rpc.Server
	server    *http.Server
}

func NewRPCServer() *Server {
	rpcServer := rpc.NewServer()
	service := SmartPoolService{}
	rpcServer.RegisterName("eth", service)
	return &Server{uint16(1633), rpcServer, &http.Server{
		Addr:    ":1633",
		Handler: rpcServer,
	}}
}

func (s Server) Start() {
	s.server.ListenAndServe()
}
