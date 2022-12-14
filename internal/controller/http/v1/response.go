package v1

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
