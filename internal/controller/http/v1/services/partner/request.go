package partner

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
		ApiKey *string `json:"api_key" validate:"required"`
	}
	detailRequest struct {
		id *uuid.UUID `param:"id" validate:"required"`
	}
)

type (
	transactionCreateReq struct {
		Amount                    decimal.Decimal `json:"amount" validate:"required"`
		Description               string          `json:"description" validate:"required"`
		Token                     string          `json:"token" validate:"required"`
		Signature                 string          `json:"signature" validate:"required"`
		SenderName                string          `json:"sender_name" validate:"required"`
		SenderBankAccountNumber   string          `json:"sender_bank_account_number" validate:"required"`
		ReceiverBankAccountNumber string          `json:"receiver_bank_account_number" validate:"required"`
	}
	bankAccountFilterReq struct {
		AccountNumber *string `url:"account_number"`
	}
)

func newListRequest() *listRequest {
	return &listRequest{
		Limit: 10,
	}
}
