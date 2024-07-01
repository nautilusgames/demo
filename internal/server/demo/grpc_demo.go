package demo

import (
	"github.com/nautilusgames/demo/api/v1/demo"
)

func NewServer() demo.DemoServer {
	return &demoServer{}
}

type demoServer struct {
	demo.UnimplementedDemoServer
}
