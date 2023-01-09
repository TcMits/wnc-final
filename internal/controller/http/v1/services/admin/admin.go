package admin

import (
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

const (
	_adminV1SubPath = "/api/admin/v1"
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
// @BasePath /api/admin/v1
func RegisterServices(
	handler iris.Party,
	// adding more usecases here
	uc1 usecase.IAdminAuthUseCase,
	uc2 usecase.IAdminMeUseCase,
	// logger
	l logger.Interface,
) {
	h := handler.Party(_adminV1SubPath)
	// routes
	{
		RegisterDocsController(h, l)
		RegisterAuthController(h, l, uc1)
		h = h.Party(
			"/me",
		)
		{
			RegisterMeController(h, l, uc2)
		}
	}
}
