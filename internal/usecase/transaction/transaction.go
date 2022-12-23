package transaction

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type (
	CustomerTransactionCreateUseCase struct {
		bKUUC      usecase.ICustomerBankAccountUpdateUseCase
		bKGFUC     usecase.ICustomerBankAccountGetFirstUseCase
		repoCreate repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput]
	}
	CustomerTransactionUpdateUseCase struct {
		repoUpdate repository.UpdateModelRepository[*model.Transaction, *model.TransactionUpdateInput]
	}
	CustomerTransactionValidateConfirmInputUseCase struct {
		cfUC usecase.ICustomerConfigUseCase
	}
	CustomerTransactionConfirmSuccessUseCase struct {
		cfUC   usecase.ICustomerConfigUseCase
		bAUUC  usecase.ICustomerBankAccountUpdateUseCase
		bAGFUC usecase.ICustomerBankAccountGetFirstUseCase
		tCUC   usecase.ICustomerTransactionCreateUseCase
		tUUC   usecase.ICustomerTransactionUpdateUseCase
	}
	CustomerTransactionListUseCase struct {
		repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
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
	}
)
