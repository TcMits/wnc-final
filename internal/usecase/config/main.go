package config

import (
	"github.com/TcMits/wnc-final/internal/usecase"
)

type (
	CustomerGetSecretUseCase struct {
		secretKey *string
	}

	CustomerConfigUseCase struct {
		*CustomerGetSecretUseCase
	}
)

func NewCustomerConfigUseCase(
	secretKey *string,
) usecase.ICustomerConfigUseCase {
	uc := &CustomerConfigUseCase{
		CustomerGetSecretUseCase: &CustomerGetSecretUseCase{secretKey: secretKey},
	}
	return uc
}

func (uc *CustomerGetSecretUseCase) GetSecret() (*string, error) {
	return uc.secretKey, nil
}
