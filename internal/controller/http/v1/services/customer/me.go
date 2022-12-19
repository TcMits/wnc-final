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
	sk, _ := uc.GetSecret()
	handler.Get("/me", middleware.Authenticator(sk, uc.GetUser), route.detail)
	handler.Options("/me", func(_ iris.Context) {})
}

// @Summary     Get profile
// @Description Get profile
// @ID          me
// @Tags  	    me
// @Accept      json
// @Produce     json
// @Success     200 {object} meResponse
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /me/ [get]
func (r *meRoute) detail(ctx iris.Context) {
	userAny, _ := ctx.User().GetRaw()
	user, _ := userAny.(*model.Customer)
	ctx.JSON(getResponse(user))
}
