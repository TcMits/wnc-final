package constructors

import (
	"time"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/logic/debt"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/mail"
)

func NewCustomerDebtListUseCase(
	repoList repository.ListModelRepository[*model.Debt, *model.DebtOrderInput, *model.DebtWhereInput],
) usecase.ICustomerDebtListUseCase {
	return &debt.CustomerDebtListUseCase{RepoList: repoList}
}
func NewCustomerDebtUpdateUseCase(
	repoUpdate repository.UpdateModelRepository[*model.Debt, *model.DebtUpdateInput],
) usecase.ICustomerDebtUpdateUseCase {
	return &debt.CustomerDebtUpdateUseCase{RepoUpdate: repoUpdate}
}

func NewCustomerDebtCreateUseCase(
	repoCreate repository.CreateModelRepository[*model.Debt, *model.DebtCreateInput],
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	taskExctor task.IExecuteTask[*task.DebtNotifyPayload],
	sk *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerDebtCreateUseCase {
	return &debt.CustomerDebtCreateUseCase{
		RepoCreate:   repoCreate,
		TaskExecutor: taskExctor,
		UC1:          NewCustomerGetFirstUseCase(rlc),
	}
}
func NewCustomerDebtValidateCreateInputUseCase(
	rlba repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	sk *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerDebtValidateCreateInputUseCase {
	return &debt.CustomerDebtValidateCreateInputUseCase{
		UC2: NewCustomerBankAccountGetFirstUseCase(rlba),
		UC3: NewCustomerGetFirstUseCase(rlc),
		UC1: NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
	}
}
func NewCustomerDebtListMineUseCase(
	repoList repository.ListModelRepository[*model.Debt, *model.DebtOrderInput, *model.DebtWhereInput],
) usecase.ICustomerDebtListMineUseCase {
	return &debt.CustomerDebtListMineUseCase{
		UC1: NewCustomerDebtListUseCase(repoList),
	}
}
func NewCustomerDebtGetFirstMineUseCase(
	repoList repository.ListModelRepository[*model.Debt, *model.DebtOrderInput, *model.DebtWhereInput],
) usecase.ICustomerDebtGetFirstMineUseCase {
	return &debt.CustomerDebtGetFirstMineUseCase{
		UC1: NewCustomerDebtListMineUseCase(repoList),
	}
}
func NewCustomerDebtCancelUseCase(
	repoUpdate repository.UpdateModelRepository[*model.Debt, *model.DebtUpdateInput],
	taskExctor task.IExecuteTask[*task.DebtNotifyPayload],
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.ICustomerDebtCancelUseCase {
	return &debt.CustomerDebtCancelUseCase{
		TaskExecutor: taskExctor,
		UC1:          NewCustomerDebtUpdateUseCase(repoUpdate),
		UC2:          NewCustomerGetFirstUseCase(rlc),
	}
}
func NewCustomerDebtValidateCancelUseCase(
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.ICustomerDebtValidateCancelUseCase {
	return &debt.CustomerDebtValidateCancelUseCase{UC1: NewCustomerGetFirstUseCase(rlc)}
}
func NewCustomerDebtValidateFulfillUseCase(
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	rlba repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.ICustomerDebtValidateFulfillUseCase {
	return &debt.CustomerDebtValidateFulfillUseCase{UC1: NewCustomerGetFirstUseCase(rlc), UC2: NewCustomerBankAccountGetFirstUseCase(rlba)}
}
func NewCustomerDebtFulfillUseCase(
	taskExctor task.IExecuteTask[*mail.EmailPayload],
	sk,
	prodOwnerName,
	feeDesc,
	debtFulfillSubjectMail,
	debtFulfillEmailTemplate *string,
	fee *float64,
	otpTimeout time.Duration,
) usecase.ICustomerDebtFulfillUseCase {
	return &debt.CustomerDebtFulfillUseCase{
		TaskExecutor:           taskExctor,
		UC1:                    NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		DebtFulfillMailTemp:    debtFulfillEmailTemplate,
		DebtFulfillSubjectMail: debtFulfillSubjectMail,
		OtpTimeout:             otpTimeout,
	}
}
func NewCustomerDebtValidateFulfillWithTokenUseCase(
	sk,
	prodOwnerName,
	feeDesc *string,
	fee *float64,
) usecase.ICustomerDebtValidateFulfillWithTokenUseCase {
	return &debt.CustomerDebtValidateFulfillWithTokenUseCase{
		UC1: NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
	}
}
func NewCustomerDebtFulfillWithTokenUseCase(
	repoFulfill repository.IDebtFullfillRepository,
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	taskExctor task.IExecuteTask[*task.DebtNotifyPayload],
) usecase.ICustomerDebtFulfillWithTokenUseCase {
	return &debt.CustomerDebtFulfillWithTokenUseCase{RepoFulfill: repoFulfill, UC1: NewCustomerGetFirstUseCase(rlc), TaskExecutor: taskExctor}
}
func NewCustomerDebtIsNextUseCase(
	repoIsNext repository.IIsNextModelRepository[*model.Debt, *model.DebtOrderInput, *model.DebtWhereInput],
) usecase.IIsNextUseCase[*model.Debt, *model.DebtOrderInput, *model.DebtWhereInput] {
	return &debt.CustomerDebtIsNextUseCase{
		UC1: NewIsNextUseCase(repoIsNext),
	}
}
func NewCustomerDebtUseCase(
	repoList repository.ListModelRepository[*model.Debt, *model.DebtOrderInput, *model.DebtWhereInput],
	repoCreate repository.CreateModelRepository[*model.Debt, *model.DebtCreateInput],
	repoUpdate repository.UpdateModelRepository[*model.Debt, *model.DebtUpdateInput],
	repoFulfill repository.IDebtFullfillRepository,
	repoIsNext repository.IIsNextModelRepository[*model.Debt, *model.DebtOrderInput, *model.DebtWhereInput],
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	rlba repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	notifyTask task.IExecuteTask[*task.DebtNotifyPayload],
	mailTask task.IExecuteTask[*mail.EmailPayload],
	sk,
	debtFulfillSubjectMail,
	debtFulfillEmailTemplate,
	prodOwnerName,
	feeDesc *string,
	fee *float64,
	otpTimeout time.Duration,
) usecase.ICustomerDebtUseCase {
	return &debt.CustomerDebtUseCase{
		ICustomerDebtListUseCase:                     NewCustomerDebtListUseCase(repoList),
		ICustomerDebtCreateUseCase:                   NewCustomerDebtCreateUseCase(repoCreate, rlc, notifyTask, sk, prodOwnerName, fee, feeDesc),
		ICustomerDebtValidateCreateInputUseCase:      NewCustomerDebtValidateCreateInputUseCase(rlba, rlc, sk, prodOwnerName, fee, feeDesc),
		ICustomerConfigUseCase:                       NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		ICustomerGetUserUseCase:                      NewCustomerGetUserUseCase(rlc),
		ICustomerDebtGetFirstMineUseCase:             NewCustomerDebtGetFirstMineUseCase(repoList),
		ICustomerDebtListMineUseCase:                 NewCustomerDebtListMineUseCase(repoList),
		ICustomerDebtCancelUseCase:                   NewCustomerDebtCancelUseCase(repoUpdate, notifyTask, rlc),
		ICustomerDebtValidateCancelUseCase:           NewCustomerDebtValidateCancelUseCase(rlc),
		ICustomerDebtValidateFulfillUseCase:          NewCustomerDebtValidateFulfillUseCase(rlc, rlba),
		ICustomerDebtFulfillUseCase:                  NewCustomerDebtFulfillUseCase(mailTask, sk, prodOwnerName, feeDesc, debtFulfillSubjectMail, debtFulfillEmailTemplate, fee, otpTimeout),
		ICustomerDebtFulfillWithTokenUseCase:         NewCustomerDebtFulfillWithTokenUseCase(repoFulfill, rlc, notifyTask),
		ICustomerDebtValidateFulfillWithTokenUseCase: NewCustomerDebtValidateFulfillWithTokenUseCase(sk, prodOwnerName, feeDesc, fee),
		IIsNextUseCase:                               NewCustomerDebtIsNextUseCase(repoIsNext),
	}
}
