package transaction

import (
	"time"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/mail"
)

type (
	CustomerTransactionUpdateUseCase struct {
		repoUpdate repository.UpdateModelRepository[*model.Transaction, *model.TransactionUpdateInput]
	}
	CustomerTransactionListUseCase struct {
		repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
	}
	CustomerTransactionCreateUseCase struct {
		cfUC                  usecase.ICustomerConfigUseCase
		repoCreate            repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput]
		taskExecutor          task.IExecuteTask[*mail.EmailPayload]
		txcConfirmSubjectMail *string
		txcConfirmMailTemp    *string
		otpTimeout            time.Duration
	}
	CustomerTransactionValidateConfirmInputUseCase struct {
		cfUC usecase.ICustomerConfigUseCase
	}
	CustomerTransactionConfirmSuccessUseCase struct {
		cfUC   usecase.ICustomerConfigUseCase
		tCRepo repository.ITransactionConfirmSuccessRepository
	}
	CustomerTransactionListMineUseCase struct {
		tLUC usecase.ICustomerTransactionListUseCase
	}
	CustomerTransactionGetFirstMineUseCase struct {
		tLMTUC usecase.ICustomerTransactionListMineUseCase
	}
	CustomerTransactionValidateCreateInputUseCase struct {
		cfUC   usecase.ICustomerConfigUseCase
		bAGFUC usecase.ICustomerBankAccountGetFirstUseCase
		cGFUC  usecase.ICustomerGetFirstUseCase
		tLUC   usecase.ICustomerTransactionListUseCase
	}
	CustomerTransactionIsNextUseCase struct {
		iNUC usecase.IIsNextUseCase[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
	}
	CustomerTransactionUseCase struct {
		usecase.ICustomerTransactionValidateCreateInputUseCase
		usecase.ICustomerTransactionCreateUseCase
		usecase.ICustomerTransactionListUseCase
		usecase.ICustomerConfigUseCase
		usecase.ICustomerGetUserUseCase
		usecase.ICustomerTransactionConfirmSuccessUseCase
		usecase.ICustomerTransactionValidateConfirmInputUseCase
		usecase.ICustomerTransactionListMineUseCase
		usecase.ICustomerTransactionGetFirstMineUseCase
		usecase.IIsNextUseCase[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
	}
)
