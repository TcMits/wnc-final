package repository

import (
	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

func GetAdminListRepository(
	client *ent.Client,
) ListModelRepository[*model.Admin, *model.AdminOrderInput, *model.AdminWhereInput] {
	return ent.NewAdminReadRepository(client)
}

func GetAdminCreateRepository(
	client *ent.Client,
) CreateModelRepository[*model.Admin, *model.AdminCreateInput] {
	return ent.NewAdminCreateRepository(client, false)
}

func GetAdminUpdateRepository(
	client *ent.Client,
) UpdateModelRepository[*model.Admin, *model.AdminUpdateInput] {
	return ent.NewAdminUpdateRepository(client, false)
}

func GetAdminDeleteRepository(
	client *ent.Client,
) DeleteModelRepository[*model.Admin] {
	return ent.NewAdminDeleteRepository(client, false)
}
