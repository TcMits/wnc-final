package contact

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
)

func (s *CustomerContactListUseCase) List(ctx context.Context, limit, offset *int, o *model.ContactOrderInput, w *model.ContactWhereInput) ([]*model.Contact, error) {
	entites, err := s.repoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.contact.implementations.CustomerContactListUseCase.List: %s", err))
	}
	return entites, nil
}
func (s *CustomerContactDeleteUseCase) Delete(ctx context.Context, e *model.Contact) error {
	err := s.repoDelete.Delete(ctx, e)
	if err != nil {
		return usecase.WrapError(fmt.Errorf("internal.usecase.contact.implementations.CustomerContactDeleteUseCase.Delete: %s", err))
	}
	return nil
}

func (s *CustomerContactListMineUseCase) ListMine(ctx context.Context, limit, offset *int, o *model.ContactOrderInput, w *model.ContactWhereInput) ([]*model.Contact, error) {
	if w == nil {
		w = new(model.ContactWhereInput)
	}
	user := usecase.GetUserAsCustomer(ctx)
	w.OwnerID = generic.GetPointer(user.ID)
	return s.cLUC.List(ctx, limit, offset, o, w)
}

func (s *CustomerContactGetFirstMineUseCase) GetFirstMine(ctx context.Context, o *model.ContactOrderInput, w *model.ContactWhereInput) (*model.Contact, error) {
	entites, err := s.cGFUC.ListMine(ctx, generic.GetPointer(1), generic.GetPointer(0), o, w)
	if err != nil {
		return nil, err
	}
	if len(entites) > 0 {
		return entites[0], nil
	}
	return nil, nil
}
func (s *CustomerContactCreateUseCase) Create(ctx context.Context, i *model.ContactCreateInput) (*model.Contact, error) {
	e, err := s.repoCreate.Create(ctx, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.contact.implementations.CustomerContactCreateUseCase.Create: %s", err))
	}
	return e, nil
}
func (s *CustomerContactUpdateUseCase) Update(ctx context.Context, e *model.Contact, i *model.ContactUpdateInput) (*model.Contact, error) {
	e, err := s.repoUpdate.Update(ctx, e, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.contact.implementations.CustomerContactUpdateUseCase.Update: %s", err))
	}
	return e, nil
}
func (s *CustomerContactValidateCreateInputUseCase) ValidateCreate(ctx context.Context, i *model.ContactCreateInput) (*model.ContactCreateInput, error) {
	if i == nil {
		return nil, usecase.WrapError(fmt.Errorf("invalid input"))
	}
	if i.BankName == "" {
		i.BankName = *s.cfUC.GetProductOwnerName()
	}
	entites, err := s.cGFUC.ListMine(ctx, generic.GetPointer(1), generic.GetPointer(0), nil, &model.ContactWhereInput{
		AccountNumber: &i.AccountNumber,
		BankName:      &i.BankName,
	})
	if err != nil {
		return nil, err
	}
	if len(entites) > 0 {
		return nil, usecase.WrapError(fmt.Errorf("the account number of the bank already existed"))
	}
	user := usecase.GetUserAsCustomer(ctx)
	i.OwnerID = user.ID
	return i, nil
}
func (s *CustomerContactValidateUpdateInputUseCase) ValidateUpdate(ctx context.Context, e *model.Contact, i *model.ContactUpdateInput) (*model.ContactUpdateInput, error) {
	if i == nil {
		return nil, usecase.WrapError(fmt.Errorf("invalid input"))
	}
	if i.BankName == nil || *i.BankName == "" {
		i.BankName = &e.BankName
	}
	if i.AccountNumber == nil || *i.AccountNumber == "" {
		i.AccountNumber = &e.AccountNumber
	}
	entites, err := s.cGFUC.ListMine(ctx, generic.GetPointer(1), generic.GetPointer(0), nil, &model.ContactWhereInput{
		AccountNumber: i.AccountNumber,
		BankName:      i.BankName,
		IDNEQ:         generic.GetPointer(e.ID),
	})
	if err != nil {
		return nil, err
	}
	if len(entites) > 0 {
		return nil, usecase.WrapError(fmt.Errorf("the account number of the bank already existed"))
	}
	i.OwnerID = nil
	return i, nil
}

func (s *CustomerContactIsNextUseCase) IsNext(ctx context.Context, limit, offset int, o *model.ContactOrderInput, w *model.ContactWhereInput) (bool, error) {
	user := usecase.GetUserAsCustomer(ctx)
	if w == nil {
		w = new(model.ContactWhereInput)
	}
	w.OwnerID = generic.GetPointer(user.ID)
	return s.iNUC.IsNext(ctx, limit, offset, o, w)
}
