package service

import "github.com/kenykendf/go-restful/internal/app/schema"

type ICategoryService interface {
	BrowseAll() ([]schema.GetCategoryResp, error)
	Create(req schema.CreateCategoryReq) error
	DetailCategory(id string) (schema.GetCategoryResp, error)
	UpdateCategory(id string, req schema.CreateCategoryReq) error
	DeleteCategory(id string) error
}
