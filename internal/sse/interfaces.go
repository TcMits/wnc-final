package sse

type (
	INotify interface {
		Notify(*MessagePayload) error
	}
)
