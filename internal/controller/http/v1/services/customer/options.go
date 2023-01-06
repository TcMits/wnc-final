package customer

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
// @Router      /options [get]
func RegisterOptionController(handler iris.Party, l logger.Interface, uc usecase.IOptionsUseCase) {
	handler.Get("/options", func(ctx iris.Context) {
		resp := new(optionsResp)
		resp.DebtStatus = uc.GetDebtStatus(ctx)
		ctx.JSON(resp)
	})
	handler.Head("/options", func(_ iris.Context) {})
	handler.Options("/options", func(_ iris.Context) {})
}