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
	typeDebtNotify = "notify:debt"
)

type (
	DebtTaskExecutor struct {
		client *asynq.Client
	}
	DebtNotifyPayload struct {
		ID     uuid.UUID
		UserID uuid.UUID
		Event  string
	}
)

func (s *DebtTaskExecutor) ExecuteTask(ctx context.Context, pl *DebtNotifyPayload) error {
	if pl == nil {
		pl = new(DebtNotifyPayload)
	}
	task, err := newTask(pl, typeDebtNotify)
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
		p := new(DebtNotifyPayload)
		if err := json.Unmarshal(t.Payload(), p); err != nil {
			return err
		}
		msgpl := new(sse.MessagePayload)
		msgpl.If = func(c *model.Customer) bool {
			return c.ID == p.UserID
		}
		msgpl.Msg = p.ID.String()
		msgpl.Event = p.Event
		l.Info("Notify debt...")
		err := b.Notify(msgpl)
		if err != nil {
			return err
		}
		return nil
	}
}

func GetDebtTaskExecutor(c *asynq.Client) IExecuteTask[*DebtNotifyPayload] {
	return &DebtTaskExecutor{
		client: c,
	}
}

func RegisterDebtTaskHandler(handler *asynq.ServeMux, l logger.Interface, b sse.INotify) {
	handler.HandleFunc(typeDebtNotify, debtTaskHandlerWrapper(b, l))
}
