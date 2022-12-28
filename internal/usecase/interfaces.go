// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/google/uuid"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks.go -package=usecase
type (
	IGetUserUseCase interface {
		GetUser(context.Context, map[string]any) (any, error)
	}
	IGetConfigUseCase interface {
		GetProductOwnerName() *string
		GetSecret() *string
	}
	iListUseCase[ModelType, ModelOrderInput, ModelWhereInput any] interface {
		List(context.Context, *int, *int, ModelOrderInput, ModelWhereInput) ([]ModelType, error)
	}
	iCreateUseCase[ModelType, ModelCreateInput any] interface {
		Create(context.Context, ModelCreateInput) (ModelType, error)
	}
	iValidateUpdateInput[ModelType, ModelUpdateInput any] interface {
		Validate(context.Context, ModelType, ModelUpdateInput) (ModelUpdateInput, error)
	}
	iValidateCreateInput[ModelCreateInput any] interface {
		Validate(context.Context, ModelCreateInput) (ModelCreateInput, error)
	}
	iUpdateUseCase[ModelType, ModelUpdateInput any] interface {
		Update(context.Context, ModelType, ModelUpdateInput) (ModelType, error)
	}
	iDetailUseCase[ModelType any] interface {
		Detail(context.Context, *uuid.UUID) (ModelType, error)
	}
	iDeleteUseCase interface {
		Delete(context.Context, *uuid.UUID) error
	}
	iEntityUseCase[ModelType, ModelOrderInput, ModelWhereInput, ModelCreateInput any] interface {
		iListUseCase[ModelType, ModelOrderInput, ModelWhereInput]
		iCreateUseCase[ModelType, ModelCreateInput]
		iDetailUseCase[ModelType]
		iDeleteUseCase
	}
	iAuthenticationUseCase[LoginInput, ModelType any] interface {
		IGetUserUseCase
		IGetConfigUseCase
		Login(context.Context, LoginInput) (any, error)
		ValidateLoginInput(context.Context, LoginInput) (LoginInput, error)
		RenewToken(context.Context, *string) (any, error)
		Logout(context.Context) error
	}
)

type (
	ICustomerConfigUseCase interface {
		IGetConfigUseCase
		GetFeeAmount() *float64
		GetFeeDesc() *string
	}
	ICustomerGetUserUseCase interface {
		IGetUserUseCase
	}
	ICustomerGetFirstUseCase interface {
		GetFirst(context.Context, *model.CustomerOrderInput, *model.CustomerWhereInput) (*model.Customer, error)
	}
	ICustomerListUseCase interface {
		iListUseCase[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput]
	}
	ICustomerMeUseCase interface {
		ICustomerConfigUseCase
		ICustomerGetUserUseCase
	}
	ICustomerAuthUseCase interface {
		iAuthenticationUseCase[*model.CustomerLoginInput, *model.Customer]
	}
	ICustomerBankAccountUpdateUseCase interface {
		iUpdateUseCase[*model.BankAccount, *model.BankAccountUpdateInput]
	}
	ICustomerUpdateUseCase interface {
		iUpdateUseCase[*model.Customer, *model.CustomerUpdateInput]
	}
	ICustomerBankAccountListUseCase interface {
		iListUseCase[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
	}
	ICustomerBankAccountGetFirstUseCase interface {
		GetFirst(context.Context, *model.BankAccountOrderInput, *model.BankAccountWhereInput) (*model.BankAccount, error)
	}
	ICustomerBankAccountValidateUpdateInputUseCase interface {
		iValidateUpdateInput[*model.BankAccount, *model.BankAccountUpdateInput]
	}
	ICustomerBankAccountUseCase interface {
		ICustomerGetUserUseCase
		ICustomerConfigUseCase
		ICustomerBankAccountUpdateUseCase
		ICustomerBankAccountValidateUpdateInputUseCase
		ICustomerBankAccountListUseCase
	}
	ICustomerTransactionValidateConfirmInputUseCase interface {
		ValidateConfirmInput(context.Context, *model.Transaction, *string, *string) error
	}
	ICustomerTransactionConfirmSuccessUseCase interface {
		ConfirmSuccess(context.Context, *model.Transaction, *string) (*model.Transaction, error)
	}
	ICustomerTransactionCreateUseCase interface {
		Create(context.Context, *model.TransactionCreateInput, bool) (*model.Transaction, error)
	}
	ICustomerTransactionValidateCreateInputUseCase interface {
		Validate(context.Context, *model.TransactionCreateInput, bool) (*model.TransactionCreateInput, error)
	}
	ICustomerTransactionListUseCase interface {
		iListUseCase[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
	}
	ICustomerTransactionGetFirstMineUseCase interface {
		GetFirstMine(context.Context, *model.TransactionOrderInput, *model.TransactionWhereInput) (*model.Transaction, error)
	}
	ICustomerTransactionListMineUseCase interface {
		ListMine(context.Context, *int, *int, *model.TransactionOrderInput, *model.TransactionWhereInput) ([]*model.Transaction, error)
	}
	ICustomerTransactionUpdateUseCase interface {
		iUpdateUseCase[*model.Transaction, *model.TransactionUpdateInput]
	}
	ICustomerTransactionUseCase interface {
		ICustomerGetUserUseCase
		ICustomerConfigUseCase
		ICustomerTransactionCreateUseCase
		ICustomerTransactionListUseCase
		ICustomerTransactionValidateCreateInputUseCase
		ICustomerTransactionConfirmSuccessUseCase
		ICustomerTransactionValidateConfirmInputUseCase
		ICustomerTransactionListMineUseCase
		ICustomerTransactionGetFirstMineUseCase
	}
	// debt
	ICustomerDebtListUseCase interface {
		iListUseCase[*model.Debt, *model.DebtOrderInput, *model.DebtWhereInput]
	}
	ICustomerDebtValidateCreateInputUseCase interface {
		iValidateCreateInput[*model.DebtCreateInput]
	}
	ICustomerDebtCreateUseCase interface {
		iCreateUseCase[*model.Debt, *model.DebtCreateInput]
	}
	ICustomerDebtGetFirstMineUseCase interface {
		GetFirstMine(context.Context, *model.DebtOrderInput, *model.DebtWhereInput) (*model.Debt, error)
	}
	ICustomerDebtListMineUseCase interface {
		ListMine(context.Context, *int, *int, *model.DebtOrderInput, *model.DebtWhereInput) ([]*model.Debt, error)
	}
	ICustomerDebtUseCase interface {
		ICustomerGetUserUseCase
		ICustomerConfigUseCase
		ICustomerDebtListUseCase
		ICustomerDebtValidateCreateInputUseCase
		ICustomerDebtCreateUseCase
		ICustomerDebtGetFirstMineUseCase
		ICustomerDebtListMineUseCase
	}
	// stream
	ICustomerStreamUseCase interface {
		ICustomerGetUserUseCase
		ICustomerConfigUseCase
	}
)
