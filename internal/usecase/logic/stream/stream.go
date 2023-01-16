package stream

import (
	"github.com/TcMits/wnc-final/internal/usecase"
)

type CustomerStreamUseCase struct {
	usecase.ICustomerConfigUseCase
	usecase.ICustomerGetUserUseCase
}
