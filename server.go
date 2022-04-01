package ztask

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
	"github.com/zhoujianxuan/ztask/tasks"
)

type Server struct {
	Srv *asynq.Server
}

func NewServer(srv *asynq.Server) *Server {
	return &Server{Srv: srv}
}

func (s *Server) Run(ctx context.Context) {
	mux := asynq.NewServeMux()
	tasks.TaskHandle(mux)

	go func() {
		if r := recover(); r != nil {
			log.Fatalf("could not run server: %v", r)
		}
		if err := s.Srv.Run(mux); err != nil {
			log.Fatalf("could not run server: %v", err)
		}
	}()

	select {
	case <-ctx.Done():
		s.Srv.Shutdown()
	}
}
