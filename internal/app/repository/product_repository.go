package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kenykendf/go-restful/internal/app/model"

	log "github.com/sirupsen/logrus"
)

type ProductRepo struct {
	DB *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) *ProductRepo {
	return &ProductRepo{DB: db}
}

func (pr *ProductRepo) Create(product model.Product) error {
	var (
		sqlStatement = `
			INSERT INTO product (
				name
				,description
				,currency
				,price
				,total_stock
				,is_active
				,category_id
				,image_url
			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		`
	)

	_, err := pr.DB.Exec(sqlStatement,
		product.Name,
		product.Description,
		product.Currency,
		product.Price,
		product.TotalStock,
		product.IsActive,
		product.CategoryID,
		product.ImageURL,
	)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepo - Create : %w", err))
		return err
	}

	return nil
}

func (pr *ProductRepo) Browse() ([]model.Product, error) {
	var (
		products     []model.Product
		sqlStatement = `
			SELECT id, name, description, currency, price, total_stock, is_active, category_id
			FROM product
			WHERE deleted_at IS NULL
		`
	)

	rows, err := pr.DB.Queryx(sqlStatement)
	if err != nil {
		log.Error(fmt.Errorf("error exec query statement : %w", err))
		return products, err
	}

	for rows.Next() {
		var product model.Product
		_ = rows.StructScan(&product)
		products = append(products, product)
	}

	return products, nil
}
