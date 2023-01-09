package stream

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/auth"
	"github.com/TcMits/wnc-final/internal/usecase/config"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

type CustomerStreamUseCase struct {
	usecase.ICustomerConfigUseCase
	usecase.ICustomerGetUserUseCase
}

func NewCustomerStreamUseCase(
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	sk *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerStreamUseCase {
	return &CustomerStreamUseCase{
		ICustomerConfigUseCase:  config.NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		ICustomerGetUserUseCase: auth.NewCustomerGetUserUseCase(rlc),
	}
}
