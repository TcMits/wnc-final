package bankaccount

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/webapi"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type (
	CustomerBankAccountUpdateUseCase struct {
		RepoUpdate repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput]
	}
	CustomerBankAccountListUseCase struct {
		RepoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
	}
	CustomerBankAccountValidateUpdateInputUseCase struct {
		UC1 usecase.ICustomerBankAccountGetFirstMineUseCase
	}
	CustomerBankAccountGetFirstUseCase struct {
		UC1 usecase.ICustomerBankAccountListUseCase
	}
	CustomerBankAccountListMineUseCase struct {
		UC1 usecase.ICustomerBankAccountListUseCase
	}
	CustomerBankAccountGetFirstMineUseCase struct {
		UC1 usecase.ICustomerBankAccountListMineUseCase
	}
	CustomerBankAccountDeleteUseCase struct {
		Repo repository.DeleteModelRepository[*model.BankAccount]
	}
	CustomerBankAccountIsNextUseCase struct {
		UC1 usecase.IIsNextUseCase[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
	}
	CustomerTPBankBankAccountGetUseCase struct {
		W1 webapi.ITPBankAPI
	}
	CustomerBankAccountUseCase struct {
		usecase.ICustomerBankAccountUpdateUseCase
		usecase.ICustomerBankAccountValidateUpdateInputUseCase
		usecase.ICustomerBankAccountDeleteUseCase
		usecase.ICustomerBankAccountListUseCase
		usecase.ICustomerConfigUseCase
		usecase.ICustomerGetUserUseCase
		usecase.ICustomerBankAccountGetFirstMineUseCase
		usecase.ICustomerBankAccountListMineUseCase
		usecase.ICustomerBankAccountGetFirstUseCase
		usecase.ICustomerTPBankBankAccountGetUseCase
		usecase.IIsNextUseCase[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
	}
)

type (
	EmployeeBankAccountValidateUpdateInputUseCase struct {
	}
	EmployeeBankAccountUpdateUseCase struct {
		RepoUpdate repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput]
	}
	EmployeeBankAccountListUseCase struct {
		RepoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
	}
	EmployeeBankAccountGetFirstUseCase struct {
		UC1 usecase.IEmployeeBankAccountListUseCase
	}
	EmployeeBankAccountDeleteUseCase struct {
		Repo repository.DeleteModelRepository[*model.BankAccount]
	}
	EmployeeBankAccountUseCase struct {
		usecase.IEmployeeConfigUseCase
		usecase.IEmployeeGetUserUseCase
		usecase.IEmployeeBankAccountDeleteUseCase
		usecase.IEmployeeBankAccountUpdateUseCase
		usecase.IEmployeeBankAccountValidateUpdateInputUseCase
		usecase.IEmployeeBankAccountGetFirstUseCase
		usecase.IEmployeeBankAccountListUseCase
		usecase.IIsNextUseCase[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
	}
)

type (
	PartnerBankAccountListUseCase struct {
		RepoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
	}
	PartnerBankAccountGetFirstUseCase struct {
		UC1 usecase.IPartnerBankAccountListUseCase
	}
	PartnerBankAccountRespGetFirstUseCase struct {
		UC1 usecase.IPartnerBankAccountListUseCase
		UC2 usecase.ICustomerGetFirstUseCase
	}
	PartnerBankAccountUseCase struct {
		usecase.IPartnerConfigUseCase
		usecase.IPartnerGetUserUseCase
		usecase.IPartnerBankAccountRespGetFirstUseCase
		usecase.IPartnerBankAccountListUseCase
		usecase.IIsNextUseCase[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
	}
)
