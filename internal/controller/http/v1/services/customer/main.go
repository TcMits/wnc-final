package customer

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/customer/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/cors"
)

const (
	_customerV1SubPath = "/customer/v1"
)

func RegisterCustomerServices(
	handler iris.Party,
	// adding more usecases here
	cUc usecase.ICustomerMeUseCase,
	aUc usecase.ICustomerAuthUseCase,
	cbac usecase.ICustomerBankAccountUseCase,
	// logger
	l logger.Interface,
) {
	// HTTP middlewares
	h := handler.Party(
		_customerV1SubPath,
		cors.New().Handler(),
		middleware.Logger(l),
	)
	// routes
	{
		RegisterDocsController(h, l)
		RegisterAuthController(h, l, aUc)
		h := h.Party(
			"/me",
		)
		{
			RegisterMeController(h, l, cUc)
			RegisterBankAccountController(h, l, cbac)
		}
	}
}
