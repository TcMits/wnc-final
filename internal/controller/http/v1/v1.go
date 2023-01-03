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

func RegisterV1HTTPServices(
	handler iris.Party,
	// adding more usecases here
	cMeUc usecase.ICustomerMeUseCase,
	cAuthUc usecase.ICustomerAuthUseCase,
	cBankAccountUc usecase.ICustomerBankAccountUseCase,
	cStreamUc usecase.ICustomerStreamUseCase,
	cTxcUc usecase.ICustomerTransactionUseCase,
	cDUc usecase.ICustomerDebtUseCase,
	cCUc usecase.ICustomerContactUseCase,
	// broker
	b *sse.Broker,
	// logger
	l logger.Interface,
) {
	handler.UseRouter(recover.New())
	RegisterHealthCheckController(handler)

	customer.RegisterCustomerServices(handler, cMeUc, cAuthUc, cBankAccountUc, cStreamUc, cTxcUc, cDUc, cCUc, b, l)
}
