package bankaccount

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
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
	CustomerBankAccountUseCase struct {
		usecase.ICustomerBankAccountUpdateUseCase
		usecase.ICustomerBankAccountValidateUpdateInputUseCase
		usecase.ICustomerBankAccountListUseCase
		usecase.ICustomerConfigUseCase
		usecase.ICustomerGetUserUseCase
		usecase.ICustomerBankAccountGetFirstMineUseCase
		usecase.ICustomerBankAccountListMineUseCase
		usecase.ICustomerBankAccountGetFirstUseCase
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
