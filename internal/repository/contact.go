package repository

import (
	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

func GetContactListRepository(
	client *ent.Client,
) ListModelRepository[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput] {
	return ent.NewContactReadRepository(client)
}

func GetContactCreateRepository(
	client *ent.Client,
) CreateModelRepository[*model.Contact, *model.ContactCreateInput] {
	return ent.NewContactCreateRepository(client, false)
}

func GetContactUpdateRepository(
	client *ent.Client,
) UpdateModelRepository[*model.Contact, *model.ContactUpdateInput] {
	return ent.NewContactUpdateRepository(client, false)
}

func GetContactDeleteRepository(
	client *ent.Client,
) DeleteModelRepository[*model.Contact] {
	return ent.NewContactDeleteRepository(client, false)
}

func GetContactIsNextRepository(
	client *ent.Client,
) IIsNextModelRepository[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput] {
	return getIsNextModelRepostiory(GetContactListRepository(client))
}
