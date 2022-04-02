package ztask

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
)

type Server struct {
	Srv *asynq.Server
}

func NewServer(r asynq.RedisConnOpt, cfg asynq.Config) *Server {
	return &Server{Srv: asynq.NewServer(r, cfg)}
}

type EasyParam struct {
	Addr     string
	Password string
	DB       int
}

func NewEasyServer(param EasyParam) *Server {
	return NewServer(asynq.RedisClientOpt{
		Addr:     param.Addr,
		Password: param.Password,
		DB:       param.DB,
	}, asynq.Config{
		Concurrency: 1,
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
	})
}

func (s *Server) Run(ctx context.Context) {
	mux := asynq.NewServeMux()
	TaskHandle(mux)

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
