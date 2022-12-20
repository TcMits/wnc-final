package transaction

import (
	"context"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/bankaccount"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/internal/usecase/me"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type (
	CustomerTransactionCreateUseCase struct {
		repoCreate repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput]
	}
	CustomerTransactionValidateCreateInputUseCase struct {
		repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
	}
	CustomerTransactionListUseCase struct {
		repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
	}
	CustomerTransactionBankAccountListUseCase struct {
		usecase.ICustomerBankAccountListUseCase
	}
	CustomerTransactionUseCase struct {
		*CustomerTransactionBankAccountListUseCase
		usecase.ICustomerTransactionValidateCreateInputUseCase
		usecase.ICustomerTransactionCreateUseCase
		usecase.ICustomerTransactionListUseCase
		usecase.ICustomerConfigUseCase
		usecase.ICustomerGetUserUseCase
	}
)

func NewCustomerTransactionValidateCreateInputUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.ICustomerTransactionValidateCreateInputUseCase {
	return &CustomerTransactionValidateCreateInputUseCase{
		repoList: repoList,
	}
}

func NewCustomerTransactionListUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.ICustomerTransactionListUseCase {
	return &CustomerTransactionListUseCase{
		repoList: repoList,
	}
}

func NewCustomerTransactionCreateUseCase(
	repoCreate repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput],
) usecase.ICustomerTransactionCreateUseCase {
	return &CustomerTransactionCreateUseCase{
		repoCreate: repoCreate,
	}
}

func NewCustomerTransactionBankAccountListUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) *CustomerTransactionBankAccountListUseCase {
	return &CustomerTransactionBankAccountListUseCase{
		ICustomerBankAccountListUseCase: bankaccount.NewCustomerBankAccountListUseCase(repoList),
	}
}

func NewCustomerTransactionUseCase(
	repoCreate repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput],
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	rlba repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	sk *string,
) usecase.ICustomerTransactionUseCase {
	return &CustomerTransactionUseCase{
		ICustomerTransactionCreateUseCase:              NewCustomerTransactionCreateUseCase(repoCreate),
		ICustomerTransactionValidateCreateInputUseCase: NewCustomerTransactionValidateCreateInputUseCase(repoList),
		ICustomerTransactionListUseCase:                NewCustomerTransactionListUseCase(repoList),
		CustomerTransactionBankAccountListUseCase:      NewCustomerTransactionBankAccountListUseCase(rlba),
		ICustomerConfigUseCase:                         config.NewCustomerConfigUseCase(sk),
		ICustomerGetUserUseCase:                        me.NewCustomerGetUserUseCase(rlc),
	}
}

func (uc *CustomerTransactionCreateUseCase) Create(ctx context.Context, i *model.TransactionCreateInput) (*model.Transaction, error) {
	return uc.repoCreate.Create(ctx, i)
}

func (uc *CustomerTransactionValidateCreateInputUseCase) Validate(ctx context.Context, i *model.TransactionCreateInput) (*model.TransactionCreateInput, error) {
	return i, nil
}

func (uc *CustomerTransactionListUseCase) List(ctx context.Context, limit, offset *int, o *model.TransactionOrderInput, w *model.TransactionWhereInput) ([]*model.Transaction, error) {
	return uc.repoList.List(ctx, limit, offset, o, w)
}

func (uc *CustomerTransactionBankAccountListUseCase) ListBankAccounts(ctx context.Context, limit, offset *int, o *model.BankAccountOrderInput, w *model.BankAccountWhereInput) ([]*model.BankAccount, error) {
	return uc.List(ctx, limit, offset, o, w)
}
