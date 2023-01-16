package constructors

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/logic/employee"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

func NewEmployeeGetUserUseCase(
	repoList repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
) usecase.IEmployeeGetUserUseCase {
	uc := &employee.EmployeeGetUserUseCase{
		UC1: NewEmployeeGetFirstUseCase(repoList),
	}
	return uc
}

func NewEmployeeGetFirstUseCase(
	repoList repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
) usecase.IEmployeeGetFirstUseCase {
	uc := &employee.EmployeeGetFirstUseCase{
		UC1: repoList,
	}
	return uc
}

func NewEmployeeListUseCase(
	repoList repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
) usecase.IEmployeeListUseCase {
	uc := &employee.EmployeeListUseCase{
		RepoList: repoList,
	}
	return uc
}
func NewEmployeeUpdateUseCase(
	repoUpdate repository.UpdateModelRepository[*model.Employee, *model.EmployeeUpdateInput],
) usecase.IEmployeeUpdateUseCase {
	return &employee.EmployeeUpdateUseCase{
		RepoUpdate: repoUpdate,
	}
}

func NewAdminEmployeeListUseCase(
	r1 repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
) usecase.IAdminEmployeeListUseCase {
	return &employee.AdminEmployeeListUseCase{
		R: r1,
	}
}
func NewAdminEmployeeGetFirstUseCase(
	r1 repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
) usecase.IAdminEmployeeGetFirstUseCase {
	return &employee.AdminEmployeeGetFirstUseCase{
		UC1: NewAdminEmployeeListUseCase(r1),
	}
}

func NewAdminEmployeeUseCase(
	r1 repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
	r2 repository.CreateModelRepository[*model.Employee, *model.EmployeeCreateInput],
	r3 repository.UpdateModelRepository[*model.Employee, *model.EmployeeUpdateInput],
	r4 repository.DeleteModelRepository[*model.Employee],
	r5 repository.IIsNextModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
	r6 repository.ListModelRepository[*model.Admin, *model.AdminOrderInput, *model.AdminWhereInput],
	secretKey,
	prodOwnerName *string,
) usecase.IAdminEmployeeUseCase {
	return &employee.AdminEmployeeUseCase{
		IAdminConfigUseCase:                 NewAdminConfigUseCase(secretKey, prodOwnerName),
		IAdminGetUserUseCase:                NewAdminGetUserUseCase(r6),
		IAdminEmployeeListUseCase:           &employee.AdminEmployeeListUseCase{R: r1},
		IAdminEmployeeGetFirstUseCase:       &employee.AdminEmployeeGetFirstUseCase{UC1: NewAdminEmployeeListUseCase(r1)},
		IAdminEmployeeCreateUseCase:         &employee.AdminEmployeeCreateUseCase{R: r2},
		IAdminEmployeeValidateCreateUseCase: &employee.AdminEmployeeValidateCreateUseCase{UC1: NewAdminEmployeeGetFirstUseCase(r1)},
		IAdminEmployeeUpdateUseCase:         &employee.AdminEmployeeUpdateUseCase{R: r3},
		IAdminEmployeeValidateUpdateUseCase: &employee.AdminEmployeeValidateUpdateUseCase{UC1: NewAdminEmployeeGetFirstUseCase(r1)},
		IAdminEmployeeDeleteUseCase:         &employee.AdminEmployeeDeleteUseCase{R: r4},
		IIsNextUseCase:                      NewIsNextUseCase(r5),
	}
}
