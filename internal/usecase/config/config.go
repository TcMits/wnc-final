package config

import (
	"github.com/TcMits/wnc-final/internal/usecase"
)

type (
	GetSecretUseCase struct {
		secretKey *string
	}
	GetProductOwnerNameUseCase struct {
		name *string
	}

	GetConfigUseCase struct {
		*GetSecretUseCase
		*GetProductOwnerNameUseCase
	}

	CustomerConfigUseCase struct {
		usecase.IGetConfigUseCase
		usecase.ICustomerGetUserUseCase
		feeAmount *float64
		feeDesc   *string
	}
)

func NewGetConfigUseCase(
	secretKey *string,
	prodOwnerName *string,
) usecase.IGetConfigUseCase {
	return &GetConfigUseCase{
		GetSecretUseCase:           &GetSecretUseCase{secretKey: secretKey},
		GetProductOwnerNameUseCase: &GetProductOwnerNameUseCase{name: prodOwnerName},
	}
}

func NewCustomerConfigUseCase(
	secretKey *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerConfigUseCase {
	uc := &CustomerConfigUseCase{
		IGetConfigUseCase: NewGetConfigUseCase(secretKey, prodOwnerName),
		feeAmount:         fee,
		feeDesc:           feeDesc,
	}
	return uc
}

func (uc *GetSecretUseCase) GetSecret() *string {
	return uc.secretKey
}

func (uc *GetProductOwnerNameUseCase) GetProductOwnerName() *string {
	return uc.name
}

func (uc *CustomerConfigUseCase) GetFeeAmount() *float64 {
	return uc.feeAmount
}
func (uc *CustomerConfigUseCase) GetFeeDesc() *string {
	return uc.feeDesc
}
