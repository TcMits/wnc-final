package employee

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/employee/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type meRoute struct {
	uc     usecase.IEmployeeMeUseCase
	logger logger.Interface
}

func RegisterMeController(handler iris.Party, l logger.Interface, uc usecase.IEmployeeMeUseCase) {
	route := &meRoute{
		uc:     uc,
		logger: l,
	}
	handler.Use(middleware.Authenticator(uc.GetSecret(), uc.GetUser))
	handler.Get("/", route.detail)
	handler.Options("/", func(_ iris.Context) {})
	handler.Head("/", func(_ iris.Context) {})
}

// @Summary     Get profile
// @Description Get profile
// @ID          me
// @Tags  	    Me
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Success     200 {object} meResponse
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /me/ [get]
func (s *meRoute) detail(ctx iris.Context) {
	e, err := s.uc.GetUserFromCtx(ctx)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	ctx.JSON(getResponse(e))
}
