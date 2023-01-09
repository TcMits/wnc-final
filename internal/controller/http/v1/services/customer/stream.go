package customer

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/customer/middleware"
	"github.com/TcMits/wnc-final/internal/sse"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

// @Summary     Receive events
// @Description Receive events
// @ID          event
// @Tags  	    Event
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Success     200 {object} eventResp
// @Failure     505 {object} errorResponse
// @Router      /stream [get]
func RegisterStreamController(handler iris.Party, l logger.Interface, broker *sse.Broker, uc usecase.ICustomerStreamUseCase) {
	h := handler.Party("/")
	h.Get("/stream", middleware.Authenticator(uc.GetSecret(), uc.GetUser), func(ctx iris.Context) {
		broker.ServeHTTP(ctx)
	})
}
