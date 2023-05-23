package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenykendf/go-restful/internal/app/schema"
	"github.com/kenykendf/go-restful/internal/pkg/handler"
)

type ProductService interface {
	Create(req *schema.CreateProductReq) error
	BrowseAll(req *schema.BrowseProductReq) ([]schema.BrowseProductResp, error)
	Detail(id string) (schema.DetailProductResp, error)
	UpdateByID(id string, req *schema.UpdateProductReq) error
	Delete(id string) error
}

type ProductController struct {
	service ProductService
}

func NewProductController(service ProductService) *ProductController {
	return &ProductController{service: service}
}

// browse product
func (cc *ProductController) BrowseProduct(ctx *gin.Context) {
	req := &schema.BrowseProductReq{}
	req.Page = ctx.GetInt("page")
	req.PageSize = ctx.GetInt("page_size")

	resp, err := cc.service.BrowseAll(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", resp)
}

func (pc *ProductController) Create(ctx *gin.Context) {
	req := &schema.CreateProductReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	// validate image types (only : jpeg & png)

	// validate image size (max 1MB)

	err := pc.service.Create(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success create product", nil)
}

// get detail product
func (cc *ProductController) Detail(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	resp, err := cc.service.Detail(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusOK, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "", resp)
}

// update product by id
func (cc *ProductController) UpdateProduct(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	req := &schema.UpdateProductReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := cc.service.UpdateByID(id, req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success update product", nil)
}

// delete product by id
func (cc *ProductController) Delete(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	err := cc.service.Delete(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success delete product", nil)
}
