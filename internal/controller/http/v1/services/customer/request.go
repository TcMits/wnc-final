package customer

import (
	"github.com/TcMits/wnc-final/ent/debt"
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
	bankAccountUpdateReq struct {
		IsForPayment bool `json:"is_for_payment" validate:"required"`
	}
	transactionCreateReq struct {
		ReceiverID    *uuid.UUID      `json:"receiver_id" validate:"required"`
		Amount        decimal.Decimal `json:"amount" validate:"required"`
		Description   string          `json:"description" validate:"required"`
		IsFeePaidByMe bool            `json:"is_fee_paid_by_me" validate:"required"`
	}
	tpBankTransactionCreateReq struct {
		Amount        decimal.Decimal `json:"amount" validate:"required"`
		Description   string          `json:"description" validate:"required"`
		AccountNumber string          `json:"account_number" validate:"required"`
		IsFeePaidByMe bool            `json:"is_fee_paid_by_me" validate:"required"`
	}
	transactionConfirmReq struct {
		Token string `json:"token" validate:"required"`
		OTP   string `json:"otp" validate:"required"`
	}
	debtCreateReq struct {
		ReceiverBankAccountNumber string          `json:"receiver_bank_account_number" validate:"required"`
		ReceiverName              string          `json:"receiver_name" validate:"required"`
		ReceiverID                uuid.UUID       `json:"receiver_id" validate:"required"`
		Description               *string         `json:"description" validate:"required"`
		Amount                    decimal.Decimal `json:"amount" validate:"required"`
	}
	debtCancelReq struct {
		Description string `json:"description" validate:"required"`
	}
	contactCreateReq struct {
		AccountNumber string `json:"account_number," validate:"required"`
		SuggestName   string `json:"suggest_name" validate:"required"`
		BankName      string `json:"bank_name" validate:"required"`
	}
	contactUpdateReq struct {
		AccountNumber string `json:"account_number"`
		SuggestName   string `json:"suggest_name"`
		BankName      string `json:"bank_name"`
	}
	changePasswordReq struct {
		OldPassword     string `json:"old_password," validate:"required"`
		Password        string `json:"password," validate:"required"`
		ConfirmPassword string `json:"confirm_password," validate:"required"`
	}
	forgetPasswordReq struct {
		Email string `json:"email," validate:"required"`
	}
	changePasswordWithTokenReq struct {
		Token           string `json:"token," validate:"required"`
		Otp             string `json:"otp," validate:"required"`
		Password        string `json:"password," validate:"required"`
		ConfirmPassword string `json:"confirm_password," validate:"required"`
	}
	debtFulfillReq struct {
		Token string `json:"token," validate:"required"`
		Otp   string `json:"otp," validate:"required"`
	}
	transactionFilterReq struct {
		OnlyDebt    bool `url:"only_debt"`
		OnlyReceive bool `url:"only_receive"`
		OnlySend    bool `url:"only_send"`
	}
	bankAccountFilterReq struct {
		AccountNumber *string `url:"account_number"`
	}
	debtFilterReq struct {
		OwnerID    *uuid.UUID   `url:"owner_id"`
		ReceiverID *uuid.UUID   `url:"receiver_id"`
		Status     *debt.Status `url:"status"`
	}
	transactionOrderReq struct {
		UpdateTimeAsc  bool `url:"update_time"`
		UpdateTimeDesc bool `url:"-update_time"`
	}
)

func newListRequest() *listRequest {
	return &listRequest{
		Limit: 10,
	}
}
