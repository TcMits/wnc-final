package constructors

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/logic/customer"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

func NewCustomerGetUserUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.ICustomerGetUserUseCase {
	uc := &customer.CustomerGetUserUseCase{
		UC1: NewCustomerGetFirstUseCase(repoList),
	}
	return uc
}
func NewCustomerUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	sk *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerUseCase {
	return &customer.CustomerUseCase{
		ICustomerConfigUseCase:   NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		ICustomerGetUserUseCase:  NewCustomerGetUserUseCase(repoList),
		ICustomerListUseCase:     NewCustomerListUseCase(repoList),
		ICustomerGetFirstUseCase: NewCustomerGetFirstUseCase(repoList),
	}
}

func NewCustomerValidateCreateUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.IEmployeeCustomerValidateCreateUseCase {
	return &customer.CustomerValidateCreateUseCase{
		UC1: NewCustomerListUseCase(repoList),
	}
}

func NewCustomerCreateUseCase(
	repoCreate repository.CreateModelRepository[*model.EmployeeCreateCustomerResp, *model.CustomerCreateInput],
) usecase.IEmployeeCustomerCreateUseCase {
	return &customer.CustomerCreateUseCase{
		RepoCreate: repoCreate,
	}
}

func NewCustomerGetFirstUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.ICustomerGetFirstUseCase {
	uc := &customer.CustomerGetFirstUseCase{
		UC1: NewCustomerListUseCase(repoList),
	}
	return uc
}

func NewCustomerUpdateUseCase(
	repoUpdate repository.UpdateModelRepository[*model.Customer, *model.CustomerUpdateInput],
) usecase.ICustomerUpdateUseCase {
	uc := &customer.CustomerUpdateUseCase{
		RepoUpdate: repoUpdate,
	}
	return uc
}

func NewCustomerListUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.ICustomerListUseCase {
	uc := &customer.CustomerListUseCase{
		RepoList: repoList,
	}
	return uc
}

func NewEmployeeCustomerUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	repoCreate repository.CreateModelRepository[*model.EmployeeCreateCustomerResp, *model.CustomerCreateInput],
	repoIsNext repository.IIsNextModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	rle repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
	secretKey *string,
	prodOwnerName *string,
) usecase.IEmployeeCustomerUseCase {
	return &customer.EmployeeCustomerUseCase{
		IEmployeeCustomerCreateUseCase:         NewCustomerCreateUseCase(repoCreate),
		IEmployeeCustomerValidateCreateUseCase: NewCustomerValidateCreateUseCase(repoList),
		IEmployeeConfigUseCase:                 NewEmployeeConfigUseCase(secretKey, prodOwnerName),
		IEmployeeCustomerListUseCase:           NewCustomerListUseCase(repoList),
		IIsNextUseCase:                         NewIsNextUseCase(repoIsNext),
		IEmployeeCustomerGetFirstUseCase:       NewCustomerGetFirstUseCase(repoList),
		IEmployeeGetUserUseCase:                NewEmployeeGetUserUseCase(rle),
	}

}
