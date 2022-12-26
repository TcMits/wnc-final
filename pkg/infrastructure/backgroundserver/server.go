package backgroundserver

import (
	"github.com/hibiken/asynq"
)

type WorkerServer struct {
	server  *asynq.Server
	handler asynq.Handler
}

func NewClient(url string, pwd string, db int) *asynq.Client {
	c := asynq.NewClient(
		asynq.RedisClientOpt{
			Addr:     url,
			Password: pwd,
			DB:       db,
		},
	)
	return c
}

func NewWorkerServer(handler asynq.Handler, url string, pwd string, db int) *WorkerServer {
	s := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     url,
			Password: pwd,
			DB:       db,
		},
		asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
		},
	)

	return &WorkerServer{
		server:  s,
		handler: handler,
	}
}

func (s *WorkerServer) Run() error {
	return s.server.Run(s.handler)
}
