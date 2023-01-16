package constructors

import (
	"time"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/logic/transaction"
	"github.com/TcMits/wnc-final/internal/webapi/tpbank"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/mail"
)

func NewCustomerTransactionListUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.ICustomerTransactionListUseCase {
	return &transaction.CustomerTransactionListUseCase{
		RepoList: repoList,
	}
}

func NewCustomerTransactionListMineUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.ICustomerTransactionListMineUseCase {
	return &transaction.CustomerTransactionListMineUseCase{
		UC1: NewCustomerTransactionListUseCase(repoList),
	}
}
func NewCustomerTransactionGetFirstMineUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.ICustomerTransactionGetFirstMineUseCase {
	return &transaction.CustomerTransactionGetFirstMineUseCase{
		UC1: NewCustomerTransactionListMineUseCase(repoList),
	}
}

func NewCustomerTransactionUpdateUseCase(
	repoUpdate repository.UpdateModelRepository[*model.Transaction, *model.TransactionUpdateInput],
) usecase.ICustomerTransactionUpdateUseCase {
	return &transaction.CustomerTransactionUpdateUseCase{
		RepoUpdate: repoUpdate,
	}
}

func NewCustomerTransactionCreateUseCase(
	taskExctor task.IExecuteTask[*mail.EmailPayload],
	repoCreate repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput],
	sk,
	prodOwnerName,
	feeDesc,
	confirmSubjectMail,
	confirmEmailTemplate *string,
	fee *float64,
	otpTimeout time.Duration,
) usecase.ICustomerTransactionCreateUseCase {
	return &transaction.CustomerTransactionCreateUseCase{
		RepoCreate:            repoCreate,
		TaskExecutor:          taskExctor,
		UC1:                   NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		TxcConfirmSubjectMail: confirmSubjectMail,
		TxcConfirmMailTemp:    confirmEmailTemplate,
		OtpTimeout:            otpTimeout,
	}
}

func NewCustomerTransactionValidateCreateInputUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
	rlba repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	sk,
	prodOwnerName,
	feeDesc *string,
	layout,
	baseUrl,
	authAPI,
	bankAccountAPI,
	validateAPI,
	createTransactionAPI,
	tpBankName,
	tpBankApiKey,
	tpBankSecretKey,
	tpBankPrivateK string,
	fee *float64,
) usecase.ICustomerTransactionValidateCreateInputUseCase {
	return &transaction.CustomerTransactionValidateCreateInputUseCase{
		UC4: NewCustomerTransactionListUseCase(repoList),
		UC2: NewCustomerBankAccountGetFirstUseCase(rlba),
		UC3: NewCustomerGetFirstUseCase(rlc),
		UC1: NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		W1: tpbank.NewTPBankAPI(
			tpBankName,
			tpBankApiKey,
			tpBankPrivateK,
			tpBankSecretKey,
			layout,
			baseUrl,
			authAPI,
			bankAccountAPI,
			createTransactionAPI,
			validateAPI,
		),
	}
}

