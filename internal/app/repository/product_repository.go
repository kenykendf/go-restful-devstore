package repository

import (
	"errors"
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

func (pr *ProductRepo) Browse(search model.BrowseProduct) ([]model.Product, error) {
	var (
		limit        = search.PageSize
		offset       = limit * (search.Page - 1)
		products     []model.Product
		sqlStatement = `
			SELECT id, name, description, currency, total_stock, is_active, category_id, image_url
			FROM product
			LIMIT $1
			OFFSET $2
		`
	)

	rows, err := pr.DB.Queryx(sqlStatement, limit, offset)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - Browse : %w", err))
		return products, err
	}

	for rows.Next() {
		var product model.Product
		err := rows.StructScan(&product)
		if err != nil {
			log.Error(fmt.Errorf("error ProductRepository - Browse : %w", err))
		}
		products = append(products, product)
	}

	return products, nil
}

// get detail product
func (pr *ProductRepo) Detail(id string) (model.Product, error) {
	var (
		sqlStatement = `
			SELECT id, name, description, currency, total_stock, is_active, category_id, image_url
			FROM product
			WHERE id = $1
		`
		product model.Product
	)
	err := pr.DB.QueryRowx(sqlStatement, id).StructScan(&product)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - Detail : %w", err))
		return product, err
	}

	return product, nil
}

// update article by id
func (pr *ProductRepo) Update(product model.Product) error {
	var (
		sqlStatement = `
			UPDATE product
			SET updated_at = NOW(),
				name = $2,
				description = $3,
				currency = $4,
				total_stock = $5,
				is_active = $6,
				category_id	 = $7
			WHERE id = $1
		`
	)

	result, err := pr.DB.Exec(sqlStatement,
		product.ID,
		product.Name,
		product.Description,
		product.Currency,
		product.TotalStock,
		product.IsActive,
		product.CategoryID,
	)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - Update : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}

// delete article by id
func (pr *ProductRepo) Delete(id string) error {
	var (
		sqlStatement = `
			DELETE FROM product
			WHERE id = $1
		`
	)

	result, err := pr.DB.Exec(sqlStatement, id)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - Delete : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}
