package partner

import (
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

// @Summary     Show options
// @Description Show all options
// @ID          option
// @Tags  	    Option
// @Accept      json
// @Produce     json
// @Success     200 {object} optionsResp
// @Failure     500 {object} errorResponse
// @Router      /api/partner/v1/options [get]
func RegisterOptionController(handler iris.Party, l logger.Interface, uc usecase.IPartnerOptionUseCase) {
	h := handler.Party("/")
	h.Get("/options", func(ctx iris.Context) {
		resp := new(optionsResp)
		resp.ActorTypes = uc.GetActorType(ctx)
		ctx.JSON(resp)
	})
	h.Head("/options", func(_ iris.Context) {})
	h.Options("/options", func(_ iris.Context) {})
}
