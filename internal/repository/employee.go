package repository

import (
	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

func GetEmployeeListRepository(
	client *ent.Client,
) ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput] {
	return ent.NewEmployeeReadRepository(client)
}

func GetEmployeeCreateRepository(
	client *ent.Client,
) CreateModelRepository[*model.Employee, *model.EmployeeCreateInput] {
	return ent.NewEmployeeCreateRepository(client, false)
}

func GetEmployeeUpdateRepository(
	client *ent.Client,
) UpdateModelRepository[*model.Employee, *model.EmployeeUpdateInput] {
	return ent.NewEmployeeUpdateRepository(client, false)
}

func GetEmployeeDeleteRepository(
	client *ent.Client,
) DeleteModelRepository[*model.Employee] {
	return ent.NewEmployeeDeleteRepository(client, false)
}
