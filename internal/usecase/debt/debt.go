package debt

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/task"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/bankaccount"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/internal/usecase/customer"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

func NewCustomerDebtListUseCase(
	repoList repository.ListModelRepository[*model.Debt, *model.DebtOrderInput, *model.DebtWhereInput],
) usecase.ICustomerDebtListUseCase {
	return &CustomerDebtListUseCase{repoList: repoList}
}
func NewCustomerDebtCreateUseCase(
	repoCreate repository.CreateModelRepository[*model.Debt, *model.DebtCreateInput],
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	taskExctor task.IExecuteTask[*task.DebtCreateNotifyPayload],
	sk *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerDebtCreateUseCase {
	return &CustomerDebtCreateUseCase{
		repoCreate:   repoCreate,
		taskExecutor: taskExctor,
		cGFUC:        customer.NewCustomerGetFirstUseCase(rlc),
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
	return &CustomerDebtValidateCreateInputUseCase{
		bAGFUC: bankaccount.NewCustomerBankAccountGetFirstUseCase(rlba),
		cGFUC:  customer.NewCustomerGetFirstUseCase(rlc),
		cfUC:   config.NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
	}
}
func NewCustomerDebtUseCase(
	repoList repository.ListModelRepository[*model.Debt, *model.DebtOrderInput, *model.DebtWhereInput],
	repoCreate repository.CreateModelRepository[*model.Debt, *model.DebtCreateInput],
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	rlba repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	taskExctor task.IExecuteTask[*task.DebtCreateNotifyPayload],
	sk *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerDebtUseCase {
	return &CustomerDebtUseCase{
		ICustomerDebtListUseCase:                NewCustomerDebtListUseCase(repoList),
		ICustomerDebtCreateUseCase:              NewCustomerDebtCreateUseCase(repoCreate, rlc, taskExctor, sk, prodOwnerName, fee, feeDesc),
		ICustomerDebtValidateCreateInputUseCase: NewCustomerDebtValidateCreateInputUseCase(rlba, rlc, sk, prodOwnerName, fee, feeDesc),
	}
}
