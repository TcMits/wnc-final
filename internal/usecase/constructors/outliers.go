package constructors

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/logic/outliers"
)

func NewIsNextUseCase[ModelType, OrderInput, WhereInput any](
	repo repository.IIsNextModelRepository[ModelType, OrderInput, WhereInput],
) usecase.IIsNextUseCase[ModelType, OrderInput, WhereInput] {
	return &outliers.IsNextUseCase[ModelType, OrderInput, WhereInput]{Repo: repo}
}
