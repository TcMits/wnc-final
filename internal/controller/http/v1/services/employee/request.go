package employee

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
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
		id uuid.UUID `param:"id" validate:"required"`
	}
)

type (
	customerCreateReq struct {
		Username    string `json:"username" validate:"required,min=6"`
		Email       string `json:"email" validate:"required,email"`
		PhoneNumber string `json:"phone_number" validate:"required,min=12"`
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

func ReadID(ctx iris.Context, req *detailRequest) error {
	id := ctx.Params().Get("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	req.id = uid
	return nil
}
