package admin

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/admin/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type meRoute struct {
	uc     usecase.IAdminMeUseCase
	logger logger.Interface
}

func RegisterMeController(handler iris.Party, l logger.Interface, uc usecase.IAdminMeUseCase) {
	h := handler.Party("/")
	route := &meRoute{
		uc:     uc,
		logger: l,
	}
	h.Use(middleware.Authenticator(uc.GetSecret(), uc.GetUser))
	h.Get("/", route.detail)
	h.Options("/", func(_ iris.Context) {})
	h.Head("/", func(_ iris.Context) {})
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
// @Router      /api/admin/v1/me/ [get]
func (s *meRoute) detail(ctx iris.Context) {
	e, err := s.uc.GetUserFromCtx(ctx)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	ctx.JSON(getResponse(e))
}
