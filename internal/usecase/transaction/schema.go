package transaction

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type (
	CustomerTransactionUpdateUseCase struct {
		repoUpdate repository.UpdateModelRepository[*model.Transaction, *model.TransactionUpdateInput]
	}
	CustomerTransactionListUseCase struct {
		repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
	}
	CustomerTransactionCreateUseCase struct {
		repoCreate repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput]
	}
	CustomerTransactionValidateConfirmInputUseCase struct {
		cfUC usecase.ICustomerConfigUseCase
	}
	CustomerTransactionConfirmSuccessUseCase struct {
		cfUC   usecase.ICustomerConfigUseCase
		tCRepo repository.ITransactionConfirmSuccessRepository
	}
	CustomerTransactionListMyTxcUseCase struct {
		tLUC usecase.ICustomerTransactionListUseCase
	}
	CustomerTransactionGetFirstMyTxcUseCase struct {
		tLMTUC usecase.ICustomerTransactionListMyTxcUseCase
	}
	CustomerTransactionValidateCreateInputUseCase struct {
		cfUC   usecase.ICustomerConfigUseCase
		bAGFUC usecase.ICustomerBankAccountGetFirstUseCase
		cGFUC  usecase.ICustomerGetFirstUseCase
		tLUC   usecase.ICustomerTransactionListUseCase
	}
	CustomerTransactionUseCase struct {
		usecase.ICustomerTransactionValidateCreateInputUseCase
		usecase.ICustomerTransactionCreateUseCase
		usecase.ICustomerTransactionListUseCase
		usecase.ICustomerConfigUseCase
		usecase.ICustomerGetUserUseCase
		usecase.ICustomerTransactionConfirmSuccessUseCase
		usecase.ICustomerTransactionValidateConfirmInputUseCase
		usecase.ICustomerTransactionListMyTxcUseCase
		usecase.ICustomerTransactionGetFirstMyTxUseCase
	}
)
