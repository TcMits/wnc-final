package repository

import (
	"context"

	"github.com/Pallinder/go-randomdata"
	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/TcMits/wnc-final/pkg/tool/transaction"
)

type (
	CustomerCreateRepository struct {
		client *ent.Client
	}
	customerCreateRepository struct {
		bALR ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
		bACR CreateModelRepository[*model.BankAccount, *model.BankAccountCreateInput]
		repo CreateModelRepository[*model.Customer, *model.CustomerCreateInput]
	}
)

func newCustomerCreateRepository(client *ent.Client) *customerCreateRepository {
	return &customerCreateRepository{
		bALR: GetBankAccountListRepository(client),
		bACR: GetBankAccountCreateRepository(client),
		repo: ent.NewCustomerCreateRepository(client, false),
	}
}

func GetCustomerListRepository(
	client *ent.Client,
) ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput] {
	return ent.NewCustomerReadRepository(client)
}

func GetCustomerCreateRepository(
	client *ent.Client,
) CreateModelRepository[*model.Customer, *model.CustomerCreateInput] {
	return &CustomerCreateRepository{client: client}
}

func GetCustomerDeleteRepository(
	client *ent.Client,
) DeleteModelRepository[*model.Customer] {
	return ent.NewCustomerDeleteRepository(client, false)
}

func GetCustomerUpdateRepository(
	client *ent.Client,
) UpdateModelRepository[*model.Customer, *model.CustomerUpdateInput] {
	return ent.NewCustomerUpdateRepository(client, false)
}

func (s *customerCreateRepository) isExist(ctx context.Context, accountNumber string) (bool, error) {
	entites, err := s.bALR.List(ctx, generic.GetPointer(1), generic.GetPointer(0), nil, &model.BankAccountWhereInput{
		AccountNumber: &accountNumber,
	})
	return len(entites) > 0, err
}

func (s *customerCreateRepository) getCandidateAccountNumber(ctx context.Context) (string, error) {
	ac := randomdata.Digits(16)
	es, err := s.isExist(ctx, ac)
	if err != nil {
		return "", err
	}
	for es {
		ac = randomdata.Digits(16)
		es, err = s.isExist(ctx, ac)
		if err != nil {
			return "", err
		}
	}
	return ac, nil
}

func (s *customerCreateRepository) create(ctx context.Context, i *model.CustomerCreateInput) (*model.Customer, error) {
	e, err := s.repo.Create(ctx, i)
	if err != nil {
		return nil, err
	}
	accountNumberCandidate, err := s.getCandidateAccountNumber(ctx)
	if err != nil {
		return nil, err
	}
	_, err = s.bACR.Create(ctx, &model.BankAccountCreateInput{
		CustomerID:    e.ID,
		IsForPayment:  generic.GetPointer(true),
		AccountNumber: &accountNumberCandidate,
	})
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (s *CustomerCreateRepository) Create(ctx context.Context, i *model.CustomerCreateInput) (*model.Customer, error) {
	var e *model.Customer
	if err := transaction.WithTx(ctx, s.client, func(tx *ent.Tx) error {
		var errInner error
		driver := newCustomerCreateRepository(tx.Client())
		e, errInner = driver.create(ctx, i)
		return errInner
	}); err != nil {
		return nil, err
	}
	return e, nil
}
