package mail_test

import (
	"testing"

	"github.com/TcMits/wnc-final/pkg/tool/mail"
	"github.com/stretchr/testify/require"

	"github.com/TcMits/wnc-final/config"
)

func TestSendMail(t *testing.T) {
	t.Parallel()
	cfg, _ := config.NewConfig()
	err := mail.SendMail(&mail.EmailPayload{
		Subject: "Sample subject",
		Message: "Sample message",
		To: []string{
			"pass email here",
		},
	}, cfg.Mail.User, cfg.Mail.Password, cfg.Mail.Host, "pass email here", cfg.Mail.Port)
	require.Nil(t, err)
}
