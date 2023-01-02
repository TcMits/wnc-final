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
	route := &authRoute{
		uc:     uc,
		logger: l,
	}
	handler.Post("/token", route.renewToken)
	handler.Post("/login", route.login)
	handler.Delete("/login", middleware.Authenticator(uc.GetSecret(), uc.GetUser), route.logout)
	handler.Options("/login", func(_ iris.Context) {})
	handler.Options("/token", func(_ iris.Context) {})
	handler.Head("/login", func(_ iris.Context) {})
}

// @Summary     Login
// @Description Login
// @ID          login
// @Tags  	    Login
// @Accept      json
// @Produce     json
// @Param       payload body loginRequest true "Login"
// @Success     200 {object} tokenPairResponse
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /login [post]
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
// @Tags  	    Logout
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Success     204  ""
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /login [Delete]
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
// @Tags  	    Renew token
// @Accept      json
// @Produce     json
// @Param       payload body renewTokenRequest true "Renew token"
// @Success     200 {object} tokenPairResponse
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /token [post]
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
