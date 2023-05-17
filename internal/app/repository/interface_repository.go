package repository

import "github.com/kenykendf/go-restful/internal/app/model"

// type ICategoryRepo interface {
// 	Browse() ([]model.Category, error)
// 	Create(category model.Category) error
// 	Detail(id string) (model.Category, error)
// 	Update(id string, category model.Category) error
// 	Delete(id string) error
// }

// type IProductRepo interface {
// 	Create(product model.Product) error
// 	Browse() ([]model.Product, error)
// }

type IUserService interface {
	Create(req model.User) error
	GetByEmail(email string) (model.User, error)
}
