package constructors

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/logic/bankaccount"
	"github.com/TcMits/wnc-final/internal/webapi/tpbank"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

func NewCustomerBankAccountDeleteUseCase(
	r repository.DeleteModelRepository[*model.BankAccount],
) usecase.ICustomerBankAccountDeleteUseCase {
	return &bankaccount.CustomerBankAccountDeleteUseCase{
		Repo: r,
	}
}

func NewCustomerTPBankBankAccountGetUseCase(
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
) usecase.ICustomerTPBankBankAccountGetUseCase {
	return &bankaccount.CustomerTPBankBankAccountGetUseCase{
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

func NewCustomerBankAccountGetFirstUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.ICustomerBankAccountGetFirstUseCase {
	return &bankaccount.CustomerBankAccountGetFirstUseCase{
		UC1: NewCustomerBankAccountListUseCase(repoList),
	}
}

func NewCustomerBankAccountListUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.ICustomerBankAccountListUseCase {
	return &bankaccount.CustomerBankAccountListUseCase{
		RepoList: repoList,
	}
}

func NewCustomerBankAccountValidateUpdateInputUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.ICustomerBankAccountValidateUpdateInputUseCase {
	return &bankaccount.CustomerBankAccountValidateUpdateInputUseCase{
		UC1: NewCustomerBankAccountGetFirstMineUseCase(repoList),
	}
}

func NewCustomerBankAccountUpdateUseCase(
	repoUpdate repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput],
) usecase.ICustomerBankAccountUpdateUseCase {
	return &bankaccount.CustomerBankAccountUpdateUseCase{
		RepoUpdate: repoUpdate,
	}
}

func NewCustomerBankAccountListMineUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.ICustomerBankAccountListMineUseCase {
	return &bankaccount.CustomerBankAccountListMineUseCase{
		UC1: NewCustomerBankAccountListUseCase(repoList),
	}
}
func NewCustomerBankAccountGetFirstMineUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.ICustomerBankAccountGetFirstMineUseCase {
	return &bankaccount.CustomerBankAccountGetFirstMineUseCase{
		UC1: NewCustomerBankAccountListMineUseCase(repoList),
	}
}
func NewCustomerBankAccountIsNextUseCase(
	repoIsNext repository.IIsNextModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.IIsNextUseCase[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput] {
	return &bankaccount.CustomerBankAccountIsNextUseCase{
		UC1: NewIsNextUseCase(repoIsNext),
	}
}

func NewCustomerBankAccountUseCase(
	repoUpdate repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput],
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	repoIsNext repository.IIsNextModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	repoDelete repository.DeleteModelRepository[*model.BankAccount],
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	sk,
	prodOwnerName *string,
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
	feeDesc *string,
) usecase.ICustomerBankAccountUseCase {
	return &bankaccount.CustomerBankAccountUseCase{
		ICustomerBankAccountUpdateUseCase:              NewCustomerBankAccountUpdateUseCase(repoUpdate),
		ICustomerBankAccountValidateUpdateInputUseCase: NewCustomerBankAccountValidateUpdateInputUseCase(repoList),
		ICustomerBankAccountListUseCase:                NewCustomerBankAccountListUseCase(repoList),
		ICustomerConfigUseCase:                         NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		ICustomerGetUserUseCase:                        NewCustomerGetUserUseCase(rlc),
		ICustomerBankAccountGetFirstMineUseCase:        NewCustomerBankAccountGetFirstMineUseCase(repoList),
		ICustomerBankAccountListMineUseCase:            NewCustomerBankAccountListMineUseCase(repoList),
		ICustomerBankAccountGetFirstUseCase:            NewCustomerBankAccountGetFirstUseCase(repoList),
		ICustomerBankAccountDeleteUseCase:              NewCustomerBankAccountDeleteUseCase(repoDelete),
		ICustomerTPBankBankAccountGetUseCase: NewCustomerTPBankBankAccountGetUseCase(
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
		),
		IIsNextUseCase: NewCustomerBankAccountIsNextUseCase(repoIsNext),
	}
}

func NewEmployeeBankAccountDeleteUseCase(
	r repository.DeleteModelRepository[*model.BankAccount],
) usecase.IEmployeeBankAccountDeleteUseCase {
	return &bankaccount.EmployeeBankAccountDeleteUseCase{
		Repo: r,
	}
}

func NewEmployeeBankAccountGetFirstUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.IEmployeeBankAccountGetFirstUseCase {
	return &bankaccount.EmployeeBankAccountGetFirstUseCase{
		UC1: NewEmployeeBankAccountListUseCase(repoList),
	}
}

func NewEmployeeBankAccountListUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.IEmployeeBankAccountListUseCase {
	return &bankaccount.EmployeeBankAccountListUseCase{
		RepoList: repoList,
	}
}

func NewEmployeeBankAccountValidateUpdateInputUseCase() usecase.IEmployeeBankAccountValidateUpdateInputUseCase {
	return &bankaccount.EmployeeBankAccountValidateUpdateInputUseCase{}
}

func NewEmployeeBankAccountUpdateUseCase(
	repoUpdate repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput],
) usecase.IEmployeeBankAccountUpdateUseCase {
	return &bankaccount.EmployeeBankAccountUpdateUseCase{
		RepoUpdate: repoUpdate,
	}
}

