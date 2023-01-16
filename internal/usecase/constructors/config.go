package constructors

import (
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/logic/config"
)

func NewGetConfigUseCase(
	secretKey *string,
	prodOwnerName *string,
) usecase.IGetConfigUseCase {
	return &config.GetConfigUseCase{
		GetSecretUseCase:           &config.GetSecretUseCase{SecretKey: secretKey},
		GetProductOwnerNameUseCase: &config.GetProductOwnerNameUseCase{Name: prodOwnerName},
	}
}

func NewEmployeeConfigUseCase(
	secretKey *string,
	prodOwnerName *string,
) usecase.IEmployeeConfigUseCase {
	return NewGetConfigUseCase(secretKey, prodOwnerName)
}
func NewAdminConfigUseCase(
	secretKey *string,
	prodOwnerName *string,
) usecase.IAdminConfigUseCase {
	return NewGetConfigUseCase(secretKey, prodOwnerName)
}
func NewPartnerConfigUseCase(
	secretKey *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.IPartnerConfigUseCase {
	return &config.PartnerConfigUseCase{
		IGetConfigUseCase: NewGetConfigUseCase(secretKey, prodOwnerName),
		FeeAmount:         fee,
		FeeDesc:           feeDesc,
	}
}

func NewCustomerConfigUseCase(
	secretKey *string,
	prodOwnerName *string,
	fee *float64,
	feeDesc *string,
) usecase.ICustomerConfigUseCase {
	uc := &config.CustomerConfigUseCase{
		IGetConfigUseCase: NewGetConfigUseCase(secretKey, prodOwnerName),
		FeeAmount:         fee,
		FeeDesc:           feeDesc,
	}
	return uc
}
