package employee

import (
	"strconv"

	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/tool/jwt"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

type (
	EntitiesResponseTemplate[EntityResponse any] struct {
		Count    uint             `json:"count"`
		Next     string           `json:"next"`
		Previous string           `json:"previous"`
		Results  []EntityResponse `json:"results"`
	}
	pagingInput[ModelType any] struct {
		limit        int
		offset       int
		noPagingResp *EntitiesResponseTemplate[any]
		isNext       bool
		entities     []ModelType
	}
	emptyResponse struct{}

	// error
	errorResponse struct {
		Message string `json:"message"`
		Code    string `json:"code"`
		Detail  string `json:"detail"`
	}
	customerResponse struct {
		ID          uuid.UUID `json:"id"`
		Username    string    `json:"username"`
		FirstName   string    `json:"first_name"`
		LastName    string    `json:"last_name"`
		PhoneNumber string    `json:"phone_number"`
		Email       string    `json:"email"`
		IsActive    bool      `json:"is_active"`
	}
	meResponse struct {
		ID        uuid.UUID `json:"id"`
		Username  string    `json:"username"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		IsActive  bool      `json:"is_active"`
	}
	tokenPairResponse struct {
		AccessToken  *string `json:"access_token"`
		RefreshToken *string `json:"refresh_token"`
	}
	// reference on docs
)

func getDefaultResponse(entity any) any {
	var result any
	switch entity.(type) {
	case *model.Customer:
		rs, _ := entity.(*model.Customer)
		result = &customerResponse{
			ID:          rs.ID,
			Username:    rs.Username,
			FirstName:   rs.FirstName,
			LastName:    rs.LastName,
			PhoneNumber: rs.PhoneNumber,
			Email:       rs.Email,
			IsActive:    rs.IsActive,
		}
	case *model.Employee:
		rs, _ := entity.(*model.Employee)
		result = &meResponse{
			ID:        rs.ID,
			Username:  rs.Username,
			FirstName: rs.FirstName,
			LastName:  rs.LastName,
			IsActive:  rs.IsActive,
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

func getResponse(entity any, args ...func(any) any) any {
	var result any = entity
	if len(args) == 0 {
		args = append(args, getDefaultResponse)
	}
	for _, t := range args {
		result = t(result)
	}
	return result
}

func getResponses[ModelType any](entities []ModelType, args ...func(any) any) *EntitiesResponseTemplate[any] {
	fr := make([]any, 0, len(entities))

	for _, entity := range entities {
		fr = append(fr, getResponse(entity, args...))
	}
	return &EntitiesResponseTemplate[any]{Results: fr}
}

func getPagingResponse[ModelType any](ctx iris.Context, i pagingInput[ModelType], args ...func(any) any) *EntitiesResponseTemplate[any] {
	var pageResp *EntitiesResponseTemplate[any]
	if i.noPagingResp != nil {
		pageResp = i.noPagingResp
	} else {
		pageResp = getResponses(i.entities, args...)
	}
	if i.isNext {
		originUrl := ctx.Request().URL
		url := *originUrl
		q := url.Query()
		q.Set("limit", strconv.Itoa(i.limit))
		offset := i.offset + i.limit
		q.Set("offset", strconv.Itoa(offset))
		url.RawQuery = q.Encode()
		pageResp.Next = url.String()
	}
	if i.offset >= i.limit {
		originUrl := ctx.Request().URL
		url := *originUrl
		q := url.Query()
		q.Set("limit", strconv.Itoa(i.limit))
		offset := i.offset - i.limit
		q.Set("offset", strconv.Itoa(offset))
		url.RawQuery = q.Encode()
		pageResp.Previous = url.String()
	}
	pageResp.Count = uint(len(pageResp.Results))
	return pageResp
}
