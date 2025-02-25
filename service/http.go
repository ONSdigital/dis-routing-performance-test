package service

import (
	"context"
	"log/slog"
	"net/http"
)

type Server struct {
	Name    string
	Addr    string
	Handler http.Handler
	server  http.Server
}

func (s *Server) Start() {
	s.server = http.Server{
		Addr:    s.Addr,
		Handler: s.Handler,
	}
	slog.Info("starting server", "name", s.Name, "addr", s.Addr)
	go func() {
		s.server.ListenAndServe()
		slog.Info("stopped server", "name", s.Name)
	}()
}

func (s *Server) Stop() {
	slog.Info("stopping server", "name", s.Name)
	s.server.Shutdown(context.Background()) // Should really use a timeout but taking a dirty shortcut for POC
}
