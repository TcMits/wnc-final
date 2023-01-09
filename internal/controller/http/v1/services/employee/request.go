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
)

func newListRequest() *listRequest {
	return &listRequest{
		Limit: 10,
	}
}