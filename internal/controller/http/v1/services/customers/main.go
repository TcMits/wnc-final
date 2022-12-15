package customers

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/middleware"
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
	cUc usecase.IMeCustomerUseCase,
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
	}
}
