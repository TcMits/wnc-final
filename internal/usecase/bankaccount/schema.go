package bankaccount

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/webapi"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type (
	CustomerBankAccountUpdateUseCase struct {
		repoUpdate repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput]
	}
	CustomerBankAccountListUseCase struct {
		repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
	}
	CustomerBankAccountValidateUpdateInputUseCase struct {
		gFMUC usecase.ICustomerBankAccountGetFirstMineUseCase
	}
	CustomerBankAccountGetFirstUseCase struct {
		bALUC usecase.ICustomerBankAccountListUseCase
	}
	CustomerBankAccountListMineUseCase struct {
		bALUC usecase.ICustomerBankAccountListUseCase
	}
	CustomerBankAccountGetFirstMineUseCase struct {
		bALMUC usecase.ICustomerBankAccountListMineUseCase
	}
	CustomerBankAccountIsNextUseCase struct {
		iNUC usecase.IIsNextUseCase[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
	}
	CustomerTPBankBankAccountGetUseCase struct {
		w1 webapi.ITPBankAPI
	}
	CustomerBankAccountUseCase struct {
		usecase.ICustomerBankAccountUpdateUseCase
		usecase.ICustomerBankAccountValidateUpdateInputUseCase
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
		repoUpdate repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput]
	}
	EmployeeBankAccountListUseCase struct {
		repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
	}
	EmployeeBankAccountGetFirstUseCase struct {
		bALUC usecase.IEmployeeBankAccountListUseCase
	}
	EmployeeBankAccountUseCase struct {
		usecase.IEmployeeConfigUseCase
		usecase.IEmployeeGetUserUseCase
		usecase.IEmployeeBankAccountUpdateUseCase
		usecase.IEmployeeBankAccountValidateUpdateInputUseCase
		usecase.IEmployeeBankAccountGetFirstUseCase
		usecase.IEmployeeBankAccountListUseCase
		usecase.IIsNextUseCase[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
	}
)

type (
	PartnerBankAccountListUseCase struct {
		repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
	}
	PartnerBankAccountGetFirstUseCase struct {
		bALUC usecase.IPartnerBankAccountListUseCase
	}
	PartnerBankAccountRespGetFirstUseCase struct {
		uc1 usecase.IPartnerBankAccountListUseCase
		uc2 usecase.ICustomerGetFirstUseCase
	}
	PartnerBankAccountUseCase struct {
		usecase.IPartnerConfigUseCase
		usecase.IPartnerGetUserUseCase
		usecase.IPartnerBankAccountRespGetFirstUseCase
		usecase.IPartnerBankAccountListUseCase
		usecase.IIsNextUseCase[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput]
	}
)
