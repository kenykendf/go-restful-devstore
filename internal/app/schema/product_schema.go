package schema

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
	Name        string
	Description string
	Currency    string
	Price       int
	TotalStock  int
	IsActive    bool
	CategoryID  int
}
