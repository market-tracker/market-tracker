package server

type IServer interface {
	Init(port int32) *IServer
	Start()
}
