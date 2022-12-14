// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks.go -package=usecase
type (
	IGetUserUseCase interface {
		GetUser(context.Context, map[string]any) (any, error)
	}
	IConfigUseCase interface {
		GetSecret() (*string, error)
	}
	IListUseCase[ModelType any] interface {
		List(context.Context, *int, *int) ([]ModelType, error)
	}
	ICreateUseCase[ModelType, ModelCreateInput any] interface {
		Create(context.Context, ModelCreateInput) (ModelType, error)
	}
	IDetailUseCase[ModelType any] interface {
		Detail(context.Context, *uint16) (ModelType, error)
	}
	IDeleteUseCase interface {
		Delete(context.Context, *uint16) error
	}
	IEntityUseCase[ModelType, ModelCreateInput any] interface {
		IListUseCase[ModelType]
		ICreateUseCase[ModelType, ModelCreateInput]
		IDetailUseCase[ModelType]
		IDeleteUseCase
	}
	IAuthenticationUseCase interface {
		IGetUserUseCase
		IConfigUseCase
	}
)
