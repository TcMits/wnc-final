package admin

import (
	"fmt"

	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/ent/transaction"
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/admin/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type transactionRoute struct {
	uc     usecase.IAdminTransactionUseCase
	logger logger.Interface
}

func RegisterTransactionController(handler iris.Party, l logger.Interface, uc usecase.IAdminTransactionUseCase) {
	h := handler.Party("/")
	route := &transactionRoute{
		uc:     uc,
		logger: l,
	}
	h.Use(middleware.Authenticator(uc.GetSecret(), uc.GetUser))
	h.Get("/transactions/{id:uuid}", route.detail)
	h.Get("/transactions", route.listing)
	h.Options("/transactions", func(_ iris.Context) {})
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
// @Param       date_start query integer false "Date start"
// @Param       date_end query integer false "Date end"
// @Param       bank_name query string false "Bank name"
// @Success     200 {object} EntitiesResponseTemplate[transactionResp]
// @Failure     500 {object} errorResponse
// @Router      /api/admin/v1/transactions [get]
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
	w := new(model.TransactionWhereInput)
	if filterReq.DateStart != nil {
		w.CreateTimeGTE = &filterReq.DateStart.t
	}
	if filterReq.DateEnd != nil {
		w.CreateTimeLTE = &filterReq.DateEnd.t
	}
	if filterReq.BankName != nil {
		w.Or = append(w.Or, &model.TransactionWhereInput{SenderBankName: filterReq.BankName}, &model.TransactionWhereInput{ReceiverBankName: filterReq.BankName})
	}
	entities, err := r.uc.List(ctx, &req.Limit, &req.Offset, or, w)
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
// @Router      /api/admin/v1/transactions/{id} [get]
func (r *transactionRoute) detail(ctx iris.Context) {
	req := new(detailRequest)
	if err := ReadID(ctx, req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	entity, err := r.uc.GetFirst(ctx, nil, &model.TransactionWhereInput{ID: &req.id})
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
