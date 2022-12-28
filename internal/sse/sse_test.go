package sse_test

import (
	"encoding/json"
	"testing"

	"github.com/TcMits/wnc-final/internal/sse"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/stretchr/testify/require"
)

func TestNotify(t *testing.T) {
	t.Parallel()
	l := logger.New(logger.DebugLevel)
	b := sse.NewBroker(l)
	messageChan := make(chan sse.MessagePayload)
	msg, _ := json.Marshal("foo")
	b.AddClient(messageChan)
	go func() {
		b.Notify(&sse.MessagePayload{
			If: func(c *model.Customer) bool {
				return c.FirstName == "foo"
			},
			Msg: msg,
		})
	}()

	data := <-messageChan
	b.RemoveClient(messageChan)
	require.Equal(t, data.Msg, msg)
	user := &model.Customer{
		FirstName: "foo",
	}
	require.True(t, data.If(user))
}
