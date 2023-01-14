package customer

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/customer/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type authRoute struct {
	uc     usecase.ICustomerAuthUseCase
	logger logger.Interface
}

func RegisterAuthController(handler iris.Party, l logger.Interface, uc usecase.ICustomerAuthUseCase) {
	h := handler.Party("/")
	route := &authRoute{
		uc:     uc,
		logger: l,
	}
	h.Post("/forget-password", route.forgetPassword)
	h.Post("/change-password-with-token", route.changePassword)
	h.Post("/token", route.renewToken)
	h.Post("/login", route.login)
	h.Delete("/login", middleware.Authenticator(uc.GetSecret(), uc.GetUser), route.logout)
	h.Options("/forget-password", func(_ iris.Context) {})
	h.Options("/change-password-with-token", func(_ iris.Context) {})
	h.Options("/login", func(_ iris.Context) {})
	h.Options("/token", func(_ iris.Context) {})
	h.Head("/forget-password", func(_ iris.Context) {})
	h.Head("/change-password-with-token", func(_ iris.Context) {})
	h.Head("/login", func(_ iris.Context) {})
	h.Head("/token", func(_ iris.Context) {})
}

// @Summary     Login
// @Description Login
// @ID          login
// @Tags  	    Authentication
// @Accept      json
// @Produce     json
// @Param       payload body loginRequest true "Login"
// @Success     200 {object} tokenPairResponse
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /api/customer/v1/login [post]
func (r *authRoute) login(ctx iris.Context) {
	request := new(loginRequest)
	if err := ctx.ReadJSON(request); err != nil {
		handleBindingError(ctx, err, r.logger, request, nil)
		return
	}
	i := &model.CustomerLoginInput{
		Username: request.Username,
		Password: request.Password,
	}
	validatedData, err := r.uc.ValidateLoginInput(ctx, i)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	res, err := r.uc.Login(ctx, validatedData)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	ctx.JSON(getResponse(res))
}

// @Summary     Logout
// @Description Logout
// @ID          logout
// @Tags  	    Authentication
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Success     204  ""
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /api/customer/v1/login [Delete]
func (r *authRoute) logout(ctx iris.Context) {
	err := r.uc.Logout(ctx)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	ctx.StatusCode(iris.StatusNoContent)
}

// @Summary     Renew token
// @Description Renew token
// @ID          renewtoken
// @Tags  	    Authentication
// @Accept      json
// @Produce     json
// @Param       payload body renewTokenRequest true "Renew token"
// @Success     200 {object} tokenPairResponse
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /api/customer/v1/token [post]
func (r *authRoute) renewToken(ctx iris.Context) {
	request := new(renewTokenRequest)
	if err := ctx.ReadJSON(request); err != nil {
		handleBindingError(ctx, err, r.logger, request, nil)
		return
	}
	res, err := r.uc.RenewToken(ctx, request.RefreshToken)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(getResponse(res))
}

// @Summary     Forget password
// @Description Forget password
// @ID          forget-password
// @Tags  	    Authentication
// @Accept      json
// @Produce     json
// @Param       payload body forgetPasswordReq true "Forget password"
// @Success     200 {object} forgetPasswordResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /api/customer/v1/forget-password [post]
func (s *authRoute) forgetPassword(ctx iris.Context) {
	request := new(forgetPasswordReq)
	if err := ctx.ReadJSON(request); err != nil {
		handleBindingError(ctx, err, s.logger, request, nil)
		return
	}
	i := &model.CustomerForgetPasswordInput{
		Email: request.Email,
	}
	i, err := s.uc.ValidateForgetPassword(ctx, i)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	res, err := s.uc.ForgetPassword(ctx, i)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	ctx.JSON(getResponse(res))
}

// @Summary     Change password with token
// @Description Change password with token
// @ID          change-password-with-token
// @Tags  	    Authentication
// @Accept      json
// @Produce     json
// @Param       payload body changePasswordWithTokenReq true "Change password with token"
// @Success     204 ""
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /api/customer/v1/change-password-with-token [post]
func (s *authRoute) changePassword(ctx iris.Context) {
	request := new(changePasswordWithTokenReq)
	if err := ctx.ReadJSON(request); err != nil {
		handleBindingError(ctx, err, s.logger, request, nil)
		return
	}
	i := &model.CustomerChangePasswordWithTokenInput{
		Token:           request.Token,
		Otp:             request.Otp,
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
	}
	i, err := s.uc.ValidateChangePasswordWithToken(ctx, i)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	err = s.uc.ChangePasswordWithToken(ctx, i)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	ctx.StatusCode(iris.StatusNoContent)
}
