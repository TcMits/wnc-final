package transaction

import (
	"time"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/webapi"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/mail"
)

type (
	CustomerTransactionUpdateUseCase struct {
		RepoUpdate repository.UpdateModelRepository[*model.Transaction, *model.TransactionUpdateInput]
	}
	CustomerTransactionListUseCase struct {
		RepoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
	}
	CustomerTransactionCreateUseCase struct {
		UC1                   usecase.ICustomerConfigUseCase
		RepoCreate            repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput]
		TaskExecutor          task.IExecuteTask[*mail.EmailPayload]
		TxcConfirmSubjectMail *string
		TxcConfirmMailTemp    *string
		OtpTimeout            time.Duration
	}
	CustomerTransactionValidateConfirmInputUseCase struct {
		UC1 usecase.ICustomerConfigUseCase
	}
	CustomerTransactionConfirmSuccessUseCase struct {
		UC1 usecase.ICustomerConfigUseCase
		R   repository.ITransactionConfirmSuccessRepository
	}
	CustomerTransactionListMineUseCase struct {
		UC1 usecase.ICustomerTransactionListUseCase
	}
	CustomerTransactionGetFirstMineUseCase struct {
		UC1 usecase.ICustomerTransactionListMineUseCase
	}
	CustomerTransactionValidateCreateInputUseCase struct {
		UC1 usecase.ICustomerConfigUseCase
		UC2 usecase.ICustomerBankAccountGetFirstUseCase
		UC3 usecase.ICustomerGetFirstUseCase
		UC4 usecase.ICustomerTransactionListUseCase
		W1  webapi.ITPBankAPI
	}
	CustomerTransactionIsNextUseCase struct {
		UC1 usecase.IIsNextUseCase[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
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

type (
	EmployeeTransactionIsNextUseCase struct {
		UC1 usecase.IIsNextUseCase[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
	}
	EmployeeTransactionListUseCase struct {
		RepoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
	}
	EmployeeTransactionGetFirstUseCase struct {
		UC1 usecase.IEmployeeTransactionListUseCase
	}
	EmployeeTransactionUseCase struct {
		usecase.IEmployeeConfigUseCase
		usecase.IEmployeeGetUserUseCase
		usecase.IEmployeeTransactionListUseCase
		usecase.IEmployeeTransactionGetFirstUseCase
		usecase.IIsNextUseCase[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
	}
)

type (
	AdminTransactionIsNextUseCase struct {
		UC1 usecase.IIsNextUseCase[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
	}
	AdminTransactionListUseCase struct {
		RepoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
	}
	AdminTransactionGetFirstUseCase struct {
		UC1 usecase.IAdminTransactionListUseCase
	}
	AdminTransactionUseCase struct {
		usecase.IAdminConfigUseCase
		usecase.IAdminGetUserUseCase
		usecase.IAdminTransactionListUseCase
		usecase.IAdminTransactionGetFirstUseCase
		usecase.IIsNextUseCase[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput]
	}
)

type (
	PartnerTransactionValidateCreateInputUseCase struct {
		UC1    usecase.IPartnerBankAccountGetFirstUseCase
		UC2    usecase.IPartnerConfigUseCase
		UC3    usecase.ICustomerGetFirstUseCase
		W1     webapi.ITPBankAPI
		Layout string
	}
	PartnerTransactionCreateUseCase struct {
		Repo repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput]
	}
	PartnerTransactionUseCase struct {
		usecase.IPartnerGetUserUseCase
		usecase.IPartnerConfigUseCase
		usecase.IPartnerTransactionValidateCreateUseCase
		usecase.IPartnerTransactionCreateUseCase
	}
)
