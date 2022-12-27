package repository

import (
	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

func GetDebtListRepository(
	client *ent.Client,
) ListModelRepository[*model.Debt, *model.DebtOrderInput, *model.DebtWhereInput] {
	return ent.NewDebtReadRepository(client)
}

func GetDebtCreateRepository(
	client *ent.Client,
) CreateModelRepository[*model.Debt, *model.DebtCreateInput] {
	return ent.NewDebtCreateRepository(client, false)
}

func GetDebtUpdateRepository(
	client *ent.Client,
) UpdateModelRepository[*model.Debt, *model.DebtUpdateInput] {
	return ent.NewDebtUpdateRepository(client, false)
}
