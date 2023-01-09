package employee

import (
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

const (
	_employeeV1SubPath = "/api/employee/v1"
)

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
// @BasePath /api/employee/v1
func RegisterEmployeeServices(
	handler iris.Party,
	// adding more usecases here
	uc1 usecase.IEmployeeCustomerUseCase,
	uc2 usecase.IEmployeeAuthUseCase,
	uc3 usecase.IEmployeeMeUseCase,
	uc4 usecase.IEmployeeBankAcountUseCase,
	uc5 usecase.IEmployeeTransactionUseCase,
	// logger
	l logger.Interface,
) {
	h := handler.Party(_employeeV1SubPath)
	// routes
	{
		RegisterDocsController(h, l)
		RegisterCustomerController(h, l, uc1)
		RegisterAuthController(h, l, uc2)
		RegisterBankAccountController(h, l, uc4)
		RegisterTransactionController(h, l, uc5)
		h = h.Party(
			"/me",
		)
		{
			RegisterMeController(h, l, uc3)
		}
	}
}
