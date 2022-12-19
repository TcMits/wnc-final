package customer

import "github.com/google/uuid"

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
	updateRequest struct {
		id *uuid.UUID `param:"id" validate:"required"`
	}
)

type (
	bankAccountUpdateRequest struct {
		IsForPayment bool `json:"is_for_payment validate:"required"`
	}
)

func newListRequest() *listRequest {
	return &listRequest{
		Limit: 10,
	}
}
