package task

import (
	"context"
	"encoding/json"

	"github.com/TcMits/wnc-final/internal/sse"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/hibiken/asynq"
)

type (
	TaskHandler = func(context.Context, *asynq.Task) error
)

func newTask[T any](pl T, typeTask string) (*asynq.Task, error) {
	payload, err := json.Marshal(pl)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(typeTask, payload), nil
}

func RegisterTask(handler *asynq.ServeMux, l logger.Interface, host, user, password, sender string, port int, b sse.INotify) {
	RegisterMailTaskHandler(handler, l, host, user, password, sender, port)
	RegisterDebtTaskHandler(handler, l, b)
}
