package model

type Category struct {
	ID          int
	Name        string
	Description string
}

type BrowseCategory struct {
	Page     int
	PageSize int
}
