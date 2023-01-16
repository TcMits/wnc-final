package constructors

import (
	"github.com/TcMits/wnc-final/internal/repository"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/logic/stream"
	"github.com/TcMits/wnc-final/pkg/entity/model"
)

func NewCustomerStreamUseCase(
	rlc repository.ListModelRepository[*model.Customer, *model.CustomerOrderInput, *model.CustomerWhereInput],
	sk *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerStreamUseCase {
	return &stream.CustomerStreamUseCase{
		ICustomerConfigUseCase:  NewCustomerConfigUseCase(sk, prodOwnerName, fee, feeDesc),
		ICustomerGetUserUseCase: NewCustomerGetUserUseCase(rlc),
	}
}
