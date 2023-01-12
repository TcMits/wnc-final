package user

import (
	"context"
	"fmt"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/employee"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type (
	EmployeeGetUserUseCase struct {
		gFUC usecase.IEmployeeGetFirstUseCase
	}
)

func NewEmployeeGetUserUseCase(
	repoList repository.ListModelRepository[*model.Employee, *model.EmployeeOrderInput, *model.EmployeeWhereInput],
) usecase.IEmployeeGetUserUseCase {
	uc := &EmployeeGetUserUseCase{
		gFUC: employee.NewEmployeeGetFirstUseCase(repoList),
	}
	return uc
}

func (s *EmployeeGetUserUseCase) GetUser(ctx context.Context, input map[string]any) (any, error) {
	usernameAny, ok := input["username"]
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("username is required"))
	}
	username, ok := usernameAny.(string)
	if !ok {
		return nil, usecase.WrapError(fmt.Errorf("wrong type of username, expected type of string, not %T", username))
	}
	u, err := s.gFUC.GetFirst(ctx, nil, &model.EmployeeWhereInput{
		Username: &username,
	})
	if err != nil {
		return nil, err
	}
	return u, nil
}
