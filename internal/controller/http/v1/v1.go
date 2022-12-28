package v1

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/customer"
	"github.com/TcMits/wnc-final/internal/sse"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
)

const _apiSubPath = "/api"

func NewHandler() *iris.Application {
	handler := iris.New()

	// validator
	handler.Validator = validator.New()

	// i18n
	handler.I18n.DefaultMessageFunc = func(
		langInput, langMatched, key string, args ...any) string {
		return ""
	}
	err := handler.I18n.Load("./locales/*/*")
	if err != nil {
		panic(err)
	}
	handler.I18n.SetDefault("en-US")

	return handler
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
func RegisterV1HTTPServices(
	handler iris.Party,
	// adding more usecases here
	cMeUc usecase.ICustomerMeUseCase,
	cAuthUc usecase.ICustomerAuthUseCase,
	cBankAccountUc usecase.ICustomerBankAccountUseCase,
	cStreamUc usecase.ICustomerStreamUseCase,
	cTxcUc usecase.ICustomerTransactionUseCase,
	cDUc usecase.ICustomerDebtUseCase,
	// broker
	b *sse.Broker,
	// logger
	l logger.Interface,
) {
	handler.UseRouter(recover.New())
	RegisterHealthCheckController(handler)

	// services
	h := handler.Party(_apiSubPath)
	{
		customer.RegisterCustomerServices(h, cMeUc, cAuthUc, cBankAccountUc, cStreamUc, cTxcUc, cDUc, b, l)
	}
}
