package model

type Product struct {
	ID          int
	Name        string
	Description string
	Currency    string
	Price       int
	TotalStock  int
	IsActive    bool
	CategoryID  int
}
