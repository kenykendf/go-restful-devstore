package repository

import (
	"github.com/kenykendf/go-restful/internal/app/model"
)

type ICategoryRepo interface {
	Browse() ([]model.Category, error)
	Create(category model.Category) error
	Detail(id string) (model.Category, error)
	Update(id string, category model.Category) error
	Delete(id string) error
}
