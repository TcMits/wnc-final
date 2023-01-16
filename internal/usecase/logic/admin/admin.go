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
		RepoList repository.ListModelRepository[*model.Admin, *model.AdminOrderInput, *model.AdminWhereInput]
	}
	AdminUpdateUseCase struct {
		RepoUpdate repository.UpdateModelRepository[*model.Admin, *model.AdminUpdateInput]
	}
	AdminGetFirstUseCase struct {
		UC1 usecase.IAdminListUseCase
	}
	AdminGetUserUseCase struct {
		UC1 usecase.IAdminGetFirstUseCase
	}
)

func (s *AdminGetUserUseCase) GetUser(ctx context.Context, input map[string]any) (any, error) {
	usernameAny, ok := input["username"]
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("username is required"))
	}
	username, ok := usernameAny.(string)
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("wrong type of username, expected type of string, not %T", username))
	}
	u, err := s.UC1.GetFirst(ctx, nil, &model.AdminWhereInput{
		Username: &username,
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *AdminGetFirstUseCase) GetFirst(ctx context.Context, o *model.AdminOrderInput, w *model.AdminWhereInput) (*model.Admin, error) {
	l, of := 1, 0
	entities, err := s.UC1.List(ctx, &l, &of, o, w)
	if err != nil {
		return nil, err
	}
	if len(entities) > 0 {
		return entities[0], nil
	}
	return nil, nil
}

func (s *AdminListUseCase) List(ctx context.Context, limit, offset *int, o *model.AdminOrderInput, w *model.AdminWhereInput) ([]*model.Admin, error) {
	entities, err := s.RepoList.List(ctx, limit, offset, o, w)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.admin.admin.AdminListUseCase.List: %s", err))
	}
	return entities, nil
}

func (s *AdminUpdateUseCase) Update(ctx context.Context, e *model.Admin, i *model.AdminUpdateInput) (*model.Admin, error) {
	e, err := s.RepoUpdate.Update(ctx, e, i)
	if err != nil {
		return nil, usecase.WrapError(fmt.Errorf("internal.usecase.logic.admin.admin.AdminUpdateUseCase.Update: %s", err))
	}
	return e, nil
}
