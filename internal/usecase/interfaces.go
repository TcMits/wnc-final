// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/TcMits/wnc-final/pkg/entity/model"
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
	iAuthenticationUseCase[LoginInput, ModelType any] interface {
		iGetUserUseCase
		iConfigUseCase
		Login(context.Context, LoginInput) (any, error)
		ValidateLoginInput(context.Context, LoginInput) (LoginInput, error)
		RenewToken(context.Context, *string) (any, error)
		Logout(context.Context, ModelType) error
	}
)

type (
	ICustomerConfigUseCase interface {
		iConfigUseCase
	}
	ICustomerGetUserUseCase interface {
		iGetUserUseCase
	}
	ICustomerMeUseCase interface {
		ICustomerConfigUseCase
		ICustomerGetUserUseCase
	}
	ICustomerAuthUseCase interface {
		iAuthenticationUseCase[*model.CustomerLoginInput, *model.Customer]
	}
)
