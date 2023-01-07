package repository

import (
	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

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
	return ent.NewBankAccountUpdateRepository(client, false)
}

func GetBankAccountIsNextRepository(
	client *ent.Client,
) IIsNextModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput] {
	return getIsNextModelRepostiory(GetBankAccountListRepository(client))
}
