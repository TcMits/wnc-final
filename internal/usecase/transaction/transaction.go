package transaction

import (
	"time"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/admin"
	"github.com/TcMits/wnc-final/internal/usecase/auth"
	"github.com/TcMits/wnc-final/internal/usecase/bankaccount"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/internal/usecase/customer"
	"github.com/TcMits/wnc-final/internal/usecase/employee"
	"github.com/TcMits/wnc-final/internal/usecase/outliers"
	"github.com/TcMits/wnc-final/internal/webapi/tpbank"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/mail"
)

func NewCustomerTransactionListUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.ICustomerTransactionListUseCase {
	return &CustomerTransactionListUseCase{
		repoList: repoList,
	}
}

func NewCustomerTransactionListMineUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.ICustomerTransactionListMineUseCase {
	return &CustomerTransactionListMineUseCase{
		tLUC: NewCustomerTransactionListUseCase(repoList),
	}
}
func NewCustomerTransactionGetFirstMineUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.ICustomerTransactionGetFirstMineUseCase {
	return &CustomerTransactionGetFirstMineUseCase{
		tLMTUC: NewCustomerTransactionListMineUseCase(repoList),
	}
}

func NewCustomerTransactionUpdateUseCase(
	repoUpdate repository.UpdateModelRepository[*model.Transaction, *model.TransactionUpdateInput],
) usecase.ICustomerTransactionUpdateUseCase {
	return &CustomerTransactionUpdateUseCase{
		repoUpdate: repoUpdate,
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
	return &CustomerTransactionCreateUseCase{
		repoCreate:            repoCreate,
		taskExecutor:          taskExctor,
		cfUC:                  config.NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		txcConfirmSubjectMail: confirmSubjectMail,
		txcConfirmMailTemp:    confirmEmailTemplate,
		otpTimeout:            otpTimeout,
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
	return &CustomerTransactionValidateCreateInputUseCase{
		tLUC:   NewCustomerTransactionListUseCase(repoList),
		bAGFUC: bankaccount.NewCustomerBankAccountGetFirstUseCase(rlba),
		cGFUC:  customer.NewCustomerGetFirstUseCase(rlc),
		cfUC:   config.NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		w1: tpbank.NewTPBankAPI(
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
	return &CustomerTransactionConfirmSuccessUseCase{
		cfUC:   config.NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		tCRepo: repo,
	}
}
func NewCustomerTransactionValidateConfirmInputUseCase(
	sk *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerTransactionValidateConfirmInputUseCase {
	return &CustomerTransactionValidateConfirmInputUseCase{
		cfUC: config.NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
	}
}

func NewCustomerTransactionIsNextUseCase(
	repoIsNext repository.IIsNextModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.IIsNextUseCase[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput] {
	return &CustomerTransactionIsNextUseCase{
		iNUC: outliers.NewIsNextUseCase(repoIsNext),
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
	return &CustomerTransactionUseCase{
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
		ICustomerConfigUseCase:                          config.NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		ICustomerGetUserUseCase:                         auth.NewCustomerGetUserUseCase(rlc),
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
	return &EmployeeTransactionIsNextUseCase{
		iNUC: outliers.NewIsNextUseCase(repoIsNext),
	}
}
func NewEmployeeTransactionListUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.IEmployeeTransactionListUseCase {
	return &EmployeeTransactionListUseCase{
		repoList: repoList,
	}
}
func NewEmployeeTransactionGetFirstUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.IEmployeeTransactionGetFirstUseCase {
	return &EmployeeTransactionGetFirstUseCase{
		tLTUC: NewEmployeeTransactionListUseCase(repoList),
	}
}

func NewEmployeeTransactionUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
	repoIsNext repository.IIsNextModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
	rle repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
	sk,
	prodOwnerName *string,
) usecase.IEmployeeTransactionUseCase {
	return &EmployeeTransactionUseCase{
		IEmployeeTransactionListUseCase:     NewEmployeeTransactionListUseCase(repoList),
		IEmployeeConfigUseCase:              config.NewEmployeeConfigUseCase(sk, prodOwnerName),
		IEmployeeGetUserUseCase:             employee.NewEmployeeGetUserUseCase(rle),
		IEmployeeTransactionGetFirstUseCase: NewEmployeeTransactionGetFirstUseCase(repoList),
		IIsNextUseCase:                      NewEmployeeTransactionIsNextUseCase(repoIsNext),
	}
}
func NewAdminTransactionIsNextUseCase(
	repoIsNext repository.IIsNextModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.IIsNextUseCase[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput] {
	return &AdminTransactionIsNextUseCase{
		iNUC: outliers.NewIsNextUseCase(repoIsNext),
	}
}
func NewAdminTransactionListUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.IAdminTransactionListUseCase {
	return &AdminTransactionListUseCase{
		repoList: repoList,
	}
}
func NewAdminTransactionGetFirstUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.IAdminTransactionGetFirstUseCase {
	return &AdminTransactionGetFirstUseCase{
		tLTUC: NewAdminTransactionListUseCase(repoList),
	}
}

func NewAdminTransactionUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
	repoIsNext repository.IIsNextModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
	rle repository.ListModelRepository[*model.Admin, *model.AdminOrderInput, *model.AdminWhereInput],
	sk,
	prodOwnerName *string,
) usecase.IAdminTransactionUseCase {
	return &AdminTransactionUseCase{
		IAdminTransactionListUseCase:     NewAdminTransactionListUseCase(repoList),
		IAdminConfigUseCase:              config.NewAdminConfigUseCase(sk, prodOwnerName),
		IAdminGetUserUseCase:             admin.NewAdminGetUserUseCase(rle),
		IAdminTransactionGetFirstUseCase: NewAdminTransactionGetFirstUseCase(repoList),
		IIsNextUseCase:                   NewAdminTransactionIsNextUseCase(repoIsNext),
	}
}

func NewPartnerTransactionCreateUseCase(
	repo repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput],
) usecase.IPartnerTransactionCreateUseCase {
	return &PartnerTransactionCreateUseCase{repo: repo}
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
	return &PartnerTransactionValidateCreateInputUseCase{
		uc1:    bankaccount.NewPartnerBankAccountGetFirstUseCase(r1),
		uc2:    config.NewPartnerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		uc3:    customer.NewCustomerGetFirstUseCase(r2),
		layout: *layout,
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
	return &PartnerTransactionUseCase{
		IPartnerTransactionValidateCreateUseCase: NewPartnerTransactionValidateCreateUseCase(r1, r2, sk, prodOwnerName, feeDesc, layout, fee),
		IPartnerTransactionCreateUseCase:         NewPartnerTransactionCreateUseCase(r3),
		IPartnerGetUserUseCase:                   auth.NewPartnerGetUserUseCase(r4),
		IPartnerConfigUseCase:                    config.NewPartnerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
	}
}
