package debt

import (
	"time"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/mail"
)

type (
	CustomerDebtListUseCase struct {
		RepoList repository.ListModelRepository[*model.Debt, *model.DebtOrderInput, *model.DebtWhereInput]
	}
	CustomerDebtCreateUseCase struct {
		TaskExecutor task.IExecuteTask[*task.DebtNotifyPayload]
		RepoCreate   repository.CreateModelRepository[*model.Debt, *model.DebtCreateInput]
		UC1          usecase.ICustomerGetFirstUseCase
	}
	CustomerDebtValidateCreateInputUseCase struct {
		UC1 usecase.ICustomerConfigUseCase
		UC2 usecase.ICustomerBankAccountGetFirstUseCase
		UC3 usecase.ICustomerGetFirstUseCase
	}
	CustomerDebtListMineUseCase struct {
		UC1 usecase.ICustomerDebtListUseCase
	}
	CustomerDebtGetFirstMineUseCase struct {
		UC1 usecase.ICustomerDebtListMineUseCase
	}
	CustomerDebtUpdateUseCase struct {
		RepoUpdate repository.UpdateModelRepository[*model.Debt, *model.DebtUpdateInput]
	}
	CustomerDebtValidateCancelUseCase struct {
		UC1 usecase.ICustomerGetFirstUseCase
	}
	CustomerDebtCancelUseCase struct {
		TaskExecutor task.IExecuteTask[*task.DebtNotifyPayload]
		UC1          usecase.ICustomerDebtUpdateUseCase
		UC2          usecase.ICustomerGetFirstUseCase
	}
	CustomerDebtValidateFulfillUseCase struct {
		UC1 usecase.ICustomerGetFirstUseCase
		UC2 usecase.ICustomerBankAccountGetFirstUseCase
	}
	CustomerDebtValidateFulfillWithTokenUseCase struct {
		UC1 usecase.ICustomerConfigUseCase
	}
	CustomerDebtFulfillUseCase struct {
		TaskExecutor           task.IExecuteTask[*mail.EmailPayload]
		DebtFulfillSubjectMail *string
		DebtFulfillMailTemp    *string
		OtpTimeout             time.Duration
		UC1                    usecase.ICustomerConfigUseCase
	}
	CustomerDebtFulfillWithTokenUseCase struct {
		RepoFulfill  repository.IDebtFullfillRepository
		TaskExecutor task.IExecuteTask[*task.DebtNotifyPayload]
		UC1          usecase.ICustomerGetFirstUseCase
	}
	CustomerDebtIsNextUseCase struct {
		UC1 usecase.IIsNextUseCase[*model.Debt, *model.DebtOrderInput, *model.DebtWhereInput]
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
		usecase.ICustomerDebtValidateFulfillUseCase
		usecase.ICustomerDebtFulfillUseCase
		usecase.ICustomerDebtFulfillWithTokenUseCase
		usecase.ICustomerDebtValidateFulfillWithTokenUseCase
		usecase.IIsNextUseCase[*model.Debt, *model.DebtOrderInput, *model.DebtWhereInput]
	}
)
