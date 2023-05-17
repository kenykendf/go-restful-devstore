package service

import "github.com/kenykendf/go-restful/internal/app/model"

type ProductRepository interface {
	Create(product model.Product) error
	Browse(search model.BrowseProduct) ([]model.Product, error)
	Detail(id string) (model.Product, error)
	Update(product model.Product) error
	Delete(id string) error
}
