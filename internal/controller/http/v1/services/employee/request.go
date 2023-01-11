package employee

import (
	"github.com/google/uuid"
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
	customerCreateReq struct {
		Username    string `json:"username" validate:"required"`
		Email       string `json:"email" validate:"required,email"`
		PhoneNumber string `json:"phone_number" validate:"required"`
	}
	bankAccountFilterReq struct {
		AccountNumber *string `url:"account_number"`
		Username      *string `url:"username"`
	}
	bankAccountUpdateReq struct {
		CashIn *float64 `json:"cash_in" validate:"required"`
	}
	transactionFilterReq struct {
		CustomerID  *uuid.UUID `url:"customer_id" validate:"required_with=OnlyDebt OnlyReceive OnlySend"`
		OnlyDebt    bool       `url:"only_debt"`
		OnlyReceive bool       `url:"only_receive"`
		OnlySend    bool       `url:"only_send"`
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
