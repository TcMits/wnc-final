package outliers

import (
	"context"

	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
)

type IsNextUseCase[ModelType, OrderInput, WhereInput any] struct {
	repo repository.IIsNextModelRepository[ModelType, OrderInput, WhereInput]
}

func (s *IsNextUseCase[ModelType, OrderInput, WhereInput]) IsNext(ctx context.Context, limit, offset int, or OrderInput, w WhereInput) (bool, error) {
	res, err := s.repo.IsNext(ctx, limit, offset, or, w)
	if err != nil {
		return false, usecase.WrapError(err)
	}
	return res, nil
}

func NewIsNextUseCase[ModelType, OrderInput, WhereInput any](
	repo repository.IIsNextModelRepository[ModelType, OrderInput, WhereInput],
) usecase.IIsNextUseCase[ModelType, OrderInput, WhereInput] {
	return &IsNextUseCase[ModelType, OrderInput, WhereInput]{repo: repo}
}
