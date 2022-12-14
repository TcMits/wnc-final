package repository

import (
	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

func GetCustomerListRepository(
	client *ent.Client,
) ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput] {
	return ent.NewCustomerReadRepository(client)
}

func GetCustomerCreateRepository(
	client *ent.Client,
) CreateModelRepository[*model.Customer, *model.CustomerCreateInput] {
	return ent.NewCustomerCreateRepository(client, false)
}

func GetCustomerDeleteRepository(
	client *ent.Client,
) DeleteModelRepository[*model.Customer] {
	return ent.NewCustomerDeleteRepository(client, false)
}
