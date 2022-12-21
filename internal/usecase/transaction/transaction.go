package transaction

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/ent/transaction"
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
	CustomerTransactionListUseCase struct {
		repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
	}
	CustomerTransactionBankAccountListUseCase struct {
		usecase.ICustomerBankAccountListUseCase
	}
	CustomerTransactionValidateCreateInputUseCase struct {
		*CustomerTransactionBankAccountListUseCase
		repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
	}
	CustomerTransactionUseCase struct {
		usecase.ICustomerTransactionValidateCreateInputUseCase
		usecase.ICustomerTransactionCreateUseCase
		usecase.ICustomerTransactionListUseCase
		usecase.ICustomerConfigUseCase
		usecase.ICustomerGetUserUseCase
	}
)

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

func NewCustomerTransactionValidateCreateInputUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
	rlba repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.ICustomerTransactionValidateCreateInputUseCase {
	return &CustomerTransactionValidateCreateInputUseCase{
		repoList: repoList,
		CustomerTransactionBankAccountListUseCase: NewCustomerTransactionBankAccountListUseCase(rlba),
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
		ICustomerTransactionValidateCreateInputUseCase: NewCustomerTransactionValidateCreateInputUseCase(repoList, rlba),
		ICustomerTransactionListUseCase:                NewCustomerTransactionListUseCase(repoList),
		ICustomerConfigUseCase:                         config.NewCustomerConfigUseCase(sk),
		ICustomerGetUserUseCase:                        me.NewCustomerGetUserUseCase(rlc),
	}
}

func (uc *CustomerTransactionCreateUseCase) Create(ctx context.Context, i *model.TransactionCreateInput) (*model.Transaction, error) {
	return uc.repoCreate.Create(ctx, i)
}

func (uc *CustomerTransactionValidateCreateInputUseCase) Validate(ctx context.Context, i *model.TransactionCreateInput) (*model.TransactionCreateInput, error) {
	if i.TransactionType == transaction.TransactionTypeInternal {
		l, o := 1, 0
		entities, err := uc.List(ctx, &l, &o, nil, &model.BankAccountWhereInput{
			ID: &i.SenderID,
		})
		if err != nil {
			return nil, err
		}
		if len(entities) > 0 {
			ba := entities[0]
			if ba.IsForPayment {
				balance := ba.Balance()
			} else {
				return nil, usecase.WrapError(fmt.Errorf("this bank account is not for payment"))
			}
		} else {
			return nil, usecase.WrapError(fmt.Errorf("this bank account is invalid"))
		}
	}
	return i, nil
}

func (uc *CustomerTransactionListUseCase) List(ctx context.Context, limit, offset *int, o *model.TransactionOrderInput, w *model.TransactionWhereInput) ([]*model.Transaction, error) {
	return uc.repoList.List(ctx, limit, offset, o, w)
}
