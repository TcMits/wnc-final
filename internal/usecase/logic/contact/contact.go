package contact

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/webapi"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type (
	CustomerContactListUseCase struct {
		RepoList repository.ListModelRepository[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput]
	}
	CustomerContactUpdateUseCase struct {
		RepoUpdate repository.UpdateModelRepository[*model.Contact, *model.ContactUpdateInput]
	}
	CustomerContactCreateUseCase struct {
		RepoCreate repository.CreateModelRepository[*model.Contact, *model.ContactCreateInput]
	}
	CustomerContactDeleteUseCase struct {
		RepoDelete repository.DeleteModelRepository[*model.Contact]
	}
	CustomerContactListMineUseCase struct {
		UC1 usecase.ICustomerContactListUseCase
	}
	CustomerContactGetFirstMineUseCase struct {
		UC1 usecase.ICustomerContactListMineUseCase
	}
	CustomerContactValidateUpdateInputUseCase struct {
		UC1 usecase.ICustomerContactGetFirstMineUseCase
	}
	CustomerContactValidateCreateInputUseCase struct {
		UC1 usecase.ICustomerContactGetFirstMineUseCase
		UC2 usecase.ICustomerConfigUseCase
		UC3 usecase.ICustomerGetFirstUseCase
		W1  webapi.ITPBankAPI
	}
	CustomerContactIsNextUseCase struct {
		UC1 usecase.IIsNextUseCase[*model.Contact, *model.ContactOrderInput, *model.ContactWhereInput]
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
