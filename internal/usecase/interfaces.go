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
	IOptionsUseCase interface {
		GetDebtStatus(context.Context) []string
		GetEvents(context.Context) []string
	}
	iListUseCase[ModelType, ModelOrderInput, ModelWhereInput any] interface {
		List(context.Context, *int, *int, ModelOrderInput, ModelWhereInput) ([]ModelType, error)
	}
	iCreateUseCase[ModelType, ModelCreateInput any] interface {
		Create(context.Context, ModelCreateInput) (ModelType, error)
	}
	iValidateUpdateInput[ModelType, ModelUpdateInput any] interface {
		ValidateUpdate(context.Context, ModelType, ModelUpdateInput) (ModelUpdateInput, error)
	}
	iValidateCreateInput[ModelCreateInput any] interface {
		ValidateCreate(context.Context, ModelCreateInput) (ModelCreateInput, error)
	}
	iUpdateUseCase[ModelType, ModelUpdateInput any] interface {
		Update(context.Context, ModelType, ModelUpdateInput) (ModelType, error)
	}
	iDetailUseCase[ModelType any] interface {
		Detail(context.Context, *uuid.UUID) (ModelType, error)
	}
	iDeleteUseCase[ModelType any] interface {
		Delete(context.Context, ModelType) error
	}
	iEntityUseCase[ModelType, ModelOrderInput, ModelWhereInput, ModelCreateInput any] interface {
		iListUseCase[ModelType, ModelOrderInput, ModelWhereInput]
		iCreateUseCase[ModelType, ModelCreateInput]
		iDetailUseCase[ModelType]
		iDeleteUseCase[ModelType]
	}
	iAuthenticationUseCase[LoginInput, ModelType any] interface {
		IGetUserUseCase
		IGetConfigUseCase
		Login(context.Context, LoginInput) (any, error)
		ValidateLoginInput(context.Context, LoginInput) (LoginInput, error)
		RenewToken(context.Context, *string) (any, error)
		Logout(context.Context) error
	}
	IIsNextUseCase[ModelType, ModelOrderInput, ModelWhereInput any] interface {
		IsNext(context.Context, int, int, ModelOrderInput, ModelWhereInput) (bool, error)
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
	ICustomerGetUserFromCtxUseCase interface {
		GetUserFromCtx(context.Context) (*model.Customer, error)
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
		ICustomerValidateChangePasswordUseCase
		ICustomerChangePasswordUseCase
		ICustomerGetUserFromCtxUseCase
	}
	// change password
	ICustomerValidateChangePasswordUseCase interface {
		ValidateChangePassword(context.Context, *model.CustomerChangePasswordInput) (*model.CustomerChangePasswordInput, error)
	}
	ICustomerChangePasswordUseCase interface {
		ChangePassword(context.Context, *model.CustomerChangePasswordInput) (*model.Customer, error)
	}
	ICustomerValidateChangePasswordWithTokenUseCase interface {
		ValidateChangePasswordWithToken(context.Context, *model.CustomerChangePasswordWithTokenInput) (*model.CustomerChangePasswordWithTokenInput, error)
	}
	ICustomerChangePasswordWithTokenUseCase interface {
		ChangePasswordWithToken(context.Context, *model.CustomerChangePasswordWithTokenInput) error
	}
	ICustomerValidateForgetPasswordUsecase interface {
		ValidateForgetPassword(context.Context, *model.CustomerForgetPasswordInput) (*model.CustomerForgetPasswordInput, error)
	}
	ICustomerForgetPasswordUseCase interface {
		ForgetPassword(context.Context, *model.CustomerForgetPasswordInput) (*model.CustomerForgetPasswordResp, error)
	}
	ICustomerAuthUseCase interface {
		iAuthenticationUseCase[*model.CustomerLoginInput, *model.Customer]
		ICustomerForgetPasswordUseCase
		ICustomerValidateForgetPasswordUsecase
		ICustomerChangePasswordWithTokenUseCase
		ICustomerValidateChangePasswordWithTokenUseCase
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
	ICustomerBankAccountGetFirstMineUseCase interface {
		GetFirstMine(context.Context, *model.BankAccountOrderInput, *model.BankAccountWhereInput) (*model.BankAccount, error)
	}
	ICustomerBankAccountListMineUseCase interface {
		ListMine(context.Context, *int, *int, *model.BankAccountOrderInput, *model.BankAccountWhereInput) ([]*model.BankAccount, error)
	}
	ICustomerBankAccountUseCase interface {
		ICustomerGetUserUseCase
		ICustomerConfigUseCase
		ICustomerBankAccountUpdateUseCase
		ICustomerBankAccountValidateUpdateInputUseCase
		ICustomerBankAccountListUseCase
		ICustomerBankAccountGetFirstMineUseCase
		ICustomerBankAccountListMineUseCase
		ICustomerBankAccountGetFirstUseCase
		IIsNextUseCase[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
	}
	ICustomerTransactionValidateConfirmInputUseCase interface {
		ValidateConfirmInput(context.Context, *model.Transaction, *model.TransactionConfirmUseCaseInput) error
	}
	ICustomerTransactionConfirmSuccessUseCase interface {
		ConfirmSuccess(context.Context, *model.Transaction, *string) (*model.Transaction, error)
	}
	ICustomerTransactionCreateUseCase interface {
		Create(context.Context, *model.TransactionCreateUseCaseInput) (*model.TransactionCreateResp, error)
	}
	ICustomerTransactionValidateCreateInputUseCase interface {
		iValidateCreateInput[*model.TransactionCreateUseCaseInput]
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
		IIsNextUseCase[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
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
	ICustomerDebtUpdateUseCase interface {
		iUpdateUseCase[*model.Debt, *model.DebtUpdateInput]
	}
	ICustomerDebtValidateCancelUseCase interface {
		ValidateCancel(context.Context, *model.Debt, *model.DebtUpdateInput) (*model.DebtUpdateInput, error)
	}
	ICustomerDebtCancelUseCase interface {
		Cancel(context.Context, *model.Debt, *model.DebtUpdateInput) (*model.Debt, error)
	}
	ICustomerDebtValidateFulfillUseCase interface {
		ValidateFulfill(context.Context, *model.Debt) error
	}
	ICustomerDebtValidateFulfillWithTokenUseCase interface {
		ValidateFulfillWithToken(context.Context, *model.Debt, *model.DebtFulfillWithTokenInput) (*model.DebtFulfillWithTokenInput, error)
	}
	ICustomerDebtFulfillUseCase interface {
		Fulfill(context.Context, *model.Debt) (*model.DebtFulfillResp, error)
	}
	ICustomerDebtFulfillWithTokenUseCase interface {
		FulfillWithToken(context.Context, *model.Debt, *model.DebtFulfillWithTokenInput) (*model.Debt, error)
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
		ICustomerDebtValidateCancelUseCase
		ICustomerDebtCancelUseCase
		ICustomerDebtValidateFulfillUseCase
		ICustomerDebtFulfillUseCase
		ICustomerDebtValidateFulfillWithTokenUseCase
		ICustomerDebtFulfillWithTokenUseCase
		IIsNextUseCase[*model.Debt, *model.DebtOrderInput, *model.DebtWhereInput]
	}
	// stream
	ICustomerStreamUseCase interface {
		ICustomerGetUserUseCase
		ICustomerConfigUseCase
	}
	// contact
	ICustomerContactListUseCase interface {
		iListUseCase[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput]
	}
	ICustomerContactListMineUseCase interface {
		ListMine(context.Context, *int, *int, *model.ContactOrderInput, *model.ContactWhereInput) ([]*model.Contact, error)
	}
	ICustomerContactCreateUseCase interface {
		iCreateUseCase[*model.Contact, *model.ContactCreateInput]
	}
	ICustomerContactValidateUpdateInputUseCase interface {
		iValidateUpdateInput[*model.Contact, *model.ContactUpdateInput]
	}
	ICustomerContactUpdateUseCase interface {
		iUpdateUseCase[*model.Contact, *model.ContactUpdateInput]
	}
	ICustomerContactDeleteUseCase interface {
		iDeleteUseCase[*model.Contact]
	}
	ICustomerContactValidateCreateInputUseCase interface {
		iValidateCreateInput[*model.ContactCreateInput]
	}
	ICustomerContactGetFirstMineUseCase interface {
		GetFirstMine(context.Context, *model.ContactOrderInput, *model.ContactWhereInput) (*model.Contact, error)
	}
	ICustomerContactUseCase interface {
		ICustomerGetUserUseCase
		ICustomerConfigUseCase
		ICustomerContactListUseCase
		ICustomerContactListMineUseCase
		ICustomerContactCreateUseCase
		ICustomerContactValidateCreateInputUseCase
		ICustomerContactGetFirstMineUseCase
		ICustomerContactUpdateUseCase
		ICustomerContactValidateUpdateInputUseCase
		ICustomerContactDeleteUseCase
		IIsNextUseCase[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput]
	}
)

type (
	IEmployeeListUseCase interface {
		iListUseCase[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput]
	}
	IEmployeeGetFirstUseCase interface {
		GetFirst(context.Context, *model.EmployeeOrderInput, *model.EmployeeWhereInput) (*model.Employee, error)
	}
	IEmployeeConfigUseCase interface {
		IGetConfigUseCase
	}
	IEmployeeGetUserUseCase interface {
		IGetUserUseCase
	}
	IEmployeeCustomerValidateCreateUseCase interface {
		iValidateCreateInput[*model.CustomerCreateInput]
	}
	IEmployeeCustomerCreateUseCase interface {
		iCreateUseCase[*model.Customer, *model.CustomerCreateInput]
	}
	IEmployeeCustomerUseCase interface {
		IEmployeeConfigUseCase
		IEmployeeGetUserUseCase
		IEmployeeCustomerCreateUseCase
		IEmployeeCustomerValidateCreateUseCase
	}
)
