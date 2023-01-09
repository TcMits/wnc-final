package admin

import (
	"time"

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
		DateStart Time    `url:"date_start,int"`
		DateEnd   Time    `url:"date_end,int"`
		BankName  *string `url:"bank_name"`
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
func newTransactionFilterReq() *transactionFilterReq {
	start := time.Now()
	end := start.Add(time.Hour * 672)
	return &transactionFilterReq{
		DateStart: Time{t: start},
		DateEnd:   Time{t: end},
	}
}
