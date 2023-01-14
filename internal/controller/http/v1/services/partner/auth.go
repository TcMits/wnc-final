package partner

import (
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type authRoute struct {
	uc     usecase.IPartnerAuthUseCase
	logger logger.Interface
}

func RegisterAuthController(handler iris.Party, l logger.Interface, uc usecase.IPartnerAuthUseCase) {
	h := handler.Party("/")
	route := &authRoute{
		uc:     uc,
		logger: l,
	}
	h.Post("/auth", route.login)
	h.Options("/auth", func(_ iris.Context) {})
	h.Head("/auth", func(_ iris.Context) {})
}

// @Summary     Authenticate
// @Description Authenticate
// @ID          authenticate
// @Tags  	    Authentication
// @Accept      json
// @Produce     json
// @Param       payload body loginRequest true "Authenticate"
// @Success     200 {object} tokenPairResponse
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /api/partner/v1/auth [post]
func (r *authRoute) login(ctx iris.Context) {
	request := new(loginRequest)
	if err := ctx.ReadJSON(request); err != nil {
		handleBindingError(ctx, err, r.logger, request, nil)
		return
	}
	i := &model.PartnerLoginInput{
		ApiKey: request.ApiKey,
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
