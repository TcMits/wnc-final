package customer

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/customer/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type customerRoute struct {
	uc     usecase.ICustomerUseCase
	logger logger.Interface
}

func RegisterCustomerController(handler iris.Party, l logger.Interface, uc usecase.ICustomerUseCase) {
	h := handler.Party("/")
	route := &customerRoute{
		uc:     uc,
		logger: l,
	}
	h.Use(middleware.Authenticator(uc.GetSecret(), uc.GetUser))
	h.Get("/customers/{id:uuid}", route.detail)
	h.Options("/customers", func(_ iris.Context) {})
	h.Head("/customers", func(_ iris.Context) {})
}

// @Summary     Get a customer
// @Description Get a customer
// @ID          customer-get
// @Tags  	    Customer
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       id path string true "ID of customer"
// @Success     200 {object} meResponse
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /api/customer/v1/customers/{id} [get]
func (s *customerRoute) detail(ctx iris.Context) {
	req := new(detailRequest)
	if err := ReadID(ctx, req); err != nil {
		handleBindingError(ctx, err, s.logger, req, nil)
		return
	}
	entity, err := s.uc.GetFirst(ctx, nil, &model.CustomerWhereInput{ID: &req.id})
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	if entity != nil {
		ctx.JSON(getResponse(entity))
	} else {
		ctx.StatusCode(iris.StatusNoContent)
		return
	}
}
