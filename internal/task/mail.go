package task

import (
	"context"

	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/TcMits/wnc-final/pkg/tool/mail"
)

type (
	EmailTaskExecutor struct {
		host     string
		user     string
		password string
		sender   string
		port     int
		l        logger.Interface
	}
)

func (s *EmailTaskExecutor) ExecuteTask(ctx context.Context, pl *mail.EmailPayload) error {
	if pl == nil {
		pl = new(mail.EmailPayload)
	}
	go mailTaskHandler(s.host, s.user, s.password, s.sender, s.port, s.l, pl)
	return nil
}
func mailTaskHandler(host, user, password, sender string, port int, l logger.Interface, p *mail.EmailPayload) {
	if err := mail.SendMail(p, user, password, host, sender, port); err != nil {
		l.Warn("Sending email failed due to: %s", err)
	}
	l.Info("Sending email successfully...")
}

func GetEmailTaskExecutor(host, user, password, sender string, port int, l logger.Interface) IExecuteTask[*mail.EmailPayload] {
	return &EmailTaskExecutor{
		host:     host,
		user:     user,
		password: password,
		sender:   sender,
		port:     port,
		l:        l,
	}
}
