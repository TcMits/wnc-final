package sse

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/customer/middleware"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type (
	MessagePayload struct {
		Msg []byte
		If  func(*model.Customer) bool
	}

	Broker struct {
		notifier       chan MessagePayload
		newClients     chan chan MessagePayload
		closingClients chan chan MessagePayload
		clients        map[chan MessagePayload]bool
		logger         logger.Interface
	}
)

func NewBroker(l logger.Interface) *Broker {
	b := &Broker{
		notifier:       make(chan MessagePayload, 1),
		newClients:     make(chan chan MessagePayload),
		closingClients: make(chan chan MessagePayload),
		clients:        make(map[chan MessagePayload]bool),
		logger:         l,
	}
	go b.listen()
	return b
}
func (b *Broker) Notify(pl *MessagePayload) error {
	b.notifier <- *pl
	return nil
}
func (b *Broker) AddClient(c chan MessagePayload) error {
	b.newClients <- c
	return nil
}
func (b *Broker) RemoveClient(c chan MessagePayload) error {
	b.closingClients <- c
	return nil
}
func (b *Broker) listen() {
	for {
		select {
		case s := <-b.newClients:
			b.clients[s] = true
			b.logger.Info("Client added. %d registered clients", len(b.clients))

		case s := <-b.closingClients:
			delete(b.clients, s)
			b.logger.Info("Removed client. %d registered clients", len(b.clients))

		case event := <-b.notifier:
			for clientMessageChan := range b.clients {
				clientMessageChan <- event
			}
		}
	}
}
func (b *Broker) ServeHTTP(ctx iris.Context) {
	flusher, ok := ctx.ResponseWriter().Flusher()
	if !ok {
		ctx.StopWithJSON(iris.StatusHTTPVersionNotSupported, struct {
			Message string `json:"message"`
			Code    string `json:"code"`
			Detail  string `json:"detail"`
		}{
			Message: "Streaming unsupported!",
			Detail:  "The HTTP version not support text/event-stream content type ",
			Code:    "UNKNOWN",
		})
		return
	}

	ctx.ContentType("application/json, text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	ctx.Header("Access-Control-Allow-Origin", "*")

	messageChan := make(chan MessagePayload)

	b.AddClient(messageChan)

	ctx.OnClose(func(iris.Context) {
		b.RemoveClient(messageChan)
	})
	user := middleware.GetUserFromCtxAsCustomer(ctx)
	for {
		data := <-messageChan
		if data.If(user) {
			ctx.Writef("data: %s\n\n", data.Msg)
			flusher.Flush()
		}
	}
}
