package backgroundtask

import (
	"github.com/hibiken/asynq"
)

type Server struct {
	server  *asynq.Server
	handler asynq.Handler
}

func NewHandler() *asynq.ServeMux {
	mux := asynq.NewServeMux()
	return mux
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

func NewServer(handler asynq.Handler, url string, pwd string, db int) *Server {
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

	return &Server{
		server:  s,
		handler: handler,
	}
}

func (s *Server) Run() error {
	return s.server.Run(s.handler)
}
