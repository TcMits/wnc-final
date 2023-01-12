package option

import (
	"context"

	"github.com/TcMits/wnc-final/internal/sse"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type (
	OptionUseCase struct {
		usecase.IGetConfigUseCase
		tpBankName string
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
	return []string{s.tpBankName}
}

func NewOptionUseCase(
	sk,
	prodOwnerName *string,
	tpBankName string,
) usecase.IOptionsUseCase {
	return &OptionUseCase{
		IGetConfigUseCase: config.NewGetConfigUseCase(sk, prodOwnerName),
	}
}
