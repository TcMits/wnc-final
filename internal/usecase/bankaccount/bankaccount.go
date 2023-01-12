package bankaccount

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/auth"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/internal/usecase/customer"
	"github.com/TcMits/wnc-final/internal/usecase/employee"
	"github.com/TcMits/wnc-final/internal/usecase/outliers"
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
		gFMUC: NewCustomerBankAccountGetFirstMineUseCase(repoList),
	}
}

func NewCustomerBankAccountUpdateUseCase(
	repoUpdate repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput],
) usecase.ICustomerBankAccountUpdateUseCase {
	return &CustomerBankAccountUpdateUseCase{
		repoUpdate: repoUpdate,
	}
}

func NewCustomerBankAccountListMineUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.ICustomerBankAccountListMineUseCase {
	return &CustomerBankAccountListMineUseCase{
		bALUC: NewCustomerBankAccountListUseCase(repoList),
	}
}
func NewCustomerBankAccountGetFirstMineUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.ICustomerBankAccountGetFirstMineUseCase {
	return &CustomerBankAccountGetFirstMineUseCase{
		bALMUC: NewCustomerBankAccountListMineUseCase(repoList),
	}
}
func NewCustomerBankAccountIsNextUseCase(
	repoIsNext repository.IIsNextModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.IIsNextUseCase[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput] {
	return &CustomerBankAccountIsNextUseCase{
		iNUC: outliers.NewIsNextUseCase(repoIsNext),
	}
}

func NewCustomerBankAccountUseCase(
	repoUpdate repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput],
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	repoIsNext repository.IIsNextModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
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
		ICustomerGetUserUseCase:                        auth.NewCustomerGetUserUseCase(rlc),
		ICustomerBankAccountGetFirstMineUseCase:        NewCustomerBankAccountGetFirstMineUseCase(repoList),
		ICustomerBankAccountListMineUseCase:            NewCustomerBankAccountListMineUseCase(repoList),
		ICustomerBankAccountGetFirstUseCase:            NewCustomerBankAccountGetFirstUseCase(repoList),
		IIsNextUseCase:                                 NewCustomerBankAccountIsNextUseCase(repoIsNext),
	}
}

func NewEmployeeBankAccountGetFirstUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.IEmployeeBankAccountGetFirstUseCase {
	return &EmployeeBankAccountGetFirstUseCase{
		bALUC: NewEmployeeBankAccountListUseCase(repoList),
	}
}

func NewEmployeeBankAccountListUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.IEmployeeBankAccountListUseCase {
	return &EmployeeBankAccountListUseCase{
		repoList: repoList,
	}
}

func NewEmployeeBankAccountValidateUpdateInputUseCase() usecase.IEmployeeBankAccountValidateUpdateInputUseCase {
	return &EmployeeBankAccountValidateUpdateInputUseCase{}
}

func NewEmployeeBankAccountUpdateUseCase(
	repoUpdate repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput],
) usecase.IEmployeeBankAccountUpdateUseCase {
	return &EmployeeBankAccountUpdateUseCase{
		repoUpdate: repoUpdate,
	}
}

func NewEmployeeBankAccountUseCase(
	repoUpdate repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput],
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	repoIsNext repository.IIsNextModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	rle repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
	sk *string,
	prodOwnerName *string,
) usecase.IEmployeeBankAcountUseCase {
	return &EmployeeBankAccountUseCase{
		IEmployeeBankAccountUpdateUseCase:              NewEmployeeBankAccountUpdateUseCase(repoUpdate),
		IEmployeeBankAccountValidateUpdateInputUseCase: NewEmployeeBankAccountValidateUpdateInputUseCase(),
		IEmployeeConfigUseCase:                         config.NewEmployeeConfigUseCase(sk, prodOwnerName),
		IEmployeeGetUserUseCase:                        employee.NewEmployeeGetUserUseCase(rle),
		IEmployeeBankAccountGetFirstUseCase:            NewEmployeeBankAccountGetFirstUseCase(repoList),
		IEmployeeBankAccountListUseCase:                NewEmployeeBankAccountListUseCase(repoList),
		IIsNextUseCase:                                 outliers.NewIsNextUseCase(repoIsNext),
	}
}

// partner
func NewPartnerBankAccountGetFirstUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.IPartnerBankAccountGetFirstUseCase {
	return &PartnerBankAccountGetFirstUseCase{
		bALUC: NewPartnerBankAccountListUseCase(repoList),
	}
}
func NewPartnerBankAccountRespGetFirstUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	r1 repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.IPartnerBankAccountRespGetFirstUseCase {
	return &PartnerBankAccountRespGetFirstUseCase{
		uc1: NewPartnerBankAccountListUseCase(repoList),
		uc2: customer.NewCustomerGetFirstUseCase(r1),
	}
}

func NewPartnerBankAccountListUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.IPartnerBankAccountListUseCase {
	return &PartnerBankAccountListUseCase{
		repoList: repoList,
	}
}
func NewPartnerBankAccountUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	repoIsNext repository.IIsNextModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	rlp repository.ListModelRepository[*model.Partner, *model.PartnerOrderInput, *model.PartnerWhereInput],
	r1 repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	sk *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.IPartnerBankAccountUseCase {
	return &PartnerBankAccountUseCase{
		IPartnerConfigUseCase:                  config.NewPartnerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		IPartnerGetUserUseCase:                 auth.NewPartnerGetUserUseCase(rlp),
		IPartnerBankAccountRespGetFirstUseCase: NewPartnerBankAccountRespGetFirstUseCase(repoList, r1),
		IPartnerBankAccountListUseCase:         NewPartnerBankAccountListUseCase(repoList),
		IIsNextUseCase:                         outliers.NewIsNextUseCase(repoIsNext),
	}
}
