package admin

import (
	"github.com/TcMits/wnc-final/internal/controller/http/v1/services/admin/middleware"
	"github.com/TcMits/wnc-final/internal/usecase"
	"github.com/TcMits/wnc-final/pkg/entity/model"
	"github.com/TcMits/wnc-final/pkg/infrastructure/logger"
	"github.com/kataras/iris/v12"
)

type employeeRoute struct {
	uc     usecase.IAdminEmployeeUseCase
	logger logger.Interface
}

func RegisterEmployeeController(handler iris.Party, l logger.Interface, uc usecase.IAdminEmployeeUseCase) {
	h := handler.Party("/")
	route := &employeeRoute{
		uc:     uc,
		logger: l,
	}
	h.Use(middleware.Authenticator(uc.GetSecret(), uc.GetUser))
	h.Get("/employees/{id:uuid}", route.detail)
	h.Put("/employees/{id:uuid}", route.update)
	h.Delete("/employees/{id:uuid}", route.delete)
	h.Get("/employees", route.listing)
	h.Post("/employees", route.create)
	h.Options("/employees", func(_ iris.Context) {})
	h.Head("/employees", func(_ iris.Context) {})
}

// @Summary     Get a employee
// @Description Get a employee
// @ID          employee-get
// @Tags  	    Employee
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       id path string true "ID of employee"
// @Success     200 {object} employeeResponse
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /employees/{id} [get]
func (s *employeeRoute) detail(ctx iris.Context) {
	req := new(detailRequest)
	if err := ReadID(ctx, req); err != nil {
		handleBindingError(ctx, err, s.logger, req, nil)
		return
	}
	entity, err := s.uc.GetFirst(ctx, nil, &model.EmployeeWhereInput{ID: &req.id})
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

// @Summary     Show employee
// @Description Show employee
// @ID          employee-listing
// @Tags  	    Employee
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Success     200 {object} EntitiesResponseTemplate[employeeResponse]
// @Failure     500 {object} errorResponse
// @Router      /employees [get]
func (s *employeeRoute) listing(ctx iris.Context) {
	req := newListRequest()
	if err := ctx.ReadQuery(req); err != nil {
		handleBindingError(ctx, err, s.logger, req, nil)
		return
	}
	entities, err := s.uc.List(ctx, &req.Limit, &req.Offset, nil, nil)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	isNext, err := s.uc.IsNext(ctx, req.Limit, req.Offset, nil, nil)
	if err != nil {
		HandleError(ctx, err, s.logger)
		return
	}
	paging := getPagingResponse(ctx, pagingInput[*model.Employee]{
		limit:    req.Limit,
		offset:   req.Offset,
		entities: entities,
		isNext:   isNext,
	})
	ctx.JSON(paging)
}

// @Summary     Create a employee
// @Description Create a employee
// @ID          employee-create
// @Tags  	    Employee
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       payload body employeeCreateReq true "Create a employee"
// @Success     201 {object} employeeResponse
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /employees [post]
func (s *employeeRoute) create(ctx iris.Context) {
	createInReq := new(employeeCreateReq)
	if err := ctx.ReadBody(createInReq); err != nil {
		handleBindingError(ctx, err, s.logger, createInReq, nil)
		return
	}
	in := &model.EmployeeCreateInput{
		Username:  createInReq.Username,
		FirstName: &createInReq.FirstName,
		LastName:  &createInReq.LastName,
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

// @Summary     Update a employee
// @Description Update a employee
// @ID          employee-update
// @Tags  	    Employee
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       payload body employeeUpdateReq true "Update a employee"
// @Param       id path string true "ID of employee"
// @Success     200 {object} employeeResponse
// @Failure     400 {object} errorResponse
// @Failure     500 {object} errorResponse
// @Router      /employees/{id} [put]
func (r *employeeRoute) update(ctx iris.Context) {
	req := new(detailRequest)
	if err := ReadID(ctx, req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	updateInReq := new(employeeUpdateReq)
	if err := ctx.ReadBody(updateInReq); err != nil {
		handleBindingError(ctx, err, r.logger, updateInReq, nil)
		return
	}
	entity, err := r.uc.GetFirst(ctx, nil, &model.EmployeeWhereInput{ID: &req.id})
	if err != nil {
		HandleError(ctx, err, r.logger)
		return
	}
	if entity == nil {
		ctx.StatusCode(iris.StatusNoContent)
	}
	i := new(model.EmployeeUpdateInput)
	if updateInReq.Username != nil {
		i.Username = updateInReq.Username
	}
	if updateInReq.FirstName != nil {
		i.FirstName = updateInReq.FirstName
	}
	if updateInReq.LastName != nil {
		i.LastName = updateInReq.LastName
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

// @Summary     Delete a employee
// @Description Delete a employee
// @ID          employee-delete
// @Tags  	    Employee
// @Security 	Bearer
// @Accept      json
// @Produce     json
// @Param       id path string true "ID of employee"
// @Success     204 ""
// @Failure     500 {object} errorResponse
// @Router      /employees/{id} [delete]
func (r *employeeRoute) delete(ctx iris.Context) {
	req := new(detailRequest)
	if err := ReadID(ctx, req); err != nil {
		handleBindingError(ctx, err, r.logger, req, nil)
		return
	}
	entity, err := r.uc.GetFirst(ctx, nil, &model.EmployeeWhereInput{ID: &req.id})
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
