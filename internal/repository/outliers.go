package repository

import (
	"context"
	"fmt"
)

type IsNextModelRepository[ModelType, OrderInputType, WhereInputType any] struct {
	repo ListModelRepository[ModelType, OrderInputType, WhereInputType]
}

func (s *IsNextModelRepository[ModelType, OrderInputType, WhereInputType]) IsNext(ctx context.Context, limit, offset int, or OrderInputType, w WhereInputType) (bool, error) {
	limit = limit + 1
	entities, err := s.repo.List(ctx, &limit, &offset, or, w)
	if err != nil {
		return false, fmt.Errorf("internal.repository.outliers.NextModelRepository.IsNext: %s", err)
	}
	return len(entities) > 0, nil
}

func getIsNextModelRepostiory[ModelType, OrderInputType, WhereInputType any](
	repo ListModelRepository[ModelType, OrderInputType, WhereInputType],
) IIsNextModelRepository[ModelType, OrderInputType, WhereInputType] {
	return &IsNextModelRepository[ModelType, OrderInputType, WhereInputType]{repo: repo}
}
