package contact

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/auth"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/internal/usecase/outliers"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

func NewCustomerContactListUseCase(
	repoList repository.ListModelRepository[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput],
) usecase.ICustomerContactListUseCase {
	return &CustomerContactListUseCase{repoList: repoList}
}

func NewCustomerContactUpdateUseCase(
	repoUpdate repository.UpdateModelRepository[*model.Contact, *model.ContactUpdateInput],
) usecase.ICustomerContactUpdateUseCase {
	return &CustomerContactUpdateUseCase{repoUpdate: repoUpdate}
}

func NewCustomerContactCreateUseCase(
	repoCreate repository.CreateModelRepository[*model.Contact, *model.ContactCreateInput],
) usecase.ICustomerContactCreateUseCase {
	return &CustomerContactCreateUseCase{repoCreate: repoCreate}
}

func NewCustomerContactDeleteUseCase(
	repoDelete repository.DeleteModelRepository[*model.Contact],
) usecase.ICustomerContactDeleteUseCase {
	return &CustomerContactDeleteUseCase{repoDelete: repoDelete}
}

func NewCustomerContactListMineUseCase(
	repoList repository.ListModelRepository[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput],
) usecase.ICustomerContactListMineUseCase {
	return &CustomerContactListMineUseCase{
		cLUC: NewCustomerContactListUseCase(repoList),
	}
}

func NewCustomerContactGetFirstMineUseCase(
	repoList repository.ListModelRepository[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput],
) usecase.ICustomerContactGetFirstMineUseCase {
	return &CustomerContactGetFirstMineUseCase{
		cGFUC: NewCustomerContactListMineUseCase(repoList),
	}
}

func NewCustomerContactValidateUpdateInputUseCase(
	repoList repository.ListModelRepository[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput],
) usecase.ICustomerContactValidateUpdateInputUseCase {
	return &CustomerContactValidateUpdateInputUseCase{cGFUC: NewCustomerContactListMineUseCase(repoList)}
}
func NewCustomerContactValidateCreateInputUseCase(
	repoList repository.ListModelRepository[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput],
	sk,
	prodOwnerName,
	feeDesc *string,
	fee *float64,
) usecase.ICustomerContactValidateCreateInputUseCase {
	return &CustomerContactValidateCreateInputUseCase{
		cGFUC: NewCustomerContactListMineUseCase(repoList),
		cfUC:  config.NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
	}
}
func NewCustomerBankAccountIsNextUseCase(
	repoIsNext repository.IIsNextModelRepository[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput],
) usecase.IIsNextUseCase[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput] {
	return &CustomerContactIsNextUseCase{
		iNUC: outliers.NewIsNextUseCase(repoIsNext),
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
	fee *float64,
) usecase.ICustomerContactUseCase {
	return &CustomerContactUseCase{
		ICustomerConfigUseCase:                     config.NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		ICustomerGetUserUseCase:                    auth.NewCustomerGetUserUseCase(rlc),
		ICustomerContactListUseCase:                NewCustomerContactListUseCase(repoList),
		ICustomerContactListMineUseCase:            NewCustomerContactListMineUseCase(repoList),
		ICustomerContactGetFirstMineUseCase:        NewCustomerContactGetFirstMineUseCase(repoList),
		ICustomerContactUpdateUseCase:              NewCustomerContactUpdateUseCase(repoUpdate),
		ICustomerContactValidateUpdateInputUseCase: NewCustomerContactValidateUpdateInputUseCase(repoList),
		ICustomerContactCreateUseCase:              NewCustomerContactCreateUseCase(repoCreate),
		ICustomerContactValidateCreateInputUseCase: NewCustomerContactValidateCreateInputUseCase(repoList, sk, prodOwnerName, feeDesc, fee),
		ICustomerContactDeleteUseCase:              NewCustomerContactDeleteUseCase(repoDelete),
		IIsNextUseCase:                             NewCustomerBankAccountIsNextUseCase(repoIsNext),
	}
}
