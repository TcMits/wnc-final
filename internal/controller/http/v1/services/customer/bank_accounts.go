package customer

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/customer/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type bankAccountRoute struct {
	uc     usecase.ICustomerBankAccountUseCase
	logger logger.Interface
}

func RegisterBankAccountController(handler iris.Party, l logger.Interface, uc usecase.ICustomerBankAccountUseCase) {
	h := handler.Party("/")
	route := &bankAccountRoute{
		uc:     uc,
		logger: l,
	}
	h.Use(middleware.Authenticator(uc.GetSecret(), uc.GetUser))
	h.Get("/bank-accounts/guest/{id:uuid}", route.guestDetail)
	h.Get("/bank-accounts/guest", route.guestListing)
	h.Get("/bank-accounts/tp-bank", route.getTPBankAcc)
	h.Put("/bank-accounts/{id:uuid}", route.update)
	h.Get("/bank-accounts/{id:uuid}", route.detail)
	h.Get("/bank-accounts", route.listing)
	h.Options("/bank-accounts/guest", func(_ iris.Context) {})
	h.Options("/bank-accounts", func(_ iris.Context) {})
	h.Head("/bank-accounts/guest", func(_ iris.Context) {})
	h.Head("/bank-accounts", func(_ iris.Context) {})
}

// @Summary     Show bank accounts
// @Description Show bank accounts
// @ID          bankaccount-listing
// @Tags  	    Bank account
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Success     200 {object} EntitiesResponseTemplate[bankAccountResp]
// @Failure     500 {object} errorResponse
// @Router      /me/bank-accounts [get]
func (r *bankAccountRoute) listing(ctx iris.Context) {
	req := newListRequest()
	if err := ctx.ReadQuery(req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	entities, err := r.uc.ListMine(ctx, &req.Limit, &req.Offset, nil, nil)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	isNext, err := r.uc.IsNext(ctx, req.Limit, req.Offset, nil, nil)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	paging := getPagingResponse(ctx, pagingInput[*model.BankAccount]{
		limit:    req.Limit,
		offset:   req.Offset,
		entities: entities,
		isNext:   isNext,
	})
	ctx.JSON(paging)
}

// @Summary     Show guest bank accounts
// @Description Show guest bank accounts
// @ID          guestbankaccount-listing
// @Tags  	    Bank account
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       account_number query string false "Bank account number"
// @Success     200 {object} EntitiesResponseTemplate[guestBankAccountResp]
// @Failure     500 {object} errorResponse
// @Router      /me/bank-accounts/guest [get]
func (r *bankAccountRoute) guestListing(ctx iris.Context) {
	req := newListRequest()
	if err := ctx.ReadQuery(req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	filterReq := new(bankAccountFilterReq)
	if err := ctx.ReadQuery(filterReq); err != nil {
		handleBindingError(ctx, err, r.logger, filterReq, nil)
		return
	}
	entities, err := r.uc.List(ctx, &req.Limit, &req.Offset, nil, &model.BankAccountWhereInput{
		AccountNumber: filterReq.AccountNumber,
	})
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	isNext, err := r.uc.IsNext(ctx, req.Limit, req.Offset, nil, &model.BankAccountWhereInput{
		AccountNumber: filterReq.AccountNumber,
	})
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	paging := getPagingResponse(ctx, pagingInput[*model.BankAccount]{
		limit:    req.Limit,
		offset:   req.Offset,
		entities: entities,
		isNext:   isNext,
	}, getGuestBankAccountResp)
	ctx.JSON(paging)
}

// @Summary     Update a bank account
// @Description Update a bank account
// @ID          bankaccount-update
// @Tags  	    Bank account
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       payload body bankAccountUpdateReq true "Update a bank account"
// @Param       id path string true "ID of bank account"
// @Success     200 {object} bankAccountResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /me/bank-accounts/{id} [put]
func (r *bankAccountRoute) update(ctx iris.Context) {
	req := new(detailRequest)
	if err := ctx.ReadParams(req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	updateInReq := new(bankAccountUpdateReq)
	if err := ctx.ReadBody(updateInReq); err != nil {
		handleBindingError(ctx, err, r.logger, updateInReq, nil)
		return
	}
	e, err := r.uc.GetFirstMine(ctx, nil, &model.BankAccountWhereInput{
		ID: req.id,
	})
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	if e == nil {
		ctx.StatusCode(iris.StatusNoContent)
	}
	i, err := r.uc.ValidateUpdate(ctx, e, &model.BankAccountUpdateInput{
		IsForPayment: &updateInReq.IsForPayment,
	})
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	e, err = r.uc.Update(ctx, e, i)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	ctx.JSON(getResponse(e))
}

// @Summary     Get a bank account
// @Description Get a bank account
// @ID          bankaccount-get
// @Tags  	    Bank account
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       id path string true "ID of bank account"
// @Success     200 {object} bankAccountResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /me/bank-accounts/{id} [get]
func (s *bankAccountRoute) detail(ctx iris.Context) {
	req := new(detailRequest)
	if err := ctx.ReadParams(req); err != nil {
		handleBindingError(ctx, err, s.logger, req, nil)
		return
	}
	entity, err := s.uc.GetFirstMine(ctx, nil, &model.BankAccountWhereInput{ID: req.id})
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

// @Summary     Get a tp bank bank account
// @Description Get a tp bank bank account
// @ID          tpbankbankaccount-get
// @Tags  	    Bank account
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       account_number query string false "Bank account number"
// @Success     200 {object} bankAccountResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /me/bank-accounts/tp-bank [get]
func (s *bankAccountRoute) getTPBankAcc(ctx iris.Context) {
	filterReq := new(bankAccountFilterReq)
	if err := ctx.ReadQuery(filterReq); err != nil {
		handleBindingError(ctx, err, s.logger, filterReq, nil)
		return
	}
	e, err := s.uc.Get(ctx, &model.BankAccountWhereInput{
		AccountNumber: filterReq.AccountNumber,
	})
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	if e != nil {
		ctx.JSON(getResponse(e))
	} else {
		ctx.StatusCode(iris.StatusNoContent)
	}
}

// @Summary     Get a guest bank account
// @Description Get a guest bank account
// @ID          guestbankaccount-get
// @Tags  	    Bank account
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       id path string true "ID of bank account"
// @Success     200 {object} guestBankAccountResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /me/bank-accounts/guest/{id} [get]
func (s *bankAccountRoute) guestDetail(ctx iris.Context) {
	req := new(detailRequest)
	if err := ctx.ReadParams(req); err != nil {
		handleBindingError(ctx, err, s.logger, req, nil)
		return
	}
	entity, err := s.uc.GetFirst(ctx, nil, &model.BankAccountWhereInput{ID: req.id})
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	if entity != nil {
		ctx.JSON(getResponse(entity, getGuestBankAccountResp))
	} else {
		ctx.StatusCode(iris.StatusNoContent)
	}
}
