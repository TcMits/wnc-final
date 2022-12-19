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
	sk, _ := uc.GetSecret()
	handler.Post("/token", route.renewToken)
	handler.Post("/login", route.login)
	handler.Delete("/login", middleware.Authenticator(sk, uc.GetUser), route.logout)
	handler.Options("/login", func(_ iris.Context) {})
	handler.Options("/token", func(_ iris.Context) {})
}

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

func (r *authRoute) logout(ctx iris.Context) {
	userAny, _ := ctx.User().GetRaw()
	user, _ := userAny.(*model.Customer)
	err := r.uc.Logout(ctx, user)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	ctx.StatusCode(iris.StatusNoContent)
}

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
