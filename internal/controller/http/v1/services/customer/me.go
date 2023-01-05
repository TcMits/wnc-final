package customer

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/customer/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type meRoute struct {
	uc     usecase.ICustomerMeUseCase
	logger logger.Interface
}

func RegisterMeController(handler iris.Party, l logger.Interface, uc usecase.ICustomerMeUseCase) {
	route := &meRoute{
		uc:     uc,
		logger: l,
	}
	handler.Use(middleware.Authenticator(uc.GetSecret(), uc.GetUser))
	handler.Post("/change-password", route.changePassword)
	handler.Get("/", route.detail)
	handler.Options("/change-password", func(_ iris.Context) {})
	handler.Options("/", func(_ iris.Context) {})
	handler.Head("/change-password", func(_ iris.Context) {})
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

// @Summary     Change password
// @Description Change password
// @ID          change-password
// @Tags  	    change-password
// @Accept      json
// @Produce     json
// @Param       payload body changePasswordReq true "Change password"
// @Success     200 {object} meResponse
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /me/change-password [post]
func (s *meRoute) changePassword(ctx iris.Context) {
	request := new(changePasswordReq)
	if err := ctx.ReadJSON(request); err != nil {
		handleBindingError(ctx, err, s.logger, request, nil)
		return
	}
	i := &model.CustomerChangePasswordInput{
		OldPassword:     request.OldPassword,
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
	}
	i, err := s.uc.ValidateChangePassword(ctx, i)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	e, err := s.uc.ChangePassword(ctx, i)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	ctx.JSON(getResponse(e))
}
