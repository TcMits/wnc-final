package customer

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/customer/middleware"
	"github.com/TcMits/wnc-final/internal/sse"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type streamRoute struct {
	uc     usecase.ICustomerStreamUseCase
	logger logger.Interface
	broker *sse.Broker
}

func RegisterStreamController(handler iris.Party, l logger.Interface, broker *sse.Broker, uc usecase.ICustomerStreamUseCase) {
	r := &streamRoute{
		uc:     uc,
		logger: l,
		broker: broker,
	}
	handler.Get("/stream", middleware.Authenticator(r.uc.GetSecret(), r.uc.GetUser), r.serve)
}

func (s *streamRoute) serve(ctx iris.Context) {
	s.broker.ServeHTTP(ctx)
}
