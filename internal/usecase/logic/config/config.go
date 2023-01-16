package config

import (
	"github.com/TcMits/wnc-final/internal/usecase"
)

type (
	GetSecretUseCase struct {
		SecretKey *string
	}
	GetProductOwnerNameUseCase struct {
		Name *string
	}

	GetConfigUseCase struct {
		*GetSecretUseCase
		*GetProductOwnerNameUseCase
	}

	CustomerConfigUseCase struct {
		usecase.IGetConfigUseCase
		FeeAmount *float64
		FeeDesc   *string
	}
	PartnerConfigUseCase struct {
		usecase.IGetConfigUseCase
		FeeAmount *float64
		FeeDesc   *string
	}
)

func NewGetConfigUseCase(
	secretKey *string,
	prodOwnerName *string,
) usecase.IGetConfigUseCase {
	return &GetConfigUseCase{
		GetSecretUseCase:           &GetSecretUseCase{SecretKey: secretKey},
		GetProductOwnerNameUseCase: &GetProductOwnerNameUseCase{Name: prodOwnerName},
	}
}

func (s *GetSecretUseCase) GetSecret() *string {
	return s.SecretKey
}

func (s *GetProductOwnerNameUseCase) GetProductOwnerName() *string {
	return s.Name
}

func (s *CustomerConfigUseCase) GetFeeAmount() *float64 {
	return s.FeeAmount
}
func (s *CustomerConfigUseCase) GetFeeDesc() *string {
	return s.FeeDesc
}

func (s *PartnerConfigUseCase) GetFeeAmount() *float64 {
	return s.FeeAmount
}
func (s *PartnerConfigUseCase) GetFeeDesc() *string {
	return s.FeeDesc
}
