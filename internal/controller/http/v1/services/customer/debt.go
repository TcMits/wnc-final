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
	h := handler.Party("/")
	route := &debtRoute{
		uc:     uc,
		logger: l,
	}
	h.Use(middleware.Authenticator(uc.GetSecret(), uc.GetUser))
	h.Put("/debts/fulfill-with-token/{id:uuid}", route.fulfillWithToken)
	h.Put("/debts/cancel/{id:uuid}", route.cancel)
	h.Put("/debts/fulfill/{id:uuid}", route.fulfill)
	h.Get("/debts/{id:uuid}", route.detail)
	h.Get("/debts", route.listing)
	h.Post("/debts", route.create)
	h.Options("/debts/cancel", func(_ iris.Context) {})
	h.Options("/debts/fulfill-with-token", func(_ iris.Context) {})
	h.Options("/debts/fulfill", func(_ iris.Context) {})
	h.Options("/debts", func(_ iris.Context) {})
	h.Head("/debts/cancel", func(_ iris.Context) {})
	h.Head("/debts/fulfill-with-token", func(_ iris.Context) {})
	h.Head("/debts/fulfill", func(_ iris.Context) {})
	h.Head("/debts", func(_ iris.Context) {})
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
// @Router      /api/customer/v1/me/debts/{id} [get]
func (s *debtRoute) detail(ctx iris.Context) {
	req := new(detailRequest)
	if err := ReadID(ctx, req); err != nil {
		handleBindingError(ctx, err, s.logger, req, nil)
		return
	}
	entity, err := s.uc.GetFirstMine(ctx, nil, &model.DebtWhereInput{ID: &req.id})
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

// @Summary     Show debt
// @Description Show debt
// @ID          debt-listing
// @Tags  	    Debt
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       owner_id query string false "ID of bank account"
// @Param       receiver_id query string false "ID of bank account"
// @Param       status query string false "Status of debt"
// @Success     200 {object} EntitiesResponseTemplate[debtResp]
// @Failure     500 {object} errorResponse
// @Router      /api/customer/v1/me/debts [get]
func (s *debtRoute) listing(ctx iris.Context) {
	req := newListRequest()
	if err := ctx.ReadQuery(req); err != nil {
		handleBindingError(ctx, err, s.logger, req, nil)
		return
	}
	filterReq := new(debtFilterReq)
	if err := ctx.ReadQuery(filterReq); err != nil {
		handleBindingError(ctx, err, s.logger, filterReq, nil)
		return
	}
	w := &model.DebtWhereInput{
		Status: filterReq.Status,
	}
	if filterReq.OwnerID != nil {
		w.OwnerID = filterReq.OwnerID
	} else if filterReq.ReceiverID != nil {
		w.ReceiverID = filterReq.ReceiverID
	}
	entities, err := s.uc.ListMine(ctx, &req.Limit, &req.Offset, nil, w)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	isNext, err := s.uc.IsNext(ctx, req.Limit, req.Offset, nil, w)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	paging := getPagingResponse(ctx, pagingInput[*model.Debt]{
		limit:    req.Limit,
		offset:   req.Offset,
		entities: entities,
		isNext:   isNext,
	})
	ctx.JSON(paging)
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
// @Router      /api/customer/v1/me/debts [post]
func (s *debtRoute) create(ctx iris.Context) {
	createInReq := new(debtCreateReq)
	if err := ctx.ReadBody(createInReq); err != nil {
		handleBindingError(ctx, err, s.logger, createInReq, nil)
		return
	}
	in := &model.DebtCreateInput{
		ReceiverID:  createInReq.ReceiverID,
		Amount:      createInReq.Amount,
		Description: createInReq.Description,
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
// @Router      /api/customer/v1/me/debts/cancel/{id} [put]
func (s *debtRoute) cancel(ctx iris.Context) {
	req := new(detailRequest)
	if err := ReadID(ctx, req); err != nil {
		handleBindingError(ctx, err, s.logger, req, nil)
		return
	}
	updateInReq := new(debtCancelReq)
	if err := ctx.ReadBody(updateInReq); err != nil {
		handleBindingError(ctx, err, s.logger, updateInReq, nil)
		return
	}
	entity, err := s.uc.GetFirstMine(ctx, nil, &model.DebtWhereInput{ID: &req.id})
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	if entity == nil {
		ctx.StatusCode(iris.StatusNoContent)
		return
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
// @Router      /api/customer/v1/me/debts/fulfill/{id} [put]
func (s *debtRoute) fulfill(ctx iris.Context) {
	req := new(detailRequest)
	if err := ReadID(ctx, req); err != nil {
		handleBindingError(ctx, err, s.logger, req, nil)
		return
	}
	entity, err := s.uc.GetFirstMine(ctx, nil, &model.DebtWhereInput{ID: &req.id})
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	if entity == nil {
		ctx.StatusCode(iris.StatusNoContent)
		return
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
// @Router      /api/customer/v1/me/debts/fulfill-with-token/{id} [put]
func (s *debtRoute) fulfillWithToken(ctx iris.Context) {
	req := new(detailRequest)
	if err := ReadID(ctx, req); err != nil {
		handleBindingError(ctx, err, s.logger, req, nil)
		return
	}
	request := new(debtFulfillReq)
	if err := ctx.ReadJSON(request); err != nil {
		handleBindingError(ctx, err, s.logger, request, nil)
		return
	}
	entity, err := s.uc.GetFirstMine(ctx, nil, &model.DebtWhereInput{ID: &req.id})
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	if entity == nil {
		ctx.StatusCode(iris.StatusNoContent)
		return
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
