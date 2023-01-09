package employee

import (
	"fmt"

	"github.com/TcMits/wnc-final/ent"
	"github.com/TcMits/wnc-final/ent/transaction"
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/admin/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/TcMits/wnc-final/pkg/tool/generic"
	"github.com/kataras/iris/v12"
)

type transactionRoute struct {
	uc     usecase.IEmployeeTransactionUseCase
	logger logger.Interface
}

func RegisterTransactionController(handler iris.Party, l logger.Interface, uc usecase.IEmployeeTransactionUseCase) {
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
// @Param       only_debt query bool false "True if only debt transaction otherwise ignored"
// @Param       sender_id query string false "ID of bank account"
// @Param       receiver_id query string false "ID of bank account"
// @Success     200 {object} EntitiesResponseTemplate[transactionResp]
// @Failure     500 {object} errorResponse
// @Router      /transactions [get]
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
	if filterReq.OnlyDebt {
		w.HasDebt = generic.GetPointer(true)
	} else if filterReq.ReceiverID != nil {
		w.ReceiverID = filterReq.ReceiverID
	} else if filterReq.SenderID != nil {
		w.SenderID = filterReq.SenderID
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
// @Router      /transactions/{id} [get]
func (r *transactionRoute) detail(ctx iris.Context) {
	req := new(detailRequest)
	if err := ctx.ReadParams(req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	entity, err := r.uc.GetFirst(ctx, nil, &model.TransactionWhereInput{ID: req.id})
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	if entity != nil {
		ctx.JSON(getResponse(entity))
	} else {
		ctx.StatusCode(iris.StatusNoContent)
	}
}
