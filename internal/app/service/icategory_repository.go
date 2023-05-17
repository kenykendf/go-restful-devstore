package service

import "github.com/kenykendf/go-restful/internal/app/model"

type CategoryRepository interface {
	Create(category model.Category) error
	Browse(category model.BrowseCategory) ([]model.Category, error)
	Update(category model.Category) error
	Detail(id string) (model.Category, error)
	Delete(id string) error
}
