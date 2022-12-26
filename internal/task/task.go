package task

import (
	"context"
	"encoding/json"

	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/TcMits/wnc-final/pkg/tool/mail"
	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	typeConfirmEmail = "email:confirm"
)

type (
	EmailTaskExecutor struct {
		client *asynq.Client
	}
	TaskHandler = func(context.Context, *asynq.Task) error
)

func newTask[T any](pl T, typeTask string) (*asynq.Task, error) {
	payload, err := json.Marshal(pl)
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(typeTask, payload), nil
}

func (s *EmailTaskExecutor) ExecuteTask(ctx context.Context, pl *mail.EmailPayload) error {
	if pl == nil {
		pl = new(mail.EmailPayload)
	}
	task, err := newTask(pl, typeConfirmEmail)
	if err != nil {
		return err
	}
	_, err = s.client.EnqueueContext(ctx, task)
	if err != nil {
		return err
	}
	return nil
}

func mailTaskHandlerWrapper(host, user, password string, port int, l logger.Interface) TaskHandler {
	return func(ctx context.Context, t *asynq.Task) error {
		p := new(mail.EmailPayload)
		if err := json.Unmarshal(t.Payload(), p); err != nil {
			return err
		}
		l.Info("Sending email...")
		if err := mail.SendMail(p, user, password, host, port); err != nil {
			return err
		}
		return nil
	}
}

func NewHandler() *asynq.ServeMux {
	mux := asynq.NewServeMux()
	return mux
}

func GetEmailTaskExecutor(c *asynq.Client) IExecuteTask[*mail.EmailPayload] {
	return &EmailTaskExecutor{
		client: c,
	}
}

func RegisterTask(handler *asynq.ServeMux, l logger.Interface, host, user, password string, port int) {
	handler.HandleFunc(typeConfirmEmail, mailTaskHandlerWrapper(host, user, password, port, l))
}
