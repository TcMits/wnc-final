package v1

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/admin"
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/customer"
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/customer/middleware"
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/employee"
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/partner"
	"github.com/TcMits/wnc-final/internal/sse"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/cors"
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
	// adding more customer usecases here
	customerUc1 usecase.ICustomerMeUseCase,
	customerUc2 usecase.ICustomerAuthUseCase,
	customerUc3 usecase.ICustomerBankAccountUseCase,
	customerUc4 usecase.ICustomerStreamUseCase,
	customerUc5 usecase.ICustomerTransactionUseCase,
	customerUc6 usecase.ICustomerDebtUseCase,
	customerUc7 usecase.ICustomerContactUseCase,
	customerUc8 usecase.IOptionsUseCase,
	// adding more employee usecases here
	employeeUc1 usecase.IEmployeeAuthUseCase,
	employeeUc2 usecase.IEmployeeMeUseCase,
	employeeUc3 usecase.IEmployeeCustomerUseCase,
	employeeUc4 usecase.IEmployeeBankAcountUseCase,
	employeeUc5 usecase.IEmployeeTransactionUseCase,
	// adding more admin usecases here
	adminUc1 usecase.IAdminMeUseCase,
	adminUc2 usecase.IAdminAuthUseCase,
	adminUc3 usecase.IAdminTransactionUseCase,
	adminUc4 usecase.IAdminEmployeeUseCase,
	// adding more partner usecases here
	partnerUc1 usecase.IPartnerAuthUseCase,
	partnerUc2 usecase.IPartnerTransactionUseCase,
	// broker
	b *sse.Broker,
	// logger
	l logger.Interface,
) {
	handler.UseRouter(recover.New())
	handler.UseRouter(middleware.Logger(l))
	handler.UseRouter(cors.New().Handler())
	RegisterHealthCheckController(handler)

	customer.RegisterServices(handler, customerUc1, customerUc2, customerUc3, customerUc4, customerUc5, customerUc6, customerUc7, customerUc8, b, l)
	employee.RegisterServices(handler, employeeUc3, employeeUc1, employeeUc2, employeeUc4, employeeUc5, l)
	admin.RegisterServices(handler, adminUc2, adminUc1, adminUc3, adminUc4, l)
	partner.RegisterServices(handler, partnerUc2, partnerUc1, l)
}
