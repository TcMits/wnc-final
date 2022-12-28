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

	handler.Put("/bank-accounts/{id:uuid}", middleware.Authenticator(uc.GetSecret(), uc.GetUser), route.update)
	handler.Get("/bank-accounts", middleware.Authenticator(uc.GetSecret(), uc.GetUser), route.listing)
	handler.Options("/bank-accounts", func(_ iris.Context) {})
}

// @Summary     Show bank accounts
// @Description Show bank accounts
// @ID          bankaccount-listing
// @Tags  	    bankaccounts
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
	entities, err := r.uc.List(ctx, &req.Limit, &req.Offset, nil, nil)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	ctx.JSON(getResponses(entities))
}

// @Summary     Update a bank account
// @Description Update a bank account
// @ID          bankaccount-update
// @Tags  	    bankaccounts
// @Accept      json
// @Produce     json
// @Param       payload body bankAccountUpdateReq true "Update a bank account"
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
