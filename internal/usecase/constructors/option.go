package constructors

import (
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/internal/usecase/logic/option"
)

func NewOptionUseCase(
	sk,
	prodOwnerName *string,
	tpBankName string,
) usecase.IOptionsUseCase {
	return &option.OptionUseCase{
		IGetConfigUseCase: NewGetConfigUseCase(sk, prodOwnerName),
		TpBankName:        tpBankName,
	}
}
