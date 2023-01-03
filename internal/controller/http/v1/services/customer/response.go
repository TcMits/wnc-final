package customer

import (
	"time"

	"github.com/TcMits/wnc-final/ent/debt"
	"github.com/TcMits/wnc-final/ent/transaction"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/jwt"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type (
	EntitiesResponseTemplate[EntityResponse any] struct {
		Results []EntityResponse `json:"results"`
	}
	emptyResponse struct{}

	// error
	errorResponse struct {
		Message string `json:"message"`
		Code    string `json:"code"`
		Detail  string `json:"detail"`
	}
	meResponse struct {
		ID          uuid.UUID `json:"id"`
		Username    string    `json:"username"`
		FirstName   string    `json:"first_name"`
		LastName    string    `json:"last_name"`
		PhoneNumber string    `json:"phone_number"`
		Email       string    `json:"email"`
		IsActive    bool      `json:"is_active"`
	}
	bankAccountResp struct {
		ID            uuid.UUID `json:"id"`
		CreateTime    time.Time `json:"create_time"`
		UpdateTime    time.Time `json:"update_time"`
		CustomerID    uuid.UUID `json:"customer_id"`
		CashIn        float64   `json:"cash_in"`
		CashOut       float64   `json:"cash_out"`
		AccountNumber string    `json:"account_number"`
		IsForPayment  bool      `json:"is_for_payment"`
	}
	guestBankAccountResp struct {
		ID            uuid.UUID `json:"id"`
		CreateTime    time.Time `json:"create_time"`
		UpdateTime    time.Time `json:"update_time"`
		CustomerID    uuid.UUID `json:"customer_id"`
		AccountNumber string    `json:"account_number"`
		IsForPayment  bool      `json:"is_for_payment"`
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
	transactionCreateResp struct {
		*transactionResp
		Token string `json:"token"`
	}
	debtResp struct {
		ID                        uuid.UUID       `json:"id"`
		CreateTime                time.Time       `json:"create_time"`
		UpdateTime                time.Time       `json:"update_time"`
		OwnerBankAccountNumber    string          `json:"owner_bank_account_number"`
		OwnerBankName             string          `json:"owner_bank_name"`
		OwnerName                 string          `json:"owner_name"`
		OwnerID                   *uuid.UUID      `json:"owner_id"`
		ReceiverBankAccountNumber string          `json:"receiver_bank_account_number"`
		ReceiverBankName          string          `json:"receiver_bank_name"`
		ReceiverName              string          `json:"receiver_name"`
		ReceiverID                *uuid.UUID      `json:"receiver_id"`
		TransactionID             *uuid.UUID      `json:"transaction_id"`
		Status                    debt.Status     `json:"status"`
		Description               string          `json:"description"`
		Amount                    decimal.Decimal `json:"amount"`
	}
	contactResp struct {
		ID            uuid.UUID `json:"id"`
		CreateTime    time.Time `json:"create_time"`
		UpdateTime    time.Time `json:"update_time"`
		OwnerID       uuid.UUID `json:"owner_id"`
		AccountNumber string    `json:"account_number"`
		SuggestName   string    `json:"suggest_name"`
		BankName      string    `json:"bank_name"`
	}
	tokenPairResponse struct {
		AccessToken  *string `json:"access_token"`
		RefreshToken *string `json:"refresh_token"`
	}

	// reference on docs
)

func getDefaultResponse(entity any) any {
	var result any
	switch entity.(type) {
	case *model.Customer:
		rs, _ := entity.(*model.Customer)
		result = &meResponse{
			ID:          rs.ID,
			Username:    rs.Username,
			FirstName:   rs.FirstName,
			LastName:    rs.LastName,
			PhoneNumber: rs.PhoneNumber,
			Email:       rs.Email,
			IsActive:    rs.IsActive,
		}
	case *model.BankAccount:
		rs, _ := entity.(*model.BankAccount)
		result = &bankAccountResp{
			ID:            rs.ID,
			CreateTime:    rs.CreateTime,
			UpdateTime:    rs.UpdateTime,
			CustomerID:    rs.CustomerID,
			CashIn:        rs.CashIn,
			CashOut:       rs.CashOut,
			AccountNumber: rs.AccountNumber,
			IsForPayment:  rs.IsForPayment,
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
	case *model.TransactionCreateResp:
		rs, _ := entity.(*model.TransactionCreateResp)
		result = &transactionCreateResp{
			transactionResp: &transactionResp{
				ID:                        rs.Transaction.ID,
				CreateTime:                rs.Transaction.CreateTime,
				UpdateTime:                rs.Transaction.UpdateTime,
				SourceTransactionID:       rs.Transaction.SourceTransactionID,
				Status:                    rs.Transaction.Status,
				ReceiverBankAccountNumber: rs.Transaction.ReceiverBankAccountNumber,
				ReceiverBankName:          rs.Transaction.ReceiverBankName,
				ReceiverName:              rs.Transaction.ReceiverName,
				ReceiverID:                rs.Transaction.ReceiverID,
				SenderBankAccountNumber:   rs.Transaction.SenderBankAccountNumber,
				SenderName:                rs.Transaction.SenderName,
				SenderID:                  rs.Transaction.SenderID,
				Amount:                    rs.Transaction.Amount,
				TransactionType:           rs.Transaction.TransactionType,
				Description:               rs.Transaction.Description,
			},
			Token: rs.Token,
		}
	case *model.Debt:
		rs, _ := entity.(*model.Debt)
		result = &debtResp{
			ID:                        rs.ID,
			CreateTime:                rs.CreateTime,
			UpdateTime:                rs.UpdateTime,
			Status:                    rs.Status,
			ReceiverBankAccountNumber: rs.ReceiverBankAccountNumber,
			ReceiverBankName:          rs.ReceiverBankName,
			ReceiverName:              rs.ReceiverName,
			ReceiverID:                rs.ReceiverID,
			OwnerBankAccountNumber:    rs.OwnerBankAccountNumber,
			OwnerName:                 rs.OwnerName,
			OwnerID:                   rs.OwnerID,
			Amount:                    rs.Amount,
			Description:               rs.Description,
			TransactionID:             rs.TransactionID,
		}
	case *model.Contact:
		rs, _ := entity.(*model.Contact)
		result = &contactResp{
			ID:            rs.ID,
			CreateTime:    rs.CreateTime,
			UpdateTime:    rs.UpdateTime,
			OwnerID:       rs.OwnerID,
			AccountNumber: rs.AccountNumber,
			SuggestName:   rs.SuggestName,
			BankName:      rs.BankName,
		}

	case *jwt.TokenPair:
		rs, _ := entity.(*jwt.TokenPair)
		result = &tokenPairResponse{
			AccessToken:  rs.AccessToken,
			RefreshToken: rs.RefreshToken,
		}
	default:
		result = entity
	}
	return result
}

func getGuestBankAccountResp(entity any) any {
	e := entity.(*model.BankAccount)
	return &guestBankAccountResp{
		ID:            e.ID,
		CreateTime:    e.CreateTime,
		UpdateTime:    e.UpdateTime,
		CustomerID:    e.CustomerID,
		AccountNumber: e.AccountNumber,
		IsForPayment:  e.IsForPayment,
	}
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

func getResponses[ModelType any](entities []ModelType, args ...func(any) any) any {
	fr := make([]any, 0, len(entities))

	for _, entity := range entities {
		fr = append(fr, getResponse(entity, args...))
	}
	return &EntitiesResponseTemplate[any]{Results: fr}
}
