package bankaccount

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
)

func (uc *CustomerBankAccountUpdateUseCase) Update(ctx context.Context, m *model.BankAccount, i *model.BankAccountUpdateInput) (*model.BankAccount, error) {
	m, err := uc.repoUpdate.Update(ctx, m, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.bankaccount.implementations.CustomerBankAccountUpdateUseCase.Update: %s", err))
	}
	return m, nil
}

func (uc *CustomerBankAccountListUseCase) List(ctx context.Context, limit, offset *int, o *model.BankAccountOrderInput, w *model.BankAccountWhereInput) ([]*model.BankAccount, error) {
	entites, err := uc.repoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.bankaccount.implementations.CustomerBankAccountListUseCase.List: %s", err))
	}
	return entites, nil
}

func (uc *CustomerBankAccountValidateUpdateInputUseCase) ValidateUpdate(ctx context.Context, m *model.BankAccount, i *model.BankAccountUpdateInput) (*model.BankAccountUpdateInput, error) {
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

func (uc *CustomerBankAccountListMineUseCase) ListMine(ctx context.Context, limit, offset *int, o *model.BankAccountOrderInput, w *model.BankAccountWhereInput) ([]*model.BankAccount, error) {
	user := usecase.GetUserAsCustomer(ctx)
	if w == nil {
		w = new(model.BankAccountWhereInput)
	}
	w.CustomerID = generic.GetPointer(user.ID)
	return uc.bALUC.List(ctx, limit, offset, o, w)
}

func (uc *CustomerBankAccountGetFirstMineUseCase) GetFirstMine(ctx context.Context, o *model.BankAccountOrderInput, w *model.BankAccountWhereInput) (*model.BankAccount, error) {
	l, of := 1, 0
	entities, err := uc.bALMUC.ListMine(ctx, &l, &of, o, w)
	if err != nil {
		return nil, err
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}

func (s *CustomerBankAccountIsNextUseCase) IsNext(ctx context.Context, limit, offset int, o *model.BankAccountOrderInput, w *model.BankAccountWhereInput) (bool, error) {
	user := usecase.GetUserAsCustomer(ctx)
	if w == nil {
		w = new(model.BankAccountWhereInput)
	}
	w.CustomerID = generic.GetPointer(user.ID)
	return s.iNUC.IsNext(ctx, limit, offset, o, w)
}

func (s *EmployeeBankAccountValidateUpdateInputUseCase) ValidateUpdate(ctx context.Context, m *model.BankAccount, i *model.BankAccountUpdateInput) (*model.BankAccountUpdateInput, error) {
	if !m.IsForPayment {
		return nil, usecase.ValidationError(fmt.Errorf("bank account not for payment"))
	}
	return i, nil
}

func (s *EmployeeBankAccountGetFirstUseCase) GetFirst(ctx context.Context, o *model.BankAccountOrderInput, w *model.BankAccountWhereInput) (*model.BankAccount, error) {
	l, of := 1, 0
	entities, err := s.bALUC.List(ctx, &l, &of, o, w)
	if err != nil {
		return nil, err
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}

func (s *EmployeeBankAccountListUseCase) List(ctx context.Context, limit, offset *int, o *model.BankAccountOrderInput, w *model.BankAccountWhereInput) ([]*model.BankAccount, error) {
	entites, err := s.repoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.bankaccount.implementations.EmployeeBankAccountListUseCase.List: %s", err))
	}
	return entites, nil
}
func (s *EmployeeBankAccountUpdateUseCase) Update(ctx context.Context, m *model.BankAccount, i *model.BankAccountUpdateInput) (*model.BankAccount, error) {
	m, err := s.repoUpdate.Update(ctx, m, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.bankaccount.implementations.CustomerBankAccountUpdateUseCase.Update: %s", err))
	}
	return m, nil
}
