package customer

import (
	"fmt"

	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/ent/transaction"
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/customer/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/kataras/iris/v12"
)

type transactionRoute struct {
	uc     usecase.ICustomerTransactionUseCase
	logger logger.Interface
}

func RegisterTransactionController(handler iris.Party, l logger.Interface, uc usecase.ICustomerTransactionUseCase) {
	h := handler.Party("/")
	route := &transactionRoute{
		uc:     uc,
		logger: l,
	}
	h.Use(middleware.Authenticator(uc.GetSecret(), uc.GetUser))
	h.Put("/transactions/tp-bank/confirm-success/{id:uuid}", route.tpBankConfirm)
	h.Post("/transactions/tp-bank", route.tpBankCreate)
	h.Put("/transactions/confirm-success/{id:uuid}", route.confirmSuccess)
	h.Get("/transactions/{id:uuid}", route.detail)
	h.Get("/transactions", route.listing)
	h.Post("/transactions", route.create)
	h.Options("/transactions/confirm-success", func(_ iris.Context) {})
	h.Options("/transactions", func(_ iris.Context) {})
	h.Head("/transactions/confirm-success", func(_ iris.Context) {})
	h.Head("/transactions", func(_ iris.Context) {})
}

// @Summary     Show transactions
// @Description Show transactions
// @ID          transaction-listing
// @Tags  	    Transaction
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       update_time query bool false "True if sort ascent by update_time otherwise ignored"
// @Param       -update_time query bool false "True if sort descent by update_time otherwise ignored"
// @Param       only_debt query bool false "True if only debt transaction otherwise ignored"
// @Param       only_receive query bool false "True if only receive transaction otherwise ignored"
// @Param       only_send query bool false "True if only send transaction otherwise ignored"
// @Success     200 {object} EntitiesResponseTemplate[transactionResp]
// @Failure     500 {object} errorResponse
// @Router      /api/customer/v1/me/transactions [get]
func (r *transactionRoute) listing(ctx iris.Context) {
	req := newListRequest()
	if err := ctx.ReadQuery(req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	orderReq := new(transactionOrderReq)
	if err := ctx.ReadQuery(orderReq); err != nil {
		handleBindingError(ctx, err, r.logger, orderReq, nil)
		return
	}
	or := new(model.TransactionOrderInput)
	if orderReq.UpdateTimeAsc {
		if o, err := ent.ParseOrderField(fmt.Sprintf("%v%v", ent.OrderDirectionAscPrefix, transaction.FieldUpdateTime)); err != nil {
			HandleError(ctx, err, r.logger)
			return
		} else {
			*or = append(*or, o)
		}
	} else if orderReq.UpdateTimeDesc {
		if o, err := ent.ParseOrderField(fmt.Sprintf("%v%v", ent.OrderDirectionDescPrefix, transaction.FieldUpdateTime)); err != nil {
			HandleError(ctx, err, r.logger)
			return
		} else {
			*or = append(*or, o)
		}
	}
	filterReq := new(transactionFilterReq)
	if err := ctx.ReadQuery(filterReq); err != nil {
		handleBindingError(ctx, err, r.logger, filterReq, nil)
		return
	}
	user := middleware.GetUserFromCtxAsCustomer(ctx)
	w := new(model.TransactionWhereInput)
	if filterReq.OnlyDebt {
		w.HasDebt = generic.GetPointer(true)
	} else if filterReq.OnlyReceive {
		w.HasReceiverWith = []*model.BankAccountWhereInput{{CustomerID: generic.GetPointer(user.ID)}}
	} else if filterReq.OnlySend {
		w.HasSenderWith = []*model.BankAccountWhereInput{{CustomerID: generic.GetPointer(user.ID)}}
	}
	entities, err := r.uc.ListMine(ctx, &req.Limit, &req.Offset, or, w)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	isNext, err := r.uc.IsNext(ctx, req.Limit, req.Offset, or, w)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	paging := getPagingResponse(ctx, pagingInput[*model.Transaction]{
		limit:    req.Limit,
		offset:   req.Offset,
		entities: entities,
		isNext:   isNext,
	})
	ctx.JSON(paging)
}

// @Summary     Create a transaction
// @Description Create a transaction
// @ID          transaction-create
// @Tags  	    Transaction
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       payload body transactionCreateReq true "Create a transaction"
// @Success     201 {object} transactionCreateResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /api/customer/v1/me/transactions [post]
func (r *transactionRoute) create(ctx iris.Context) {
	createInReq := new(transactionCreateReq)
	if err := ctx.ReadBody(createInReq); err != nil {
		handleBindingError(ctx, err, r.logger, createInReq, nil)
		return
	}
	in := &model.TransactionCreateUseCaseInput{
		TransactionCreateInput: &model.TransactionCreateInput{
			ReceiverID:      createInReq.ReceiverID,
			Amount:          createInReq.Amount,
			Description:     &createInReq.Description,
			TransactionType: transaction.TransactionTypeInternal,
		},
		IsFeePaidByMe: createInReq.IsFeePaidByMe,
	}
	in, err := r.uc.ValidateCreate(ctx, in)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	entity, err := r.uc.Create(ctx, in)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	ctx.JSON(getResponse(entity))
}

// @Summary     Create a tp bank transaction
// @Description Create a tp bank transaction
// @ID          tpbanktransaction-create
// @Tags  	    Transaction
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       payload body tpBankTransactionCreateReq true "Create a tp bank transaction"
// @Success     201 {object} transactionCreateResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /api/customer/v1/me/transactions/tp-bank [post]
func (r *transactionRoute) tpBankCreate(ctx iris.Context) {
	createInReq := new(tpBankTransactionCreateReq)
	if err := ctx.ReadBody(createInReq); err != nil {
		handleBindingError(ctx, err, r.logger, createInReq, nil)
		return
	}
	in := &model.TransactionCreateUseCaseInput{
		TransactionCreateInput: &model.TransactionCreateInput{
			ReceiverBankAccountNumber: createInReq.AccountNumber,
			Amount:                    createInReq.Amount,
			Description:               &createInReq.Description,
			TransactionType:           transaction.TransactionTypeExternal,
		},
		IsFeePaidByMe: createInReq.IsFeePaidByMe,
	}
	in, err := r.uc.ValidateCreate(ctx, in)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	entity, err := r.uc.Create(ctx, in)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	ctx.JSON(getResponse(entity))
}

// @Summary     Get a transaction
// @Description Get a transaction
// @ID          transaction-get
// @Tags  	    Transaction
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       id path string true "ID of transaction"
// @Success     200 {object} transactionResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /api/customer/v1/me/transactions/{id} [get]
func (r *transactionRoute) detail(ctx iris.Context) {
	req := new(detailRequest)
	if err := ReadID(ctx, req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	entity, err := r.uc.GetFirstMine(ctx, nil, &model.TransactionWhereInput{ID: &req.id})
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	if entity != nil {
		ctx.JSON(getResponse(entity))
	} else {
		ctx.StatusCode(iris.StatusNoContent)
		return
	}
}

// @Summary     Confirm success
// @Description Confirm success a transaction
// @ID          transaction-confirmsuccess
// @Tags  	    Transaction
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       payload body transactionConfirmReq true "Confirm a transaction"
// @Param       id path string true "ID of transaction"
// @Success     200 {object} transactionResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /api/customer/v1/me/transactions/confirm-success/{id} [put]
func (r *transactionRoute) confirmSuccess(ctx iris.Context) {
	req := new(detailRequest)
	if err := ReadID(ctx, req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	entity, err := r.uc.GetFirstMine(ctx, nil, &model.TransactionWhereInput{ID: &req.id})
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	if entity != nil {
		confirmReq := new(transactionConfirmReq)
		if err := ctx.ReadBody(confirmReq); err != nil {
			handleBindingError(ctx, err, r.logger, confirmReq, nil)
			return
		}
		err = r.uc.ValidateConfirmInput(ctx, entity, &model.TransactionConfirmUseCaseInput{
			Token: confirmReq.Token,
			Otp:   confirmReq.OTP,
		})
		if err != nil {
			HandleError(ctx, err, r.logger)
			return
		}
		entity, err = r.uc.ConfirmSuccess(ctx, entity, &confirmReq.Token)
		if err != nil {
			HandleError(ctx, err, r.logger)
			return
		}
		ctx.JSON(getResponse(entity))
	} else {
		ctx.StatusCode(iris.StatusNoContent)
		return
	}
}

// @Summary     Confirm a tp bank transaction
// @Description Confirm a tp bank transaction
// @ID          tpbanktransaction-confirmsuccess
// @Tags  	    Transaction
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       payload body transactionConfirmReq true "Confirm a transaction"
// @Param       id path string true "ID of transaction"
// @Success     200 {object} transactionResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /api/customer/v1/me/transactions/tp-bank/confirm-success/{id} [put]
func (r *transactionRoute) tpBankConfirm(ctx iris.Context) {
	req := new(detailRequest)
	if err := ReadID(ctx, req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	entity, err := r.uc.GetFirstMine(ctx, nil, &model.TransactionWhereInput{ID: &req.id})
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	if entity != nil {
		confirmReq := new(transactionConfirmReq)
		if err := ctx.ReadBody(confirmReq); err != nil {
			handleBindingError(ctx, err, r.logger, confirmReq, nil)
			return
		}
		err = r.uc.ValidateConfirmInput(ctx, entity, &model.TransactionConfirmUseCaseInput{
			Token: confirmReq.Token,
			Otp:   confirmReq.OTP,
		})
		if err != nil {
			HandleError(ctx, err, r.logger)
			return
		}
		entity, err = r.uc.ConfirmSuccess(ctx, entity, &confirmReq.Token)
		if err != nil {
			HandleError(ctx, err, r.logger)
			return
		}
		ctx.JSON(getResponse(entity))
	} else {
		ctx.StatusCode(iris.StatusNoContent)
		return
	}
}
