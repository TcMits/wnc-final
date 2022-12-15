// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/google/uuid"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks.go -package=usecase
type (
	iGetUserUseCase interface {
		GetUser(context.Context, map[string]any) (any, error)
	}
	iConfigUseCase interface {
		GetSecret() (*string, error)
	}
	iListUseCase[ModelType any] interface {
		List(context.Context, *int, *int) ([]ModelType, error)
	}
	iCreateUseCase[ModelType, ModelCreateInput any] interface {
		Create(context.Context, ModelCreateInput) (ModelType, error)
	}
	iDetailUseCase[ModelType any] interface {
		Detail(context.Context, *uuid.UUID) (ModelType, error)
	}
	iDeleteUseCase interface {
		Delete(context.Context, *uuid.UUID) error
	}
	iEntityUseCase[ModelType, ModelCreateInput any] interface {
		iListUseCase[ModelType]
		iCreateUseCase[ModelType, ModelCreateInput]
		iDetailUseCase[ModelType]
		iDeleteUseCase
	}
	iAuthenticationUseCase interface {
		iGetUserUseCase
		iConfigUseCase
	}
)

type (
	ICustomerConfigUseCase interface {
		iConfigUseCase
	}
	ICustomerAuthenticationUseCase interface {
		iAuthenticationUseCase
	}
	IMeCustomerUseCase interface {
		ICustomerAuthenticationUseCase
	}
)
