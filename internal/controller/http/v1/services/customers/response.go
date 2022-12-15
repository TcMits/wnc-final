package customers

import (
	"github.com/TcMits/wnc-final/pkg/entity/model"
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

	// reference on docs
)

func getResponse(entity any) any {
	var result any
	switch entity.(type) {
	default:
		result = entity
	}
	return result
}

func getEntityResponse(entity any) any {
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
	default:
		result = entity
	}
	return result
}

func getEntitiesResponse[ModelType any](entities []ModelType) any {
	fr := make([]any, 0, len(entities))

	for _, entity := range entities {
		fr = append(fr, getEntityResponse(entity))
	}
	return &EntitiesResponseTemplate[any]{Results: fr}
}
