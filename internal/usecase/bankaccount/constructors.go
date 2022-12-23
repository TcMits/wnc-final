package bankaccount

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/internal/usecase/me"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

func NewCustomerBankAccountGetFirstUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.ICustomerBankAccountGetFirstUseCase {
	return &CustomerBankAccountGetFirstUseCase{
		bALUC: NewCustomerBankAccountListUseCase(repoList),
	}
}

func NewCustomerBankAccountListUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.ICustomerBankAccountListUseCase {
	return &CustomerBankAccountListUseCase{
		repoList: repoList,
	}
}

func NewCustomerBankAccountValidateUpdateInputUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.ICustomerBankAccountValidateUpdateInputUseCase {
	return &CustomerBankAccountValidateUpdateInputUseCase{
		bALUC: NewCustomerBankAccountListUseCase(repoList),
	}
}

func NewCustomerBankAccountUpdateUseCase(
	repoUpdate repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput],
) usecase.ICustomerBankAccountUpdateUseCase {
	return &CustomerBankAccountUpdateUseCase{
		repoUpdate: repoUpdate,
	}
}

func NewCustomerBankAccountUseCase(
	repoUpdate repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput],
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	sk *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerBankAccountUseCase {
	return &CustomerBankAccountUseCase{
		ICustomerBankAccountUpdateUseCase:              NewCustomerBankAccountUpdateUseCase(repoUpdate),
		ICustomerBankAccountValidateUpdateInputUseCase: NewCustomerBankAccountValidateUpdateInputUseCase(repoList),
		ICustomerBankAccountListUseCase:                NewCustomerBankAccountListUseCase(repoList),
		ICustomerConfigUseCase:                         config.NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		ICustomerGetUserUseCase:                        me.NewCustomerGetUserUseCase(rlc),
	}
}
