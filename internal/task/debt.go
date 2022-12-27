package task

import (
	"context"
	"encoding/json"

	"github.com/TcMits/wnc-final/internal/sse"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
)

const (
	typeDebtCreateNotify = "notify:debt-create"
)

type (
	DebtCreateTaskExecutor struct {
		client *asynq.Client
	}
	DebtCreateNotifyPayload struct {
		UserID uuid.UUID
	}
)

func (s *DebtCreateTaskExecutor) ExecuteTask(ctx context.Context, pl *DebtCreateNotifyPayload) error {
	if pl == nil {
		pl = new(DebtCreateNotifyPayload)
	}
	task, err := newTask(pl, typeDebtCreateNotify)
	if err != nil {
		return err
	}
	_, err = s.client.EnqueueContext(ctx, task)
	if err != nil {
		return err
	}
	return nil
}

func debtTaskHandlerWrapper(b sse.INotify, l logger.Interface) TaskHandler {
	return func(ctx context.Context, t *asynq.Task) error {
		p := new(DebtCreateNotifyPayload)
		if err := json.Unmarshal(t.Payload(), p); err != nil {
			return err
		}
		pl := new(sse.MessagePayload)
		pl.If = func(c *model.Customer) bool {
			return c.ID == p.UserID
		}
		var err error
		pl.Msg, err = json.Marshal("hello world")
		if err != nil {
			return err
		}
		err = b.Notify(pl)
		if err != nil {
			return err
		}
		return nil
	}
}
