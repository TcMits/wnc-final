package option

import (
	"context"

	"github.com/TcMits/wnc-final/internal/sse"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type (
	OptionUseCase struct {
		usecase.IGetConfigUseCase
		TpBankName string
	}
	PartnerOptionUseCase struct {
	}
)

func (s *OptionUseCase) GetDebtStatus(ctx context.Context) []string {
	return model.DebtStatus
}
func (s *OptionUseCase) GetEvents(ctx context.Context) []string {
	return sse.Events
}
func (s *OptionUseCase) GetTransactionStatus(ctx context.Context) []string {
	return model.TransactionStatus
}
func (s *OptionUseCase) GetPartners(ctx context.Context) []string {
	return []string{s.TpBankName}
}
