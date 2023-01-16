package constructors

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/logic/me"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

func NewCustomerChangePasswordUseCase(
	repoUpdate repository.UpdateModelRepository[*model.Customer, *model.CustomerUpdateInput],
) usecase.ICustomerChangePasswordUseCase {
	return &me.CustomerChangePasswordUseCase{UC1: NewCustomerUpdateUseCase(repoUpdate)}
}
func NewCustomerValidateChangePasswordUseCase() usecase.ICustomerValidateChangePasswordUseCase {
	return &me.CustomerValidateChangePasswordUseCase{}
}
func NewCustomerGetUserFromCtxUserCase() usecase.ICustomerGetUserFromCtxUseCase {
	return &me.CustomerGetUserFromCtx{}
}

func NewCustomerMeUseCase(
	repoList repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	repoUpdate repository.UpdateModelRepository[*model.Customer, *model.CustomerUpdateInput],
	sk *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerMeUseCase {
	uc := &me.CustomerMeUseCase{
		ICustomerConfigUseCase:                 NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		ICustomerGetUserUseCase:                NewCustomerGetUserUseCase(repoList),
		ICustomerChangePasswordUseCase:         NewCustomerChangePasswordUseCase(repoUpdate),
		ICustomerValidateChangePasswordUseCase: NewCustomerValidateChangePasswordUseCase(),
		ICustomerGetUserFromCtxUseCase:         NewCustomerGetUserFromCtxUserCase(),
	}
	return uc
}
func NewEmployeeGetUserFromCtxUserCase() usecase.IEmployeeGetUserFromCtxUseCase {
	return &me.EmployeeGetUserFromCtx{}
}

func NewEmployeeMeUseCase(
	repoList repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
	repoUpdate repository.UpdateModelRepository[*model.Employee, *model.EmployeeUpdateInput],
	sk *string,
	prodOwnerName *string,
) usecase.IEmployeeMeUseCase {
	uc := &me.EmployeeMeUseCase{
		IEmployeeConfigUseCase:         NewEmployeeConfigUseCase(sk, prodOwnerName),
		IEmployeeGetUserUseCase:        NewEmployeeGetUserUseCase(repoList),
		IEmployeeGetUserFromCtxUseCase: NewEmployeeGetUserFromCtxUserCase(),
	}
	return uc
}
func NewAdminGetUserFromCtxUserCase() usecase.IAdminGetUserFromCtxUseCase {
	return &me.AdminGetUserFromCtx{}
}

func NewAdminMeUseCase(
	repoList repository.ListModelRepository[*model.Admin, *model.AdminOrderInput, *model.AdminWhereInput],
	repoUpdate repository.UpdateModelRepository[*model.Admin, *model.AdminUpdateInput],
	sk *string,
	prodOwnerName *string,
) usecase.IAdminMeUseCase {
	uc := &me.AdminMeUseCase{
		IAdminConfigUseCase:         NewAdminConfigUseCase(sk, prodOwnerName),
		IAdminGetUserUseCase:        NewAdminGetUserUseCase(repoList),
		IAdminGetUserFromCtxUseCase: NewAdminGetUserFromCtxUserCase(),
	}
	return uc
}
