package partner

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/partner/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type transactionRoute struct {
	uc     usecase.IPartnerTransactionUseCase
	logger logger.Interface
}

func RegisterTransactionController(handler iris.Party, l logger.Interface, uc usecase.IPartnerTransactionUseCase) {
	h := handler.Party("/")
	route := &transactionRoute{
		uc:     uc,
		logger: l,
	}
	h.Use(middleware.Authenticator(uc.GetSecret(), uc.GetUser))
	h.Post("/transactions/validate", route.validate)
	h.Post("/transactions", route.create)
	h.Options("/transactions/validate", func(_ iris.Context) {})
	h.Options("/transactions", func(_ iris.Context) {})
	h.Head("/transactions/validate", func(_ iris.Context) {})
	h.Head("/transactions", func(_ iris.Context) {})
}

// @Summary     Create a transaction
// @Description Create a transaction
// @ID          transaction-create
// @Tags  	    Transaction
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       payload body transactionCreateReq true "Create a transaction"
// @Success     201 {object} transactionResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /api/partner/v1/transactions [post]
func (r *transactionRoute) create(ctx iris.Context) {
	createInReq := new(transactionCreateReq)
	if err := ctx.ReadBody(createInReq); err != nil {
		handleBindingError(ctx, err, r.logger, createInReq, nil)
		return
	}
	in := &model.PartnerTransactionCreateInput{
		TransactionCreateInput: &model.TransactionCreateInput{
			Amount:                    createInReq.Amount,
			Description:               &createInReq.Description,
			SenderName:                createInReq.SenderName,
			SenderBankAccountNumber:   createInReq.SenderBankAccountNumber,
			ReceiverBankAccountNumber: createInReq.ReceiverBankAccountNumber,
		},
		Token:     createInReq.Token,
		Signature: createInReq.Signature,
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
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(getResponse(entity))
}

// @Summary     Validate before create transaction
// @Description Validate before create transaction
// @ID          transaction-validate
// @Tags  	    Transaction
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       payload body transactionCreateReq true "Validate before create transaction"
// @Success     204 ""
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /api/partner/v1/transactions/validate [post]
func (r *transactionRoute) validate(ctx iris.Context) {
	createInReq := new(transactionCreateReq)
	if err := ctx.ReadBody(createInReq); err != nil {
		handleBindingError(ctx, err, r.logger, createInReq, nil)
		return
	}
	in := &model.PartnerTransactionCreateInput{
		TransactionCreateInput: &model.TransactionCreateInput{
			Amount:                    createInReq.Amount,
			Description:               &createInReq.Description,
			SenderName:                createInReq.SenderName,
			SenderBankAccountNumber:   createInReq.SenderBankAccountNumber,
			ReceiverBankAccountNumber: createInReq.ReceiverBankAccountNumber,
		},
		Token:     createInReq.Token,
		Signature: createInReq.Signature,
	}
	_, err := r.uc.ValidateCreate(ctx, in)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	ctx.StatusCode(iris.StatusNoContent)
}
