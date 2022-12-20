package bankaccount

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/internal/usecase/me"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type (
	CustomerBankAccountUpdateUseCase struct {
		repoUpdate repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput]
	}
	CustomerBankAccountValidateUpdateInputUseCase struct {
		repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
	}
	CustomerBankAccountListUseCase struct {
		repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
	}
	CustomerBankAccountUseCase struct {
		usecase.ICustomerBankAccountUpdateUseCase
		usecase.ICustomerBankAccountValidateUpdateInputUseCase
		usecase.ICustomerBankAccountListUseCase
		usecase.ICustomerConfigUseCase
		usecase.ICustomerGetUserUseCase
	}
)

func NewCustomerBankAccountListUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.ICustomerBankAccountListUseCase {
	return &CustomerBankAccountListUseCase{
		repoList: repoList,
	}
}

func NewCustomerBankAccountValidateUpdateInputUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.ICustomerBankAccountValidateUpdateInputUseCase {
	return &CustomerBankAccountValidateUpdateInputUseCase{
		repoList: repoList,
	}
}

func NewCustomerBankAccountUpdateUseCase(
	repoUpdate repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput],
) usecase.ICustomerBankAccountUpdateUseCase {
	return &CustomerBankAccountUpdateUseCase{
		repoUpdate: repoUpdate,
	}
}

func NewCustomerBankAccountUseCase(
	repoUpdate repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput],
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	sk *string,
) usecase.ICustomerBankAccountUseCase {
	return &CustomerBankAccountUseCase{
		ICustomerBankAccountUpdateUseCase:              NewCustomerBankAccountUpdateUseCase(repoUpdate),
		ICustomerBankAccountValidateUpdateInputUseCase: NewCustomerBankAccountValidateUpdateInputUseCase(repoList),
		ICustomerBankAccountListUseCase:                NewCustomerBankAccountListUseCase(repoList),
		ICustomerConfigUseCase:                         config.NewCustomerConfigUseCase(sk),
		ICustomerGetUserUseCase:                        me.NewCustomerGetUserUseCase(rlc),
	}
}

func (uc *CustomerBankAccountUpdateUseCase) Update(ctx context.Context, m *model.BankAccount, i *model.BankAccountUpdateInput) (*model.BankAccount, error) {
	return uc.repoUpdate.Update(ctx, m, i)
}

func (uc *CustomerBankAccountListUseCase) List(ctx context.Context, limit, offset *int, o *model.BankAccountOrderInput, w *model.BankAccountWhereInput) ([]*model.BankAccount, error) {
	return uc.repoList.List(ctx, limit, offset, o, w)
}

func (uc *CustomerBankAccountValidateUpdateInputUseCase) Validate(ctx context.Context, m *model.BankAccount, i *model.BankAccountUpdateInput) (*model.BankAccountUpdateInput, error) {
	userAny := ctx.Value("user")
	if userAny == nil {
		return nil, usecase.WrapError(fmt.Errorf("user is invalid"))
	}
	user, ok := userAny.(*model.Customer)
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("user is invalid"))
	}
	if user.ID != m.CustomerID {
		return nil, usecase.WrapError(fmt.Errorf("the bank account is not owned by user"))
	}
	if *i.IsForPayment {
		l, o := 1, 0
		iFP := true
		entities, err := uc.repoList.List(ctx, &l, &o, nil, &model.BankAccountWhereInput{
			IsForPayment: &iFP,
		})
		if err != nil {
			return nil, err
		}
		if len(entities) > 0 {
			return nil, usecase.WrapError(fmt.Errorf("payment account already existed"))
		}
	}
	return i, nil
}
