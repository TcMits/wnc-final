package customer

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/customer/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type contactRoute struct {
	uc     usecase.ICustomerContactUseCase
	logger logger.Interface
}

func RegisterContactController(handler iris.Party, l logger.Interface, uc usecase.ICustomerContactUseCase) {
	h := handler.Party("/")
	route := &contactRoute{
		uc:     uc,
		logger: l,
	}
	h.Use(middleware.Authenticator(uc.GetSecret(), uc.GetUser))
	h.Get("/contacts/{id:uuid}", route.detail)
	h.Put("/contacts/{id:uuid}", route.update)
	h.Delete("/contacts/{id:uuid}", route.delete)
	h.Get("/contacts", route.listing)
	h.Post("/contacts", route.create)
	h.Options("/contacts", func(_ iris.Context) {})
	h.Head("/contacts", func(_ iris.Context) {})
}

// @Summary     Get a contact
// @Description Get a contact
// @ID          contact-get
// @Tags  	    Contact
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       id path string true "ID of contact"
// @Success     200 {object} contactResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /contacts/{id} [get]
func (s *contactRoute) detail(ctx iris.Context) {
	req := new(detailRequest)
	if err := ctx.ReadParams(req); err != nil {
		handleBindingError(ctx, err, s.logger, req, nil)
		return
	}
	entity, err := s.uc.GetFirstMine(ctx, nil, &model.ContactWhereInput{ID: req.id})
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	if entity != nil {
		ctx.JSON(getResponse(entity))
	} else {
		ctx.StatusCode(iris.StatusNoContent)
	}
}

// @Summary     Show contact
// @Description Show contact
// @ID          contact-listing
// @Tags  	    Contact
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Success     200 {object} EntitiesResponseTemplate[contactResp]
// @Failure     500 {object} errorResponse
// @Router      /contacts [get]
func (s *contactRoute) listing(ctx iris.Context) {
	req := newListRequest()
	if err := ctx.ReadQuery(req); err != nil {
		handleBindingError(ctx, err, s.logger, req, nil)
		return
	}
	entities, err := s.uc.ListMine(ctx, &req.Limit, &req.Offset, nil, nil)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	isNext, err := s.uc.IsNext(ctx, req.Limit, req.Offset, nil, nil)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	paging := getPagingResponse(ctx, pagingInput[*model.Contact]{
		limit:    req.Limit,
		offset:   req.Offset,
		entities: entities,
		isNext:   isNext,
	})
	ctx.JSON(paging)
}

// @Summary     Create a contact
// @Description Create a contact
// @ID          contact-create
// @Tags  	    Contact
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       payload body contactCreateReq true "Create a contact"
// @Success     201 {object} contactResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /contacts [post]
func (s *contactRoute) create(ctx iris.Context) {
	createInReq := new(contactCreateReq)
	if err := ctx.ReadBody(createInReq); err != nil {
		handleBindingError(ctx, err, s.logger, createInReq, nil)
		return
	}
	in := &model.ContactCreateInput{
		AccountNumber: createInReq.AccountNumber,
		SuggestName:   createInReq.SuggestName,
		BankName:      createInReq.BankName,
	}
	in, err := s.uc.ValidateCreate(ctx, in)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	entity, err := s.uc.Create(ctx, in)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	ctx.JSON(getResponse(entity))
}

// @Summary     Update a contact
// @Description Update a contact
// @ID          contact-update
// @Tags  	    Contact
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       payload body contactUpdateReq true "Update a contact"
// @Param       id path string true "ID of contact"
// @Success     200 {object} contactResp
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /contacts/{id} [put]
func (r *contactRoute) update(ctx iris.Context) {
	req := new(detailRequest)
	if err := ctx.ReadParams(req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	updateInReq := new(contactUpdateReq)
	if err := ctx.ReadBody(updateInReq); err != nil {
		handleBindingError(ctx, err, r.logger, updateInReq, nil)
		return
	}
	entity, err := r.uc.GetFirstMine(ctx, nil, &model.ContactWhereInput{ID: req.id})
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	if entity == nil {
		ctx.StatusCode(iris.StatusNoContent)
	}
	i := &model.ContactUpdateInput{
		AccountNumber: &updateInReq.AccountNumber,
		SuggestName:   &updateInReq.SuggestName,
		BankName:      &updateInReq.BankName,
	}
	i, err = r.uc.ValidateUpdate(ctx, entity, i)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	entity, err = r.uc.Update(ctx, entity, i)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	ctx.JSON(getResponse(entity))
}

// @Summary     Delete a contact
// @Description Delete a contact
// @ID          contact-delete
// @Tags  	    Contact
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       id path string true "ID of contact"
// @Success     204 ""
// @Failure     500 {object} errorResponse
// @Router      /contacts/{id} [delete]
func (r *contactRoute) delete(ctx iris.Context) {
	req := new(detailRequest)
	if err := ctx.ReadParams(req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	entity, err := r.uc.GetFirstMine(ctx, nil, &model.ContactWhereInput{ID: req.id})
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	if entity == nil {
		ctx.StatusCode(iris.StatusNoContent)
	}
	err = r.uc.Delete(ctx, entity)
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	ctx.StatusCode(iris.StatusNoContent)
}
