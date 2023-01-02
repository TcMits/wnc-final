package repository

import (
	"context"

	"github.com/TcMits/wnc-final/pkg/entity/model"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks.go -package=repository

type (
	ListModelRepository[ModelType, OrderInputType, WhereInputType any] interface {
		List(context.Context, *int, *int, OrderInputType, WhereInputType) ([]ModelType, error)
	}
	CreateModelRepository[ModelType, CreateInputType any] interface {
		Create(context.Context, CreateInputType) (ModelType, error)
	}
	DeleteModelRepository[ModelType any] interface {
		Delete(context.Context, ModelType) error
	}
	UpdateModelRepository[ModelType any, UpdateInputType any] interface {
		Update(context.Context, ModelType, UpdateInputType) (ModelType, error)
	}
)

type (
	ITransactionConfirmSuccessRepository interface {
		ConfirmSuccess(context.Context, *model.Transaction, *model.TransactionCreateInput) (*model.Transaction, error)
	}
)

type (
	IDebtFullfillRepository interface {
		Fulfill(context.Context, *model.Debt, *model.DebtUpdateInput) (*model.Debt, error)
	}
)
