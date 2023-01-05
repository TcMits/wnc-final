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
	handler.Use(middleware.Authenticator(uc.GetSecret(), uc.GetUser))
	handler.Put("/debts/fulfill-with-token/{id:uuid}", route.fulfillWithToken)
	handler.Put("/debts/cancel/{id:uuid}", route.cancel)
	handler.Put("/debts/fulfill/{id:uuid}", route.fulfill)
	handler.Get("/debts/{id:uuid}", route.detail)
	handler.Get("/debts", route.listing)
	handler.Post("/debts", route.create)
	handler.Options("/debts/cancel", func(_ iris.Context) {})
	handler.Options("/debts/fulfill-with-token", func(_ iris.Context) {})
	handler.Options("/debts/fulfill", func(_ iris.Context) {})
	handler.Options("/debts", func(_ iris.Context) {})
	handler.Head("/debts/cancel", func(_ iris.Context) {})
	handler.Head("/debts/fulfill-with-token", func(_ iris.Context) {})
	handler.Head("/debts/fulfill", func(_ iris.Context) {})
	handler.Head("/debts", func(_ iris.Context) {})
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
	ctx.JSON(getResponse(entity))
}

// @Summary     Cancel a debt
// @Description Cancel a debt
// @ID          debt-cancel
// @Tags  	    Debt
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       id path string true "ID of debt"
// @Param       payload body debtCancelReq true "Cancel a debt"
// @Success     200 {object} debtResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /debts/cancel/{id} [put]
func (s *debtRoute) cancel(ctx iris.Context) {
	req := new(detailRequest)
	if err := ctx.ReadParams(req); err != nil {
		handleBindingError(ctx, err, s.logger, req, nil)
		return
	}
	updateInReq := new(debtCancelReq)
	if err := ctx.ReadBody(updateInReq); err != nil {
		handleBindingError(ctx, err, s.logger, updateInReq, nil)
		return
	}
	entity, err := s.uc.GetFirstMine(ctx, nil, &model.DebtWhereInput{ID: req.id})
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	if entity == nil {
		ctx.StatusCode(iris.StatusNoContent)
	}
	i := &model.DebtUpdateInput{
		Description: &updateInReq.Description,
	}
	i, err = s.uc.ValidateCancel(ctx, entity, i)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	entity, err = s.uc.Cancel(ctx, entity, i)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	ctx.JSON(getResponse(entity))
}

// @Summary     Fulfill a debt
// @Description Fulfill a debt
// @ID          debt-fulfill
// @Tags  	    Debt
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       id path string true "ID of debt"
// @Success     200 {object} debtFulfillResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /debts/fulfill/{id} [put]
func (s *debtRoute) fulfill(ctx iris.Context) {
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
	if entity == nil {
		ctx.StatusCode(iris.StatusNoContent)
	}
	err = s.uc.ValidateFulfill(ctx, entity)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	res, err := s.uc.Fulfill(ctx, entity)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	ctx.JSON(getResponse(res))
}

// @Summary     Fulfill a debt with token
// @Description Fulfill a debt with token
// @ID          fulfill-debt-with-token
// @Tags  	    Debt
// @Accept      json
// @Produce     json
// @Param       id path string true "ID of debt"
// @Param       payload body debtFulfillReq true "Fulfill a debt with token"
// @Success     200 {object} debtResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /debts/fulfill-with-token/{id} [put]
func (s *debtRoute) fulfillWithToken(ctx iris.Context) {
	req := new(detailRequest)
	if err := ctx.ReadParams(req); err != nil {
		handleBindingError(ctx, err, s.logger, req, nil)
		return
	}
	request := new(debtFulfillReq)
	if err := ctx.ReadJSON(request); err != nil {
		handleBindingError(ctx, err, s.logger, request, nil)
		return
	}
	entity, err := s.uc.GetFirstMine(ctx, nil, &model.DebtWhereInput{ID: req.id})
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	if entity == nil {
		ctx.StatusCode(iris.StatusNoContent)
	}
	i := &model.DebtFulfillWithTokenInput{
		Otp:   request.Otp,
		Token: request.Token,
	}
	i, err = s.uc.ValidateFulfillWithToken(ctx, entity, i)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	res, err := s.uc.FulfillWithToken(ctx, entity, i)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	ctx.JSON(getResponse(res))
}
