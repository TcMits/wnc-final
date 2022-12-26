package task

import "context"

type (
	IExecuteTask[T any] interface {
		ExecuteTask(context.Context, T) error
	}
)
