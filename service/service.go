package service

import (
	"fmt"
	"github.com/ONSdigital/dis-routing-performance-test/handle-anything-server"

	_ "github.com/ONSdigital/dis-routing-performance-test/handle-anything-server"
)

type Service struct {
	HandleAnythingPort   int
	handleAnythingServer Server
}

func (s *Service) Run() error {

	s.handleAnythingServer = Server{
		Name:    "handleAnythingServer",
		Addr:    fmt.Sprintf("localhost:%d", s.HandleAnythingPort),
		Handler: handle_anything_server.Handler(),
	}

	s.handleAnythingServer.Start()

	return nil
}

func (s *Service) Shutdown() {
	s.handleAnythingServer.Stop()
}
