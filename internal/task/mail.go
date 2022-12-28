package task

import (
	"context"
	"encoding/json"

	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/TcMits/wnc-final/pkg/tool/mail"
	"github.com/hibiken/asynq"
)

const (
	typeConfirmEmail = "email:confirm"
)

type (
	EmailTaskExecutor struct {
		client *asynq.Client
	}
)

func (s *EmailTaskExecutor) ExecuteTask(ctx context.Context, pl *mail.EmailPayload) error {
	if pl == nil {
		pl = new(mail.EmailPayload)
	}
	task, err := NewTask(pl, typeConfirmEmail)
	if err != nil {
		return err
	}
	_, err = s.client.EnqueueContext(ctx, task)
	if err != nil {
		return err
	}
	return nil
}

func MailTaskHandlerWrapper(host, user, password, sender string, port int, l logger.Interface) TaskHandler {
	return func(ctx context.Context, t *asynq.Task) error {
		p := new(mail.EmailPayload)
		if err := json.Unmarshal(t.Payload(), p); err != nil {
			return err
		}
		l.Info("Sending email...")
		if err := mail.SendMail(p, user, password, host, sender, port); err != nil {
			return err
		}
		return nil
	}
}
func GetEmailTaskExecutor(c *asynq.Client) IExecuteTask[*mail.EmailPayload] {
	return &EmailTaskExecutor{
		client: c,
	}
}
