package debt

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type (
	CustomerDebtListUseCase struct {
		repoList repository.ListModelRepository[*model.Debt, *model.DebtOrderInput, *model.DebtWhereInput]
	}
	CustomerDebtCreateUseCase struct {
		taskExecutor task.IExecuteTask[*task.DebtCreateNotifyPayload]
		repoCreate   repository.CreateModelRepository[*model.Debt, *model.DebtCreateInput]
		cGFUC        usecase.ICustomerGetFirstUseCase
	}
	CustomerDebtValidateCreateInputUseCase struct {
		cfUC   usecase.ICustomerConfigUseCase
		bAGFUC usecase.ICustomerBankAccountGetFirstUseCase
		cGFUC  usecase.ICustomerGetFirstUseCase
	}
	CustomerDebtListMineUseCase struct {
		dLUC usecase.ICustomerDebtListUseCase
	}
	CustomerDebtGetFirstMineUseCase struct {
		dLMUC usecase.ICustomerDebtListMineUseCase
	}
	CustomerDebtUseCase struct {
		usecase.ICustomerConfigUseCase
		usecase.ICustomerGetUserUseCase
		usecase.ICustomerDebtListUseCase
		usecase.ICustomerDebtCreateUseCase
		usecase.ICustomerDebtValidateCreateInputUseCase
		usecase.ICustomerDebtGetFirstMineUseCase
		usecase.ICustomerDebtListMineUseCase
	}
)
