package customer

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/customer/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type debtRoute struct {
	uc     usecase.ICustomerDebtUseCase
	logger logger.Interface
}

func RegisterDebtController(handler iris.Party, l logger.Interface, uc usecase.ICustomerDebtUseCase) {
	route := &debtRoute{
		uc:     uc,
		logger: l,
	}
	handler.Get("/debts/{id:uuid}", middleware.Authenticator(uc.GetSecret(), uc.GetUser), route.detail)
	handler.Get("/debts", middleware.Authenticator(uc.GetSecret(), uc.GetUser), route.listing)
	handler.Post("/debts", middleware.Authenticator(uc.GetSecret(), uc.GetUser), route.create)
	handler.Options("/debts", func(_ iris.Context) {})
}

// @Summary     Get a debt
// @Description Get a debt
// @ID          debt-get
// @Tags  	    Debt
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       id path string true "ID of debt"
// @Success     200 {object} debtResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /debts/{id} [get]
func (s *debtRoute) detail(ctx iris.Context) {
	req := new(detailRequest)
	if err := ctx.ReadParams(req); err != nil {
		handleBindingError(ctx, err, s.logger, req, nil)
		return
	}
	entity, err := s.uc.GetFirstMine(ctx, nil, &model.DebtWhereInput{ID: req.id})
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	if entity != nil {
		ctx.JSON(getResponse(entity))
	} else {
		ctx.StatusCode(iris.StatusNoContent)
	}
}

// @Summary     Show debt
// @Description Show debt
// @ID          debt-listing
// @Tags  	    Debt
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Success     200 {object} debtResp
// @Failure     500 {object} errorResponse
// @Router      /debts [get]
func (s *debtRoute) listing(ctx iris.Context) {
	req := newListRequest()
	if err := ctx.ReadQuery(req); err != nil {
		handleBindingError(ctx, err, s.logger, req, nil)
		return
	}
	entities, err := s.uc.ListMine(ctx, &req.Limit, &req.Offset, nil, nil)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	ctx.JSON(getResponses(entities))
}

// @Summary     Create a debt
// @Description Create a debt
// @ID          debt-create
// @Tags  	    Debt
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       payload body debtCreateReq true "Create a debt"
// @Success     201 {object} debtResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /debts [post]
func (s *debtRoute) create(ctx iris.Context) {
	createInReq := new(debtCreateReq)
	if err := ctx.ReadBody(createInReq); err != nil {
		handleBindingError(ctx, err, s.logger, createInReq, nil)
		return
	}
	in := &model.DebtCreateInput{
		ReceiverBankAccountNumber: createInReq.ReceiverBankAccountNumber,
		ReceiverName:              createInReq.ReceiverName,
		ReceiverID:                createInReq.ReceiverID,
		Amount:                    createInReq.Amount,
		Description:               createInReq.Description,
	}
	in, err := s.uc.Validate(ctx, in)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	entity, err := s.uc.Create(ctx, in)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	ctx.JSON(getResponse(entity))
}
