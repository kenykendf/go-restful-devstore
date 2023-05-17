package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenykendf/go-restful/internal/app/schema"
	"github.com/kenykendf/go-restful/internal/pkg/handler"
	"github.com/kenykendf/go-restful/internal/pkg/reason"
	"github.com/kenykendf/go-restful/internal/pkg/validator"
	"github.com/sirupsen/logrus"
)

type CategoryService interface {
	Create(req *schema.CreateCategoryReq) error
	BrowseAll(req *schema.BrowseCategoryReq) ([]schema.GetCategoryResp, error)
	DetailCategory(id string) (schema.GetCategoryResp, error)
	UpdateCategory(id string, req *schema.UpdateCategoryReq) error
	DeleteCategory(id string) error
}

type CategoryController struct {
	service CategoryService
}

func NewCategoryController(service CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (cc *CategoryController) CreateCategory(ctx *gin.Context) {
	req := &schema.CreateCategoryReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := cc.service.Create(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "category created"})
}

func (cc *CategoryController) BrowseCategory(ctx *gin.Context) {
	req := &schema.BrowseCategoryReq{}
	req.Page = ctx.GetInt("page")
	req.PageSize = ctx.GetInt("page_size")

	resp, err := cc.service.BrowseAll(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", resp)
}

func (cc *CategoryController) DetailCategory(ctx *gin.Context) {
	categoryID := ctx.Param("id")

	resp, err := cc.service.DetailCategory(categoryID)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", resp)
}

func (cc *CategoryController) UpdateCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	req := &schema.UpdateCategoryReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	validate := validator.Check(&req)
	if validate {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.CategoryCannotUpdate)
		return
	}

	err := cc.service.UpdateCategory(id, req)
	if err != nil {
		logrus.Error(fmt.Errorf("error updating category : %w", err))
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "category has been updated", nil)
}

func (cc *CategoryController) DeleteCategory(ctx *gin.Context) {
	categoryID := ctx.Param("id")

	err := cc.service.DeleteCategory(categoryID)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "category has been deleted", nil)

}
