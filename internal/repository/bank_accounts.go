package repository

import (
	"context"

	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
)

type bankAccountUpdateRepository struct {
	repo UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput]
}

func (s *bankAccountUpdateRepository) Update(ctx context.Context, e *model.BankAccount, i *model.BankAccountUpdateInput) (*model.BankAccount, error) {
	if i == nil {
		return e, nil
	}
	if i.CashIn == nil {
		i.CashIn = generic.GetPointer(float64(0))
	}
	if i.CashOut == nil {
		i.CashOut = generic.GetPointer(float64(0))
	}
	i.CashIn = generic.GetPointer(e.CashIn + *i.CashIn)
	i.CashOut = generic.GetPointer(e.CashOut + *i.CashOut)
	return s.repo.Update(ctx, e, i)
}

func GetBankAccountListRepository(
	client *ent.Client,
) ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput] {
	return ent.NewBankAccountReadRepository(client)
}

func GetBankAccountCreateRepository(
	client *ent.Client,
) CreateModelRepository[*model.BankAccount, *model.BankAccountCreateInput] {
	return ent.NewBankAccountCreateRepository(client, false)
}

func GetBankAccountDeleteRepository(
	client *ent.Client,
) DeleteModelRepository[*model.BankAccount] {
	return ent.NewBankAccountDeleteRepository(client, false)
}

func GetBankAccountUpdateRepository(
	client *ent.Client,
) UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput] {
	return &bankAccountUpdateRepository{repo: ent.NewBankAccountUpdateRepository(client, false)}
}

func GetBankAccountIsNextRepository(
	client *ent.Client,
) IIsNextModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput] {
	return getIsNextModelRepostiory(GetBankAccountListRepository(client))
}