func NewEmployeeBankAccountUseCase(
	repoUpdate repository.UpdateModelRepository[*model.BankAccount, *model.BankAccountUpdateInput],
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	repoIsNext repository.IIsNextModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	repoDelete repository.DeleteModelRepository[*model.BankAccount],
	rle repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
	sk *string,
	prodOwnerName *string,
) usecase.IEmployeeBankAcountUseCase {
	return &bankaccount.EmployeeBankAccountUseCase{
		IEmployeeBankAccountUpdateUseCase:              NewEmployeeBankAccountUpdateUseCase(repoUpdate),
		IEmployeeBankAccountValidateUpdateInputUseCase: NewEmployeeBankAccountValidateUpdateInputUseCase(),
		IEmployeeConfigUseCase:                         NewEmployeeConfigUseCase(sk, prodOwnerName),
		IEmployeeGetUserUseCase:                        NewEmployeeGetUserUseCase(rle),
		IEmployeeBankAccountGetFirstUseCase:            NewEmployeeBankAccountGetFirstUseCase(repoList),
		IEmployeeBankAccountListUseCase:                NewEmployeeBankAccountListUseCase(repoList),
		IEmployeeBankAccountDeleteUseCase:              NewEmployeeBankAccountDeleteUseCase(repoDelete),
		IIsNextUseCase:                                 NewIsNextUseCase(repoIsNext),
	}
}

// partner
func NewPartnerBankAccountGetFirstUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.IPartnerBankAccountGetFirstUseCase {
	return &bankaccount.PartnerBankAccountGetFirstUseCase{
		UC1: NewPartnerBankAccountListUseCase(repoList),
	}
}
func NewPartnerBankAccountRespGetFirstUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
	r1 repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
) usecase.IPartnerBankAccountRespGetFirstUseCase {
	return &bankaccount.PartnerBankAccountRespGetFirstUseCase{
		UC1: NewPartnerBankAccountListUseCase(repoList),
		UC2: NewCustomerGetFirstUseCase(r1),
	}
}

func NewPartnerBankAccountListUseCase(
	repoList repository.ListModelRepository[*model.BankAccount, *model.BankAccountOrderInput, *model.BankAccountWhereInput],
) usecase.IPartnerBankAccountListUseCase {
	return &bankaccount.PartnerBankAccountListUseCase{
		RepoList: repoList,
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
	return &bankaccount.PartnerBankAccountUseCase{
		IPartnerConfigUseCase:                  NewPartnerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		IPartnerGetUserUseCase:                 NewPartnerGetUserUseCase(rlp),
		IPartnerBankAccountRespGetFirstUseCase: NewPartnerBankAccountRespGetFirstUseCase(repoList, r1),
		IPartnerBankAccountListUseCase:         NewPartnerBankAccountListUseCase(repoList),
		IIsNextUseCase:                         NewIsNextUseCase(repoIsNext),
	}
}
