package customer

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type (
	listRequest struct {
		Limit  int `url:"limit"`
		Offset int `url:"offset"`
	}
	renewTokenRequest struct {
		RefreshToken *string `json:"refresh_token" validate:"required"`
	}
	loginRequest struct {
		Username *string `json:"username" validate:"required"`
		Password *string `json:"password" validate:"required"`
	}
	detailRequest struct {
		id *uuid.UUID `param:"id" validate:"required"`
	}
)

type (
	bankAccountUpdateRequest struct {
		IsForPayment bool `json:"is_for_payment" validate:"required"`
	}
	transactionCreateRequest struct {
		ReceiverBankAccountNumber string          `json:"receiver_bank_account_number" validate:"required"`
		ReceiverBankName          string          `json:"receiver_bank_name" validate:"required"`
		ReceiverName              string          `json:"receiver_name" validate:"required"`
		ReceiverID                *uuid.UUID      `json:"receiver_id" validate:"required"`
		Amount                    decimal.Decimal `json:"amount" validate:"required"`
		Description               string          `json:"description" validate:"required"`
		SenderID                  *uuid.UUID      `json:"sender_id" validate:"required"`
		IsFeePaidByMe             bool            `json:"is_fee_paid_by_me" validate:"required"`
	}
	transactionConfirmRequest struct {
		Token string `json:"token" validate:"required"`
	}
)

func newListRequest() *listRequest {
	return &listRequest{
		Limit: 10,
	}
}
