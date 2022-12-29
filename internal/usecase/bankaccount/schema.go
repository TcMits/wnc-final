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
		bALUC usecase.ICustomerBankAccountListUseCase
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
	CustomerBankAccountUseCase struct {
		usecase.ICustomerBankAccountUpdateUseCase
		usecase.ICustomerBankAccountValidateUpdateInputUseCase
		usecase.ICustomerBankAccountListUseCase
		usecase.ICustomerConfigUseCase
		usecase.ICustomerGetUserUseCase
		usecase.ICustomerBankAccountGetFirstMineUseCase
		usecase.ICustomerBankAccountListMineUseCase
		usecase.ICustomerBankAccountGetFirstUseCase
	}
)
