package bankaccount

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

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
		entities, err := uc.bALUC.List(ctx, &l, &o, nil, &model.BankAccountWhereInput{
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

func (uc *CustomerBankAccountGetFirstUseCase) GetFirst(ctx context.Context, o *model.BankAccountOrderInput, w *model.BankAccountWhereInput) (*model.BankAccount, error) {
	l, of := 1, 0
	entities, err := uc.bALUC.List(ctx, &l, &of, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.bankaccount.GetFirst: %s", err))
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}
