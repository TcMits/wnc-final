package customer

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/customer/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type transactionRoute struct {
	uc     usecase.ICustomerTransactionUseCase
	logger logger.Interface
}

func RegisterTransactionController(handler iris.Party, l logger.Interface, uc usecase.ICustomerTransactionUseCase) {
	route := &transactionRoute{
		uc:     uc,
		logger: l,
	}
	handler.Put("/transactions/confirm-success/{id:uuid}", middleware.Authenticator(uc.GetSecret(), uc.GetUser), route.confirmSuccess)
	handler.Get("/transactions/{id:uuid}", middleware.Authenticator(uc.GetSecret(), uc.GetUser), route.detail)
	handler.Get("/transactions", middleware.Authenticator(uc.GetSecret(), uc.GetUser), route.listing)
	handler.Post("/transactions", middleware.Authenticator(uc.GetSecret(), uc.GetUser), route.create)
	handler.Options("/transactions", func(_ iris.Context) {})
}

// @Summary     Show transactions
// @Description Show transactions
// @ID          transaction-listing
// @Tags  	    Transaction
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Success     200 {object} transactionResp
// @Failure     500 {object} errorResponse
// @Router      /transactions [get]
func (r *transactionRoute) listing(ctx iris.Context) {
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
// @Router      /transactions [post]
func (r *transactionRoute) create(ctx iris.Context) {
	createInReq := new(transactionCreateReq)
	if err := ctx.ReadBody(createInReq); err != nil {
		handleBindingError(ctx, err, r.logger, createInReq, nil)
		return
	}
	in := &model.TransactionCreateInput{
		ReceiverBankAccountNumber: createInReq.ReceiverBankAccountNumber,
		ReceiverBankName:          createInReq.ReceiverBankName,
		ReceiverName:              createInReq.ReceiverName,
		ReceiverID:                createInReq.ReceiverID,
		Amount:                    createInReq.Amount,
		Description:               &createInReq.Description,
	}
	in, err := r.uc.Validate(ctx, in, createInReq.IsFeePaidByMe)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	entity, err := r.uc.Create(ctx, in, createInReq.IsFeePaidByMe)
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
// @Router      /transactions/{id} [get]
func (r *transactionRoute) detail(ctx iris.Context) {
	req := new(detailRequest)
	if err := ctx.ReadParams(req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	entity, err := r.uc.GetFirstMine(ctx, nil, &model.TransactionWhereInput{ID: req.id})
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
// @Router      /transactions/{id} [put]
func (r *transactionRoute) confirmSuccess(ctx iris.Context) {
	req := new(detailRequest)
	if err := ctx.ReadParams(req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	entity, err := r.uc.GetFirstMine(ctx, nil, &model.TransactionWhereInput{ID: req.id})
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
		err = r.uc.ValidateConfirmInput(ctx, entity, &confirmReq.Token, &confirmReq.Token)
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
	}
}
