package employee

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/employee/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type bankAccountRoute struct {
	uc     usecase.IEmployeeBankAcountUseCase
	logger logger.Interface
}

func RegisterBankAccountController(handler iris.Party, l logger.Interface, uc usecase.IEmployeeBankAcountUseCase) {
	h := handler.Party("/")
	route := &bankAccountRoute{
		uc:     uc,
		logger: l,
	}
	h.Use(middleware.Authenticator(uc.GetSecret(), uc.GetUser))
	h.Put("/bank-accounts/{id:uuid}", route.update)
	h.Get("/bank-accounts/{id:uuid}", route.detail)
	h.Delete("/bank-accounts/{id:uuid}", route.delete)
	h.Get("/bank-accounts", route.listing)
	h.Options("/bank-accounts", func(_ iris.Context) {})
	h.Head("/bank-accounts", func(_ iris.Context) {})
}

// @Summary     Delete a bank account
// @Description Delete a bank account
// @ID          bankaccount-delete
// @Tags  	    Bank account
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       id path string true "ID of bank account"
// @Success     204 ""
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /bank-accounts/{id} [delete]
func (s *bankAccountRoute) delete(ctx iris.Context) {
	req := new(detailRequest)
	if err := ReadID(ctx, req); err != nil {
		handleBindingError(ctx, err, s.logger, req, nil)
		return
	}
	entity, err := s.uc.GetFirst(ctx, nil, &model.BankAccountWhereInput{ID: &req.id})
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	if entity != nil {
		err = s.uc.Delete(ctx, entity)
		if err != nil {
			HandleError(ctx, err, s.logger)
			return
		}
	}
	ctx.StatusCode(iris.StatusNoContent)
}

// @Summary     Show bank accounts
// @Description Show bank accounts
// @ID          bankaccount-listing
// @Tags  	    Bank account
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       account_number query string false "Account number of bank account"
// @Param       username query string false "Username of bank account"
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
	} else if filterReq.Username != nil {
		w.HasCustomerWith = []*model.CustomerWhereInput{
			{
				Username: filterReq.Username,
			},
		}
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

// @Summary     Deposit a bank account
// @Description Deposit a bank account
// @ID          bankaccount-deposit
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
	if err := ReadID(ctx, req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	updateInReq := new(bankAccountUpdateReq)
	if err := ctx.ReadBody(updateInReq); err != nil {
		handleBindingError(ctx, err, r.logger, updateInReq, nil)
		return
	}
	e, err := r.uc.GetFirst(ctx, nil, &model.BankAccountWhereInput{
		ID: &req.id,
	})
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	if e == nil {
		ctx.StatusCode(iris.StatusNoContent)
		return
	}
	i, err := r.uc.ValidateUpdate(ctx, e, &model.BankAccountUpdateInput{
		CashIn: updateInReq.CashIn,
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
// @Router      /bank-accounts/{id} [get]
func (s *bankAccountRoute) detail(ctx iris.Context) {
	req := new(detailRequest)
	if err := ReadID(ctx, req); err != nil {
		handleBindingError(ctx, err, s.logger, req, nil)
		return
	}
	entity, err := s.uc.GetFirst(ctx, nil, &model.BankAccountWhereInput{ID: &req.id})
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
