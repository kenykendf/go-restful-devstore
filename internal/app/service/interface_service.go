package service

import (
	"github.com/kenykendf/go-restful/internal/app/schema"
)

type ICategoryService interface {
	BrowseAll() ([]schema.GetCategoryResp, error)
	Create(req *schema.CreateCategoryReq) error
	DetailCategory(id string) (schema.GetCategoryResp, error)
	UpdateCategory(id string, req schema.UpdateCategoryReq) error
	DeleteCategory(id string) error
}

type IProductService interface {
	Create(req schema.CreateProductReq) error
	BrowseAll() ([]schema.GetProductResp, error)
}

type IUserService interface {
	Create(req *schema.CreateUser) error
}
