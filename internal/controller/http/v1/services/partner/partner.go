package partner

import (
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

const (
	_v1SubPath = "/api/partner/v1"
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
// @BasePath /api/partner/v1
func RegisterServices(
	handler iris.Party,
	// adding more usecases here
	uc1 usecase.IPartnerTransactionUseCase,
	uc2 usecase.IPartnerAuthUseCase,
	uc3 usecase.IPartnerBankAccountUseCase,
	// logger
	l logger.Interface,
) {
	h := handler.Party(_v1SubPath)
	// routes
	{
		RegisterDocsController(h, l)
		RegisterAuthController(h, l, uc2)
		RegisterTransactionController(h, l, uc1)
		RegisterBankAccountController(h, l, uc3)
	}
}
