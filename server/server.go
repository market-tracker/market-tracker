package server

import (
	"fmt"

	"github.com/market-tracker/market-tracker/app"
)

type Server struct {
	port int32
}

var server *Server

func (s *Server) Start(callback func()) {
	addr := fmt.Sprintf(":%d", s.port)
	a := app.GetInstance()
	a.Start()
	if err := a.Run(addr); err != nil {
		callback()
	}
}

func InitServer(port int32) *Server {
	if server != nil {
		return server
	}
	return &Server{port: port}
}
