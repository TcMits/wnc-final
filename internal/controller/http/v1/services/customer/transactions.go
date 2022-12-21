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
	sk, _ := uc.GetSecret()
	handler.Get("/transactions/{id:uuid}", middleware.Authenticator(sk, uc.GetUser), route.detail)
	handler.Get("/transactions", middleware.Authenticator(sk, uc.GetUser), route.listing)
	handler.Post("/transactions", middleware.Authenticator(sk, uc.GetUser), route.create)
	handler.Options("/transactions", func(_ iris.Context) {})
}

func (r *transactionRoute) listing(ctx iris.Context) {
	req := newListRequest()
	if err := ctx.ReadQuery(req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	uAny, err := ctx.User().GetRaw()
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	user, _ := uAny.(*model.Customer)
	entities, err := r.uc.List(ctx, &req.Limit, &req.Offset, nil, &model.TransactionWhereInput{
		Or: []*model.TransactionWhereInput{
			{
				HasReceiverWith: []*model.BankAccountWhereInput{
					{
						CustomerID: &user.ID,
					},
				},
				HasSenderWith: []*model.BankAccountWhereInput{
					{
						CustomerID: &user.ID,
					},
				},
			},
		},
	})
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
	in, err := r.uc.Validate(ctx, in)
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
	uAny, err := ctx.User().GetRaw()
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	user, _ := uAny.(*model.Customer)
	l, o := 1, 0
	entities, err := r.uc.List(ctx, &l, &o, nil, &model.TransactionWhereInput{
		ID: req.id,
		Or: []*model.TransactionWhereInput{
			{
				HasReceiverWith: []*model.BankAccountWhereInput{
					{
						CustomerID: &user.ID,
					},
				},
				HasSenderWith: []*model.BankAccountWhereInput{
					{
						CustomerID: &user.ID,
					},
				},
			},
		},
	})
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	if len(entities) > 0 {
		entity := entities[0]
		ctx.JSON(getResponse(entity))
	} else {
		ctx.StatusCode(iris.StatusNoContent)
	}
}
