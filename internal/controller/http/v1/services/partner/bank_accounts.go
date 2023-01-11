package partner

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/partner/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type bankAccountRoute struct {
	uc     usecase.IPartnerBankAccountUseCase
	logger logger.Interface
}

func RegisterBankAccountController(handler iris.Party, l logger.Interface, uc usecase.IPartnerBankAccountUseCase) {
	h := handler.Party("/")
	route := &bankAccountRoute{
		uc:     uc,
		logger: l,
	}
	h.Use(middleware.Authenticator(uc.GetSecret(), uc.GetUser))
	h.Get("/bank-accounts", route.listing)
	h.Options("/bank-accounts", func(_ iris.Context) {})
	h.Head("/bank-accounts", func(_ iris.Context) {})
}

// @Summary     Show bank accounts
// @Description Show bank accounts
// @ID          bankaccount-listing
// @Tags  	    Bank account
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       account_number query string false "Account number of bank account"
// @Success     200 {object} EntitiesResponseTemplate[bankAccountResp]
// @Failure     500 {object} errorResponse
// @Router      /bank-accounts [get]
func (r *bankAccountRoute) listing(ctx iris.Context) {
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
	w := new(model.BankAccountWhereInput)
	if filterReq.AccountNumber != nil {
		w.AccountNumber = filterReq.AccountNumber
	}
	entities, err := r.uc.List(ctx, &req.Limit, &req.Offset, nil, w)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	isNext, err := r.uc.IsNext(ctx, req.Limit, req.Offset, nil, w)
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
