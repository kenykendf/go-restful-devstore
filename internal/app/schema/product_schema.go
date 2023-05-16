package schema

import "mime/multipart"

type GetProductResp struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Currency    string `json:"currency"`
	Price       int    `json:"price"`
	TotalStock  int    `json:"total_stock"`
	IsActive    bool   `json:"is_active"`
	CategoryID  int    `json:"category_id"`
}

type CreateProductReq struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	Currency    string                `form:"currency"`
	Price       int                   `form:"price"`
	TotalStock  int                   `form:"total_stock"`
	IsActive    bool                  `form:"is_active"`
	CategoryID  int                   `form:"category_id"`
	Image       *multipart.FileHeader `form:"image"`
}
