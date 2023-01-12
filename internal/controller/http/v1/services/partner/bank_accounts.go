package partner

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/partner/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
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

// @Summary     Get bank account
// @Description Get bank account
// @ID          bankaccount-listing
// @Tags  	    Bank account
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       account_number query string false "Account number of bank account"
// @Success     200 {object} bankAccountResp
// @Failure     500 {object} errorResponse
// @Router      /bank-accounts [get]
func (r *bankAccountRoute) listing(ctx iris.Context) {
	filterReq := new(bankAccountFilterReq)
	if err := ctx.ReadQuery(filterReq); err != nil {
		handleBindingError(ctx, err, r.logger, filterReq, nil)
		return
	}
	w := &model.BankAccountWhereInput{
		IsForPayment: generic.GetPointer(true),
	}
	if filterReq.AccountNumber != nil {
		w.AccountNumber = filterReq.AccountNumber
	}
	e, err := r.uc.GetFirst(ctx, nil, w)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	if e == nil {
		ctx.StatusCode(iris.StatusNoContent)
		return
	}
	ctx.JSON(getResponse(e))
}
