package constructors

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/logic/admin"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

func NewAdminGetFirstUseCase(
	repoList repository.ListModelRepository[*model.Admin, *model.AdminOrderInput, *model.AdminWhereInput],
) usecase.IAdminGetFirstUseCase {
	uc := &admin.AdminGetFirstUseCase{
		UC1: NewAdminListUseCase(repoList),
	}
	return uc
}

func NewAdminListUseCase(
	repoList repository.ListModelRepository[*model.Admin, *model.AdminOrderInput, *model.AdminWhereInput],
) usecase.IAdminListUseCase {
	uc := &admin.AdminListUseCase{
		RepoList: repoList,
	}
	return uc
}
func NewAdminUpdateUseCase(
	repoUpdate repository.UpdateModelRepository[*model.Admin, *model.AdminUpdateInput],
) usecase.IAdminUpdateUseCase {
	return &admin.AdminUpdateUseCase{
		RepoUpdate: repoUpdate,
	}
}
func NewAdminGetUserUseCase(
	repoList repository.ListModelRepository[*model.Admin, *model.AdminOrderInput, *model.AdminWhereInput],
) usecase.IAdminGetUserUseCase {
	uc := &admin.AdminGetUserUseCase{
		UC1: NewAdminGetFirstUseCase(repoList),
	}
	return uc
}
