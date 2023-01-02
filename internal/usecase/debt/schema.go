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
		taskExecutor task.IExecuteTask[*task.DebtNotifyPayload]
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
	CustomerDebtUpdateUseCase struct {
		repoUpdate repository.UpdateModelRepository[*model.Debt, *model.DebtUpdateInput]
	}
	CustomerDebtValidateCancelUseCase struct {
		cGFUC usecase.ICustomerGetFirstUseCase
	}
	CustomerDebtCancelUseCase struct {
		taskExecutor task.IExecuteTask[*task.DebtNotifyPayload]
		dUUc         usecase.ICustomerDebtUpdateUseCase
		cGFUC        usecase.ICustomerGetFirstUseCase
	}
	CustomerDebtUseCase struct {
		usecase.ICustomerConfigUseCase
		usecase.ICustomerGetUserUseCase
		usecase.ICustomerDebtListUseCase
		usecase.ICustomerDebtCreateUseCase
		usecase.ICustomerDebtValidateCreateInputUseCase
		usecase.ICustomerDebtGetFirstMineUseCase
		usecase.ICustomerDebtListMineUseCase
		usecase.ICustomerDebtValidateCancelUseCase
		usecase.ICustomerDebtCancelUseCase
	}
)