func NewCustomerTransactionConfirmSuccessUseCase(
	repo repository.ITransactionConfirmSuccessRepository,
	sk *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerTransactionConfirmSuccessUseCase {
	return &transaction.CustomerTransactionConfirmSuccessUseCase{
		UC1: NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		R:   repo,
	}
}
func NewCustomerTransactionValidateConfirmInputUseCase(
	sk *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerTransactionValidateConfirmInputUseCase {
	return &transaction.CustomerTransactionValidateConfirmInputUseCase{
		UC1: NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
	}
}

func NewCustomerTransactionIsNextUseCase(
	repoIsNext repository.IIsNextModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.IIsNextUseCase[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput] {
	return &transaction.CustomerTransactionIsNextUseCase{
		UC1: NewIsNextUseCase(repoIsNext),
	}
}

func NewCustomerTransactionUseCase(
	taskExctor task.IExecuteTask[*mail.EmailPayload],
	repoConfirm repository.ITransactionConfirmSuccessRepository,
	repoCreate repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput],
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
	repoUpdate repository.UpdateModelRepository[*model.Transaction, *model.TransactionUpdateInput],
	repoIsNext repository.IIsNextModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	rlba repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	rUBa repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput],
	sk,
	prodOwnerName,
	feeDesc,
	confirmEmailSubject,
	confirmEmailTemplate *string,
	layout,
	baseUrl,
	authAPI,
	bankAccountAPI,
	validateAPI,
	createTransactionAPI,
	tpBankName,
	tpBankApiKey,
	tpBankSecretKey,
	tpBankPrivateK string,
	fee *float64,
	otpTimeout time.Duration,
) usecase.ICustomerTransactionUseCase {
	return &transaction.CustomerTransactionUseCase{
		ICustomerTransactionCreateUseCase: NewCustomerTransactionCreateUseCase(taskExctor, repoCreate, sk, prodOwnerName, feeDesc, confirmEmailSubject, confirmEmailTemplate, fee, otpTimeout),
		ICustomerTransactionValidateCreateInputUseCase: NewCustomerTransactionValidateCreateInputUseCase(repoList, rlba, rlc, sk, prodOwnerName, feeDesc,
			layout,
			baseUrl,
			authAPI,
			bankAccountAPI,
			validateAPI,
			createTransactionAPI,
			tpBankName,
			tpBankApiKey,
			tpBankSecretKey,
			tpBankPrivateK,
			fee,
		),
		ICustomerTransactionListUseCase:                 NewCustomerTransactionListUseCase(repoList),
		ICustomerConfigUseCase:                          NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		ICustomerGetUserUseCase:                         NewCustomerGetUserUseCase(rlc),
		ICustomerTransactionConfirmSuccessUseCase:       NewCustomerTransactionConfirmSuccessUseCase(repoConfirm, sk, prodOwnerName, fee, feeDesc),
		ICustomerTransactionListMineUseCase:             NewCustomerTransactionListMineUseCase(repoList),
		ICustomerTransactionGetFirstMineUseCase:         NewCustomerTransactionGetFirstMineUseCase(repoList),
		ICustomerTransactionValidateConfirmInputUseCase: NewCustomerTransactionValidateConfirmInputUseCase(sk, prodOwnerName, fee, feeDesc),
		IIsNextUseCase:                                  NewCustomerTransactionIsNextUseCase(repoIsNext),
	}
}

func NewEmployeeTransactionIsNextUseCase(
	repoIsNext repository.IIsNextModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.IIsNextUseCase[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput] {
	return &transaction.EmployeeTransactionIsNextUseCase{
		UC1: NewIsNextUseCase(repoIsNext),
	}
}
func NewEmployeeTransactionListUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.IEmployeeTransactionListUseCase {
	return &transaction.EmployeeTransactionListUseCase{
		RepoList: repoList,
	}
}
func NewEmployeeTransactionGetFirstUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.IEmployeeTransactionGetFirstUseCase {
	return &transaction.EmployeeTransactionGetFirstUseCase{
		UC1: NewEmployeeTransactionListUseCase(repoList),
	}
}

func NewEmployeeTransactionUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
	repoIsNext repository.IIsNextModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
	rle repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
	sk,
	prodOwnerName *string,
) usecase.IEmployeeTransactionUseCase {
	return &transaction.EmployeeTransactionUseCase{
		IEmployeeTransactionListUseCase:     NewEmployeeTransactionListUseCase(repoList),
		IEmployeeConfigUseCase:              NewEmployeeConfigUseCase(sk, prodOwnerName),
		IEmployeeGetUserUseCase:             NewEmployeeGetUserUseCase(rle),
		IEmployeeTransactionGetFirstUseCase: NewEmployeeTransactionGetFirstUseCase(repoList),
		IIsNextUseCase:                      NewEmployeeTransactionIsNextUseCase(repoIsNext),
	}
}
func NewAdminTransactionIsNextUseCase(
	repoIsNext repository.IIsNextModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.IIsNextUseCase[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput] {
	return &transaction.AdminTransactionIsNextUseCase{
		UC1: NewIsNextUseCase(repoIsNext),
	}
}
func NewAdminTransactionListUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.IAdminTransactionListUseCase {
	return &transaction.AdminTransactionListUseCase{
		RepoList: repoList,
	}
}
func NewAdminTransactionGetFirstUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.IAdminTransactionGetFirstUseCase {
	return &transaction.AdminTransactionGetFirstUseCase{
		UC1: NewAdminTransactionListUseCase(repoList),
	}
}

func NewAdminTransactionUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
	repoIsNext repository.IIsNextModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
	rle repository.ListModelRepository[*model.Admin, *model.AdminOrderInput, *model.AdminWhereInput],
	sk,
	prodOwnerName *string,
) usecase.IAdminTransactionUseCase {
	return &transaction.AdminTransactionUseCase{
		IAdminTransactionListUseCase:     NewAdminTransactionListUseCase(repoList),
		IAdminConfigUseCase:              NewAdminConfigUseCase(sk, prodOwnerName),
		IAdminGetUserUseCase:             NewAdminGetUserUseCase(rle),
		IAdminTransactionGetFirstUseCase: NewAdminTransactionGetFirstUseCase(repoList),
		IIsNextUseCase:                   NewAdminTransactionIsNextUseCase(repoIsNext),
	}
}

func NewPartnerTransactionCreateUseCase(
	repo repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput],
) usecase.IPartnerTransactionCreateUseCase {
	return &transaction.PartnerTransactionCreateUseCase{Repo: repo}
}

func NewPartnerTransactionValidateCreateUseCase(
	r1 repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	r2 repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	sk,
	prodOwnerName,
	feeDesc,
	layout *string,
	fee *float64,
) usecase.IPartnerTransactionValidateCreateUseCase {
	return &transaction.PartnerTransactionValidateCreateInputUseCase{
		UC1:    NewPartnerBankAccountGetFirstUseCase(r1),
		UC2:    NewPartnerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		UC3:    NewCustomerGetFirstUseCase(r2),
		Layout: *layout,
	}
}

func NewPartnerTransactionUseCase(
	r1 repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	r2 repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	r3 repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput],
	r4 repository.ListModelRepository[*model.Partner, *model.PartnerOrderInput, *model.PartnerWhereInput],
	sk,
	prodOwnerName,
	feeDesc,
	layout *string,
	fee *float64,
) usecase.IPartnerTransactionUseCase {
	return &transaction.PartnerTransactionUseCase{
		IPartnerTransactionValidateCreateUseCase: NewPartnerTransactionValidateCreateUseCase(r1, r2, sk, prodOwnerName, feeDesc, layout, fee),
		IPartnerTransactionCreateUseCase:         NewPartnerTransactionCreateUseCase(r3),
		IPartnerGetUserUseCase:                   NewPartnerGetUserUseCase(r4),
		IPartnerConfigUseCase:                    NewPartnerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
	}
}
