package constructors

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/logic/contact"
	"github.com/TcMits/wnc-final/internal/webapi/tpbank"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

func NewCustomerContactListUseCase(
	repoList repository.ListModelRepository[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput],
) usecase.ICustomerContactListUseCase {
	return &contact.CustomerContactListUseCase{RepoList: repoList}
}

func NewCustomerContactUpdateUseCase(
	repoUpdate repository.UpdateModelRepository[*model.Contact, *model.ContactUpdateInput],
) usecase.ICustomerContactUpdateUseCase {
	return &contact.CustomerContactUpdateUseCase{RepoUpdate: repoUpdate}
}

func NewCustomerContactCreateUseCase(
	repoCreate repository.CreateModelRepository[*model.Contact, *model.ContactCreateInput],
) usecase.ICustomerContactCreateUseCase {
	return &contact.CustomerContactCreateUseCase{RepoCreate: repoCreate}
}

func NewCustomerContactDeleteUseCase(
	repoDelete repository.DeleteModelRepository[*model.Contact],
) usecase.ICustomerContactDeleteUseCase {
	return &contact.CustomerContactDeleteUseCase{RepoDelete: repoDelete}
}

func NewCustomerContactListMineUseCase(
	repoList repository.ListModelRepository[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput],
) usecase.ICustomerContactListMineUseCase {
	return &contact.CustomerContactListMineUseCase{
		UC1: NewCustomerContactListUseCase(repoList),
	}
}

func NewCustomerContactGetFirstMineUseCase(
	repoList repository.ListModelRepository[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput],
) usecase.ICustomerContactGetFirstMineUseCase {
	return &contact.CustomerContactGetFirstMineUseCase{
		UC1: NewCustomerContactListMineUseCase(repoList),
	}
}

func NewCustomerContactValidateUpdateInputUseCase(
	repoList repository.ListModelRepository[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput],
) usecase.ICustomerContactValidateUpdateInputUseCase {
	return &contact.CustomerContactValidateUpdateInputUseCase{UC1: NewCustomerContactGetFirstMineUseCase(repoList)}
}
func NewCustomerContactValidateCreateInputUseCase(
	repoList repository.ListModelRepository[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput],
	r1 repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
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
) usecase.ICustomerContactValidateCreateInputUseCase {
	return &contact.CustomerContactValidateCreateInputUseCase{
		UC1: NewCustomerContactGetFirstMineUseCase(repoList),
		UC2: NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		UC3: NewCustomerGetFirstUseCase(r1),
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
func NewCustomerContactIsNextUseCase(
	repoIsNext repository.IIsNextModelRepository[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput],
) usecase.IIsNextUseCase[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput] {
	return &contact.CustomerContactIsNextUseCase{
		UC1: NewIsNextUseCase(repoIsNext),
	}
}

func NewCustomerContactUseCase(
	repoList repository.ListModelRepository[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput],
	repoUpdate repository.UpdateModelRepository[*model.Contact, *model.ContactUpdateInput],
	repoCreate repository.CreateModelRepository[*model.Contact, *model.ContactCreateInput],
	repoDelete repository.DeleteModelRepository[*model.Contact],
	repoIsNext repository.IIsNextModelRepository[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput],
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
) usecase.ICustomerContactUseCase {
	return &contact.CustomerContactUseCase{
		ICustomerConfigUseCase:                     NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		ICustomerGetUserUseCase:                    NewCustomerGetUserUseCase(rlc),
		ICustomerContactListUseCase:                NewCustomerContactListUseCase(repoList),
		ICustomerContactListMineUseCase:            NewCustomerContactListMineUseCase(repoList),
		ICustomerContactGetFirstMineUseCase:        NewCustomerContactGetFirstMineUseCase(repoList),
		ICustomerContactUpdateUseCase:              NewCustomerContactUpdateUseCase(repoUpdate),
		ICustomerContactValidateUpdateInputUseCase: NewCustomerContactValidateUpdateInputUseCase(repoList),
		ICustomerContactCreateUseCase:              NewCustomerContactCreateUseCase(repoCreate),
		ICustomerContactValidateCreateInputUseCase: NewCustomerContactValidateCreateInputUseCase(repoList, rlc, sk, prodOwnerName, feeDesc,
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
		ICustomerContactDeleteUseCase: NewCustomerContactDeleteUseCase(repoDelete),
		IIsNextUseCase:                NewCustomerContactIsNextUseCase(repoIsNext),
	}
}
