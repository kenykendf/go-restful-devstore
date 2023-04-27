package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenykendf/go-restful/internal/app/schema"
	"github.com/kenykendf/go-restful/internal/app/service"
	"github.com/sirupsen/logrus"
)

type CategoryController struct {
	service service.ICategoryService
}

func NewCategoryController(service service.ICategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (cc *CategoryController) CreateCategory(ctx *gin.Context) {
	var req schema.CreateCategoryReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = cc.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "category created"})
}

func (cc *CategoryController) BrowseCategory(ctx *gin.Context) {
	resp, err := cc.service.BrowseAll()
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

func (cc *CategoryController) DetailCategory(ctx *gin.Context) {
	categoryID := ctx.Param("id")

	resp, err := cc.service.DetailCategory(categoryID)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

func (cc *CategoryController) UpdateCategory(ctx *gin.Context) {
	categoryID := ctx.Param("id")
	var req schema.CreateCategoryReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = cc.service.UpdateCategory(categoryID, req)
	if err != nil {
		logrus.Error(fmt.Errorf("error updating category : %w", err))
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "category has been updated"})
}

func (cc *CategoryController) DeleteCategory(ctx *gin.Context) {
	categoryID := ctx.Param("id")

	err := cc.service.DeleteCategory(categoryID)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "category has been deleted"})
}
