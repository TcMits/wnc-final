package task

import (
	"context"
	"encoding/json"

	"github.com/TcMits/wnc-final/internal/sse"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/google/uuid"
)

type (
	DebtCreateTaskExecutor struct {
		b sse.INotify
		l logger.Interface
	}
	DebtCreateNotifyPayload struct {
		UserID uuid.UUID
	}
)

func (s *DebtCreateTaskExecutor) ExecuteTask(ctx context.Context, pl *DebtCreateNotifyPayload) error {
	if pl == nil {
		pl = new(DebtCreateNotifyPayload)
	}
	go func() {
		msgpl := new(sse.MessagePayload)
		msgpl.If = func(c *model.Customer) bool {
			return c.ID == pl.UserID
		}
		var err error
		msgpl.Msg, err = json.Marshal("hello world")
		if err != nil {
			s.l.Warn("Notify failed due to: %s", err)
		}
		err = s.b.Notify(msgpl)
		if err != nil {
			s.l.Warn("Notify failed due to: %s", err)
		}
	}()
	return nil
}

func GetDebtTaskExecutor(b sse.INotify, l logger.Interface) IExecuteTask[*DebtCreateNotifyPayload] {
	return &DebtCreateTaskExecutor{
		b: b,
		l: l,
	}
}
