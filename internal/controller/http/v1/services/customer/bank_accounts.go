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
	route := &bankAccountRoute{
		uc:     uc,
		logger: l,
	}

	handler.Get("/bank-accounts/guest/{id:uuid}", middleware.Authenticator(uc.GetSecret(), uc.GetUser), route.guestDetail)
	handler.Get("/bank-accounts/guest", middleware.Authenticator(uc.GetSecret(), uc.GetUser), route.guestListing)
	handler.Put("/bank-accounts/{id:uuid}", middleware.Authenticator(uc.GetSecret(), uc.GetUser), route.update)
	handler.Get("/bank-accounts/{id:uuid}", middleware.Authenticator(uc.GetSecret(), uc.GetUser), route.detail)
	handler.Get("/bank-accounts", middleware.Authenticator(uc.GetSecret(), uc.GetUser), route.listing)
	handler.Options("/bank-accounts/guest", func(_ iris.Context) {})
	handler.Options("/bank-accounts", func(_ iris.Context) {})
}

// @Summary     Show bank accounts
// @Description Show bank accounts
// @ID          bankaccount-listing
// @Tags  	    Bank account
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Success     200 {object} bankAccountResp
// @Failure     500 {object} errorResponse
// @Router      /bank-accounts [get]
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
	ctx.JSON(getResponses(entities))
}

// @Summary     Show guest bank accounts
// @Description Show guest bank accounts
// @ID          guestbankaccount-listing
// @Tags  	    Guest bank account
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Success     200 {object} guestBankAccountResp
// @Failure     500 {object} errorResponse
// @Router      /bank-accounts/guest [get]
func (r *bankAccountRoute) guestListing(ctx iris.Context) {
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
	ctx.JSON(getResponses(entities, getGuestBankAccountResp))
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
// @Router      /bank-accounts/{id} [put]
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
	l, o := 1, 0
	entities, err := r.uc.List(ctx, &l, &o, nil, &model.BankAccountWhereInput{
		ID: req.id,
	})
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	if len(entities) > 0 {
		entity := entities[0]
		in := &model.BankAccountUpdateInput{
			IsForPayment: &updateInReq.IsForPayment,
		}
		in, err = r.uc.Validate(ctx, entity, in)
		if err != nil {
			HandleError(ctx, err, r.logger)
			return
		}
		entity, err = r.uc.Update(ctx, entity, in)
		if err != nil {
			HandleError(ctx, err, r.logger)
			return
		}
		ctx.JSON(getResponse(entity))
	} else {
		ctx.StatusCode(iris.StatusNoContent)
	}
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
// @Router      /bank-accounts/guest/{id} [get]
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

// @Summary     Get a guest bank account
// @Description Get a guest bank account
// @ID          guestbankaccount-get
// @Tags  	    Guest bank account
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       id path string true "ID of bank account"
// @Success     200 {object} guestBankAccountResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /bank-accounts/{id} [get]
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
