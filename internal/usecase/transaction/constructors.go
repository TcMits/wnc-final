package transaction

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/bankaccount"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/internal/usecase/customer"
	"github.com/TcMits/wnc-final/internal/usecase/me"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

func NewCustomerTransactionListUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.ICustomerTransactionListUseCase {
	return &CustomerTransactionListUseCase{
		repoList: repoList,
	}
}

func NewCustomerTransactionListMyTxcUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.ICustomerTransactionListMyTxcUseCase {
	return &CustomerTransactionListMyTxcUseCase{
		tLUC: NewCustomerTransactionListUseCase(repoList),
	}
}
func NewCustomerTransactionGetFirstMyTxUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
) usecase.ICustomerTransactionGetFirstMyTxUseCase {
	return &CustomerTransactionGetFirstMyTxUseCase{
		tLMTUC: NewCustomerTransactionListMyTxcUseCase(repoList),
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
	repoCreate repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput],
	rLBa repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	rUBa repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput],
) usecase.ICustomerTransactionCreateUseCase {
	return &CustomerTransactionCreateUseCase{
		repoCreate: repoCreate,
		bKGFUC:     bankaccount.NewCustomerBankAccountGetFirstUseCase(rLBa),
		bKUUC:      bankaccount.NewCustomerBankAccountUpdateUseCase(rUBa),
	}
}

func NewCustomerTransactionValidateCreateInputUseCase(
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
	rlba repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	sk *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerTransactionValidateCreateInputUseCase {
	return &CustomerTransactionValidateCreateInputUseCase{
		tLUC:   NewCustomerTransactionListUseCase(repoList),
		bAGFUC: bankaccount.NewCustomerBankAccountGetFirstUseCase(rlba),
		cGFUC:  customer.NewCustomerGetFirstUseCase(rlc),
		cfUC:   config.NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
	}
}

func NewCustomerTransactionConfirmSuccessUseCase(
	repoCreate repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput],
	repoUpdate repository.UpdateModelRepository[*model.Transaction, *model.TransactionUpdateInput],
	rlba repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	rUBa repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput],
	sk *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerTransactionConfirmSuccessUseCase {
	return &CustomerTransactionConfirmSuccessUseCase{
		cfUC:   config.NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		bAUUC:  bankaccount.NewCustomerBankAccountUpdateUseCase(rUBa),
		bAGFUC: bankaccount.NewCustomerBankAccountGetFirstUseCase(rlba),
		tCUC:   NewCustomerTransactionCreateUseCase(repoCreate, rlba, rUBa),
		tUUC:   NewCustomerTransactionUpdateUseCase(repoUpdate),
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

func NewCustomerTransactionUseCase(
	repoCreate repository.CreateModelRepository[*model.Transaction, *model.TransactionCreateInput],
	repoList repository.ListModelRepository[*model.Transaction, *model.TransactionOrderInput, *model.TransactionWhereInput],
	repoUpdate repository.UpdateModelRepository[*model.Transaction, *model.TransactionUpdateInput],
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	rlba repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	rUBa repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput],
	sk *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerTransactionUseCase {
	return &CustomerTransactionUseCase{
		ICustomerTransactionCreateUseCase:              NewCustomerTransactionCreateUseCase(repoCreate, rlba, rUBa),
		ICustomerTransactionValidateCreateInputUseCase: NewCustomerTransactionValidateCreateInputUseCase(repoList, rlba, rlc, sk, prodOwnerName, fee, feeDesc),
		ICustomerTransactionListUseCase:                NewCustomerTransactionListUseCase(repoList),
		ICustomerConfigUseCase:                         config.NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		ICustomerGetUserUseCase:                        me.NewCustomerGetUserUseCase(rlc),
		ICustomerTransactionConfirmSuccessUseCase:      NewCustomerTransactionConfirmSuccessUseCase(repoCreate, repoUpdate, rlba, rUBa, sk, prodOwnerName, fee, feeDesc),
		ICustomerTransactionListMyTxcUseCase:           NewCustomerTransactionListMyTxcUseCase(repoList),
		ICustomerTransactionGetFirstMyTxUseCase:        NewCustomerTransactionGetFirstMyTxUseCase(repoList),
	}
}
