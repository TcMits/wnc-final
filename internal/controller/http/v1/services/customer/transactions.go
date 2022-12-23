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

func (r *transactionRoute) listing(ctx iris.Context) {
	req := newListRequest()
	if err := ctx.ReadQuery(req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	entities, err := r.uc.ListMyTxc(ctx, &req.Limit, &req.Offset, nil, nil)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	ctx.JSON(getResponses(entities))
}
func (r *transactionRoute) create(ctx iris.Context) {
	createInReq := new(transactionCreateRequest)
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
		SenderID:                  *createInReq.SenderID,
	}
	in, err := r.uc.Validate(ctx, in, createInReq.IsFeePaidByMe)
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

func (r *transactionRoute) detail(ctx iris.Context) {
	req := new(detailRequest)
	if err := ctx.ReadParams(req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	entity, err := r.uc.GetFirstMyTxc(ctx, nil, &model.TransactionWhereInput{ID: req.id})
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

func (r *transactionRoute) confirmSuccess(ctx iris.Context) {
	req := new(detailRequest)
	if err := ctx.ReadParams(req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	entity, err := r.uc.GetFirstMyTxc(ctx, nil, &model.TransactionWhereInput{ID: req.id})
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	if entity != nil {
		confirmReq := new(transactionConfirmRequest)
		if err := ctx.ReadBody(confirmReq); err != nil {
			handleBindingError(ctx, err, r.logger, confirmReq, nil)
			return
		}
		err = r.uc.ValidateConfirmInput(ctx, entity, &confirmReq.Token)
		if err != nil {
			HandleError(ctx, err, r.logger)
			return
		}
		entity, err = r.uc.ConfirmAsSuccess(ctx, entity, &confirmReq.Token)
		if err != nil {
			HandleError(ctx, err, r.logger)
			return
		}
		ctx.JSON(getResponse(entity))
	} else {
		ctx.StatusCode(iris.StatusNoContent)
	}
}
