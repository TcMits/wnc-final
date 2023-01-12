package contact

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/webapi"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type (
	CustomerContactListUseCase struct {
		repoList repository.ListModelRepository[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput]
	}
	CustomerContactUpdateUseCase struct {
		repoUpdate repository.UpdateModelRepository[*model.Contact, *model.ContactUpdateInput]
	}
	CustomerContactCreateUseCase struct {
		repoCreate repository.CreateModelRepository[*model.Contact, *model.ContactCreateInput]
	}
	CustomerContactDeleteUseCase struct {
		repoDelete repository.DeleteModelRepository[*model.Contact]
	}
	CustomerContactListMineUseCase struct {
		cLUC usecase.ICustomerContactListUseCase
	}
	CustomerContactGetFirstMineUseCase struct {
		cGFUC usecase.ICustomerContactListMineUseCase
	}
	CustomerContactValidateUpdateInputUseCase struct {
		gFMUc usecase.ICustomerContactGetFirstMineUseCase
	}
	CustomerContactValidateCreateInputUseCase struct {
		gFMUc usecase.ICustomerContactGetFirstMineUseCase
		cfUC  usecase.ICustomerConfigUseCase
		uc4   usecase.ICustomerGetFirstUseCase
		w1    webapi.ITPBankAPI
	}
	CustomerContactIsNextUseCase struct {
		iNUC usecase.IIsNextUseCase[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput]
	}
	CustomerContactUseCase struct {
		usecase.ICustomerConfigUseCase
		usecase.ICustomerGetUserUseCase
		usecase.ICustomerContactListUseCase
		usecase.ICustomerContactListMineUseCase
		usecase.ICustomerContactGetFirstMineUseCase
		usecase.ICustomerContactUpdateUseCase
		usecase.ICustomerContactValidateUpdateInputUseCase
		usecase.ICustomerContactCreateUseCase
		usecase.ICustomerContactValidateCreateInputUseCase
		usecase.ICustomerContactDeleteUseCase
		usecase.IIsNextUseCase[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput]
	}
)
