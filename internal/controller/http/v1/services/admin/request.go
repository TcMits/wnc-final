package admin

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
	transactionFilterReq struct {
		DateStart *Time   `url:"date_start,int"`
		DateEnd   *Time   `url:"date_end,int"`
		BankName  *string `url:"bank_name"`
	}
	transactionOrderReq struct {
		UpdateTimeAsc  bool `url:"update_time"`
		UpdateTimeDesc bool `url:"-update_time"`
	}
	employeeCreateReq struct {
		Username  string `json:"username" validate:"required"`
		FirstName string `json:"first_name" validate:"required"`
		LastName  string `json:"last_name" validate:"required"`
	}
	employeeUpdateReq struct {
		Username  *string `json:"username"`
		FirstName *string `json:"first_name"`
		LastName  *string `json:"last_name"`
	}
)

func newListRequest() *listRequest {
	return &listRequest{
		Limit: 10,
	}
}
