package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenykendf/go-restful/internal/app/schema"
	"github.com/kenykendf/go-restful/internal/app/service"
)

type ProductController struct {
	service service.IProductService
}

func NewProductController(service service.IProductService) *ProductController {
	return &ProductController{service: service}
}

func (pc *ProductController) Create(ctx *gin.Context) {
	var req schema.CreateProductReq

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	err = pc.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "product created"})
}

func (pc *ProductController) BrowseProduct(ctx *gin.Context) {
	resp, err := pc.service.BrowseAll()
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}
