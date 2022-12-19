package customer

import (
	"time"

	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/jwt"
	"github.com/google/uuid"
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
	tokenPairResponse struct {
		AccessToken  *string `json:"access_token"`
		RefreshToken *string `json:"refresh_token"`
	}

	// reference on docs
)

func getResponse(entity any) any {
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

func getResponses[ModelType any](entities []ModelType) any {
	fr := make([]any, 0, len(entities))

	for _, entity := range entities {
		fr = append(fr, getResponse(entity))
	}
	return &EntitiesResponseTemplate[any]{Results: fr}
}
