package admin

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type (
	AdminListUseCase struct {
		repoList repository.ListModelRepository[*model.Admin, *model.AdminOrderInput, *model.AdminWhereInput]
	}
	AdminUpdateUseCase struct {
		repoUpdate repository.UpdateModelRepository[*model.Admin, *model.AdminUpdateInput]
	}
	AdminGetFirstUseCase struct {
		eLUC usecase.IAdminListUseCase
	}
)

func NewAdminGetFirstUseCase(
	repoList repository.ListModelRepository[*model.Admin, *model.AdminOrderInput, *model.AdminWhereInput],
) usecase.IAdminGetFirstUseCase {
	uc := &AdminGetFirstUseCase{
		eLUC: repoList,
	}
	return uc
}

func NewAdminListUseCase(
	repoList repository.ListModelRepository[*model.Admin, *model.AdminOrderInput, *model.AdminWhereInput],
) usecase.IAdminListUseCase {
	uc := &AdminListUseCase{
		repoList: repoList,
	}
	return uc
}
func NewAdminUpdateUseCase(
	repoUpdate repository.UpdateModelRepository[*model.Admin, *model.AdminUpdateInput],
) usecase.IAdminUpdateUseCase {
	return &AdminUpdateUseCase{
		repoUpdate: repoUpdate,
	}
}

func (uc *AdminGetFirstUseCase) GetFirst(ctx context.Context, o *model.AdminOrderInput, w *model.AdminWhereInput) (*model.Admin, error) {
	l, of := 1, 0
	entities, err := uc.eLUC.List(ctx, &l, &of, o, w)
	if err != nil {
		return nil, err
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}

func (uc *AdminListUseCase) List(ctx context.Context, limit, offset *int, o *model.AdminOrderInput, w *model.AdminWhereInput) ([]*model.Admin, error) {
	entities, err := uc.repoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.admin.admin.AdminListUseCase.List: %s", err))
	}
	return entities, nil
}

func (s *AdminUpdateUseCase) Update(ctx context.Context, e *model.Admin, i *model.AdminUpdateInput) (*model.Admin, error) {
	e, err := s.repoUpdate.Update(ctx, e, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.admin.admin.AdminUpdateUseCase.Update: %s", err))
	}
	return e, nil
}
