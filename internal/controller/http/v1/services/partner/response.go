package partner

import (
	"strconv"
	"time"

	"github.com/TcMits/wnc-final/ent/transaction"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/jwt"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/shopspring/decimal"
)

type (
	EntitiesResponseTemplate[EntityResponse any] struct {
		Count    uint             `json:"count"`
		Next     string           `json:"next"`
		Previous string           `json:"previous"`
		Results  []EntityResponse `json:"results"`
	}
	pagingInput[ModelType any] struct {
		limit        int
		offset       int
		noPagingResp *EntitiesResponseTemplate[any]
		isNext       bool
		entities     []ModelType
	}
	emptyResponse struct{}

	// error
	errorResponse struct {
		Message string `json:"message"`
		Code    string `json:"code"`
		Detail  string `json:"detail"`
	}
	tokenPairResponse struct {
		AccessToken  *string `json:"access_token"`
		RefreshToken *string `json:"refresh_token"`
	}
	transactionResp struct {
		ID                        uuid.UUID                   `json:"id"`
		CreateTime                time.Time                   `json:"create_time"`
		UpdateTime                time.Time                   `json:"update_time"`
		SourceTransactionID       *uuid.UUID                  `json:"source_transaction_id"`
		Status                    transaction.Status          `json:"status"`
		ReceiverBankAccountNumber string                      `json:"receiver_bank_account_number"`
		ReceiverBankName          string                      `json:"receiver_bank_name"`
		ReceiverName              string                      `json:"receiver_name"`
		ReceiverID                *uuid.UUID                  `json:"receiver_id"`
		SenderBankAccountNumber   string                      `json:"sender_bank_account_number"`
		SenderBankName            string                      `json:"sender_bank_name"`
		SenderName                string                      `json:"sender_name"`
		SenderID                  *uuid.UUID                  `json:"sender_id"`
		Amount                    decimal.Decimal             `json:"amount"`
		TransactionType           transaction.TransactionType `json:"transaction_type"`
		Description               string                      `json:"description"`
	}
	bankAccountResp struct {
		ID            uuid.UUID `json:"id"`
		CreateTime    time.Time `json:"create_time"`
		UpdateTime    time.Time `json:"update_time"`
		CustomerID    uuid.UUID `json:"customer_id"`
		AccountNumber string    `json:"account_number"`
		IsForPayment  bool      `json:"is_for_payment"`
	}
	optionsResp struct {
		ActorTypes []string `json:"actor_types"`
	}
	// reference on docs
)

func getDefaultResponse(entity any) any {
	var result any
	switch entity.(type) {
	case *jwt.TokenPair:
		rs, _ := entity.(*jwt.TokenPair)
		result = &tokenPairResponse{
			AccessToken:  rs.AccessToken,
			RefreshToken: rs.RefreshToken,
		}
	case *model.Transaction:
		rs, _ := entity.(*model.Transaction)
		result = &transactionResp{
			ID:                        rs.ID,
			CreateTime:                rs.CreateTime,
			UpdateTime:                rs.UpdateTime,
			SourceTransactionID:       rs.SourceTransactionID,
			Status:                    rs.Status,
			ReceiverBankAccountNumber: rs.ReceiverBankAccountNumber,
			ReceiverBankName:          rs.ReceiverBankName,
			ReceiverName:              rs.ReceiverName,
			ReceiverID:                rs.ReceiverID,
			SenderBankAccountNumber:   rs.SenderBankAccountNumber,
			SenderName:                rs.SenderName,
			SenderID:                  rs.SenderID,
			Amount:                    rs.Amount,
			TransactionType:           rs.TransactionType,
			Description:               rs.Description,
		}
	case *model.BankAccount:
		rs, _ := entity.(*model.BankAccount)
		result = &bankAccountResp{
			ID:            rs.ID,
			CreateTime:    rs.CreateTime,
			UpdateTime:    rs.UpdateTime,
			CustomerID:    rs.CustomerID,
			AccountNumber: rs.AccountNumber,
			IsForPayment:  rs.IsForPayment,
		}
	default:
		result = entity
	}
	return result
}

func getResponse(entity any, args ...func(any) any) any {
	var result any = entity
	if len(args) == 0 {
		args = append(args, getDefaultResponse)
	}
	for _, t := range args {
		result = t(result)
	}
	return result
}

func getResponses[ModelType any](entities []ModelType, args ...func(any) any) *EntitiesResponseTemplate[any] {
	fr := make([]any, 0, len(entities))

	for _, entity := range entities {
		fr = append(fr, getResponse(entity, args...))
	}
	return &EntitiesResponseTemplate[any]{Results: fr}
}

func getPagingResponse[ModelType any](ctx iris.Context, i pagingInput[ModelType], args ...func(any) any) *EntitiesResponseTemplate[any] {
	var pageResp *EntitiesResponseTemplate[any]
	if i.noPagingResp != nil {
		pageResp = i.noPagingResp
	} else {
		pageResp = getResponses(i.entities, args...)
	}
	if i.isNext {
		originUrl := ctx.Request().URL
		url := *originUrl
		q := url.Query()
		q.Set("limit", strconv.Itoa(i.limit))
		offset := i.offset + i.limit
		q.Set("offset", strconv.Itoa(offset))
		url.RawQuery = q.Encode()
		pageResp.Next = url.String()
	}
	if i.offset >= i.limit {
		originUrl := ctx.Request().URL
		url := *originUrl
		q := url.Query()
		q.Set("limit", strconv.Itoa(i.limit))
		offset := i.offset - i.limit
		q.Set("offset", strconv.Itoa(offset))
		url.RawQuery = q.Encode()
		pageResp.Previous = url.String()
	}
	pageResp.Count = uint(len(pageResp.Results))
	return pageResp
}
