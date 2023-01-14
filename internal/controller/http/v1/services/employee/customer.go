package employee

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/employee/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type customerRoute struct {
	uc     usecase.IEmployeeCustomerUseCase
	logger logger.Interface
}

func RegisterCustomerController(handler iris.Party, l logger.Interface, uc usecase.IEmployeeCustomerUseCase) {
	h := handler.Party("/")
	route := &customerRoute{
		uc:     uc,
		logger: l,
	}
	h.Use(middleware.Authenticator(uc.GetSecret(), uc.GetUser))
	h.Get("/customers/{id:uuid}", route.detail)
	h.Post("/customers", route.create)
	h.Get("/customers", route.listing)
	h.Options("/customers", func(_ iris.Context) {})
	h.Head("/customers", func(_ iris.Context) {})
}

// @Summary     Create a customer
// @Description Create a customer
// @ID          customer-create
// @Tags  	    Customer
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       payload body customerCreateReq true "Create a customer"
// @Success     201 {object} customerCreateResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /api/employee/v1/customers [post]
func (s *customerRoute) create(ctx iris.Context) {
	createInReq := new(customerCreateReq)
	if err := ctx.ReadBody(createInReq); err != nil {
		handleBindingError(ctx, err, s.logger, createInReq, nil)
		return
	}
	in := &model.CustomerCreateInput{
		Username:    createInReq.Username,
		Email:       createInReq.Email,
		PhoneNumber: createInReq.PhoneNumber,
	}
	in, err := s.uc.ValidateCreate(ctx, in)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	entity, err := s.uc.Create(ctx, in)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(getResponse(entity))
}

// @Summary     Show customers
// @Description Show customers
// @ID          customer-listing
// @Tags  	    Customer
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Success     200 {object} EntitiesResponseTemplate[customerResponse]
// @Failure     500 {object} errorResponse
// @Router      /api/employee/v1/customers [get]
func (r *customerRoute) listing(ctx iris.Context) {
	req := newListRequest()
	if err := ctx.ReadQuery(req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	entities, err := r.uc.List(ctx, &req.Limit, &req.Offset, nil, nil)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	isNext, err := r.uc.IsNext(ctx, req.Limit, req.Offset, nil, nil)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	paging := getPagingResponse(ctx, pagingInput[*model.Customer]{
		limit:    req.Limit,
		offset:   req.Offset,
		entities: entities,
		isNext:   isNext,
	})
	ctx.JSON(paging)
}

// @Summary     Get a customer
// @Description Get a customer
// @ID          customer-get
// @Tags  	    Customer
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       id path string true "ID of customer"
// @Success     200 {object} customerResponse
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /api/employee/v1/customers/{id} [get]
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
