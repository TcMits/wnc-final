package bankaccount

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
)

func (s *CustomerBankAccountDeleteUseCase) Delete(ctx context.Context, e *model.BankAccount) error {
	err := s.Repo.Delete(ctx, e)
	if err != nil {
		return usecase.WrapError(fmt.Errorf("internal.usecase.logic.bankaccount.implementations.CustomerBankAccountDeleteUseCase.Delete: %s", err))
	}
	return nil
}

func (s *CustomerTPBankBankAccountGetUseCase) Get(ctx context.Context, w *model.BankAccountWhereInput) (*model.BankAccountPartner, error) {
	e, err := s.W1.Get(ctx, &model.WhereInputPartner{
		AccountNumber: *w.AccountNumber,
	})
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.bankaccount.implementations.CustomerTPBankBankAccountGetUseCase.Get: %s", err))
	}
	return e, nil
}

func (s *CustomerBankAccountUpdateUseCase) Update(ctx context.Context, m *model.BankAccount, i *model.BankAccountUpdateInput) (*model.BankAccount, error) {
	m, err := s.RepoUpdate.Update(ctx, m, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.bankaccount.implementations.CustomerBankAccountUpdateUseCase.Update: %s", err))
	}
	return m, nil
}

func (s *CustomerBankAccountListUseCase) List(ctx context.Context, limit, offset *int, o *model.BankAccountOrderInput, w *model.BankAccountWhereInput) ([]*model.BankAccount, error) {
	entites, err := s.RepoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.bankaccount.implementations.CustomerBankAccountListUseCase.List: %s", err))
	}
	return entites, nil
}

func (s *CustomerBankAccountValidateUpdateInputUseCase) ValidateUpdate(ctx context.Context, m *model.BankAccount, i *model.BankAccountUpdateInput) (*model.BankAccountUpdateInput, error) {
	user := usecase.GetUserAsCustomer(ctx)
	if user.ID != m.CustomerID {
		return nil, usecase.ValidationError(fmt.Errorf("the bank account is not owned by user"))
	}
	if *i.IsForPayment {
		e, err := s.UC1.GetFirstMine(ctx, nil, &model.BankAccountWhereInput{
			IsForPayment: generic.GetPointer(true),
		})
		if err != nil {
			return nil, err
		}
		if e != nil {
			return nil, usecase.ValidationError(fmt.Errorf("payment account already existed"))
		}
	}
	return i, nil
}

func (s *CustomerBankAccountGetFirstUseCase) GetFirst(ctx context.Context, o *model.BankAccountOrderInput, w *model.BankAccountWhereInput) (*model.BankAccount, error) {
	l, of := 1, 0
	entities, err := s.UC1.List(ctx, &l, &of, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.bankaccount.GetFirst: %s", err))
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}

func (s *CustomerBankAccountListMineUseCase) ListMine(ctx context.Context, limit, offset *int, o *model.BankAccountOrderInput, w *model.BankAccountWhereInput) ([]*model.BankAccount, error) {
	user := usecase.GetUserAsCustomer(ctx)
	if w == nil {
		w = new(model.BankAccountWhereInput)
	}
	w.CustomerID = generic.GetPointer(user.ID)
	return s.UC1.List(ctx, limit, offset, o, w)
}

func (s *CustomerBankAccountGetFirstMineUseCase) GetFirstMine(ctx context.Context, o *model.BankAccountOrderInput, w *model.BankAccountWhereInput) (*model.BankAccount, error) {
	l, of := 1, 0
	entities, err := s.UC1.ListMine(ctx, &l, &of, o, w)
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
	return s.UC1.IsNext(ctx, limit, offset, o, w)
}

func (s *EmployeeBankAccountDeleteUseCase) Delete(ctx context.Context, e *model.BankAccount) error {
	err := s.Repo.Delete(ctx, e)
	if err != nil {
		return usecase.WrapError(fmt.Errorf("internal.usecase.logic.bankaccount.implementations.EmployeeBankAccountDeleteUseCase.Delete: %s", err))
	}
	return nil
}

func (s *EmployeeBankAccountValidateUpdateInputUseCase) ValidateUpdate(ctx context.Context, m *model.BankAccount, i *model.BankAccountUpdateInput) (*model.BankAccountUpdateInput, error) {
	if !m.IsForPayment {
		return nil, usecase.ValidationError(fmt.Errorf("bank account not for payment"))
	}
	return i, nil
}

func (s *EmployeeBankAccountGetFirstUseCase) GetFirst(ctx context.Context, o *model.BankAccountOrderInput, w *model.BankAccountWhereInput) (*model.BankAccount, error) {
	l, of := 1, 0
	entities, err := s.UC1.List(ctx, &l, &of, o, w)
	if err != nil {
		return nil, err
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}

func (s *EmployeeBankAccountListUseCase) List(ctx context.Context, limit, offset *int, o *model.BankAccountOrderInput, w *model.BankAccountWhereInput) ([]*model.BankAccount, error) {
	entites, err := s.RepoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.bankaccount.implementations.EmployeeBankAccountListUseCase.List: %s", err))
	}
	return entites, nil
}
func (s *EmployeeBankAccountUpdateUseCase) Update(ctx context.Context, m *model.BankAccount, i *model.BankAccountUpdateInput) (*model.BankAccount, error) {
	m, err := s.RepoUpdate.Update(ctx, m, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.bankaccount.implementations.CustomerBankAccountUpdateUseCase.Update: %s", err))
	}
	return m, nil
}

// partner
func (s *PartnerBankAccountListUseCase) List(ctx context.Context, limit, offset *int, o *model.BankAccountOrderInput, w *model.BankAccountWhereInput) ([]*model.BankAccount, error) {
	entites, err := s.RepoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.bankaccount.implementations.PartnerBankAccountListUseCase.List: %s", err))
	}
	return entites, nil
}
func (s *PartnerBankAccountGetFirstUseCase) GetFirst(ctx context.Context, o *model.BankAccountOrderInput, w *model.BankAccountWhereInput) (*model.BankAccount, error) {
	l, of := 1, 0
	entities, err := s.UC1.List(ctx, &l, &of, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.bankaccount.GetFirst: %s", err))
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}

func (s *PartnerBankAccountRespGetFirstUseCase) GetFirst(ctx context.Context, o *model.BankAccountOrderInput, w *model.BankAccountWhereInput) (*model.PartnerBankAccountResp, error) {
	l, of := 1, 0
	entities, err := s.UC1.List(ctx, &l, &of, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.bankaccount.GetFirst: %s", err))
	}
	if len(entities) > 0 {
		e := entities[0]
		c, err := s.UC2.GetFirst(ctx, nil, &model.CustomerWhereInput{ID: generic.GetPointer(e.CustomerID)})
		if err != nil {
			return nil, err
		}
		r := &model.PartnerBankAccountResp{
			AccountNumber: e.AccountNumber,
		}
		r.Name = "Not valid"
		if c != nil {
			r.Name = c.GetName()
		}
		return r, nil
	}
	return nil, nil
}
