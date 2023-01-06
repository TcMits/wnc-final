package option

import (
	"context"

	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type OptionUseCase struct{}

func (s *OptionUseCase) GetDebtStatus(ctx context.Context) []string {
	return model.DebtStatus
}

func NewOptionUseCase() usecase.IOptionsUseCase {
	return new(OptionUseCase)
}
