package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kenykendf/go-restful/internal/app/model"
	log "github.com/sirupsen/logrus"
)

type CategoryRepo struct {
	DB *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) *CategoryRepo {
	return &CategoryRepo{DB: db}
}

func (cr *CategoryRepo) Create(category model.Category) error {
	var (
		sqlStatement = `
			INSERT INTO categories (name, description)
			VALUES ($1, $2)
		`
	)
	_, err := cr.DB.Exec(sqlStatement, category.Name, category.Description)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CategoryRepo) Browse(search model.BrowseCategory) ([]model.Category, error) {
	var (
		limit        = search.PageSize
		offset       = limit * (search.Page - 1)
		categories   []model.Category
		sqlStatement = `
			SELECT id, name, description
			FROM categories
			LIMIT $1
			OFFSET $2
		`
	)

	rows, err := cr.DB.Queryx(sqlStatement, limit, offset)
	if err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - Browse : %w", err))
		return categories, err
	}

	for rows.Next() {
		var category model.Category
		err := rows.StructScan(&category)
		if err != nil {
			log.Error(fmt.Errorf("error CategoryRepository - Browse : %w", err))
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (cr *CategoryRepo) Detail(id string) (model.Category, error) {
	var (
		category     model.Category
		sqlStatement = `
			SELECT id, name, description
			FROM categories
			WHERE id = $1
			AND deleted_at IS NULL
		`
	)
	rows := cr.DB.QueryRow(sqlStatement, id)

	err := rows.Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func (cr *CategoryRepo) Update(category model.Category) error {
	var (
		sqlStatement = `
			UPDATE categories
			SET updated_at = now(), name = $2, description = $3
			WHERE id = $1
			`
	)
	_, err := cr.DB.Exec(sqlStatement, category.ID, category.Name, category.Description)
	if err != nil {
		log.Error(fmt.Errorf("error updating category : %w", err))
		return err
	}

	return nil
}

func (cr *CategoryRepo) Delete(id string) error {
	var (
		sqlStatement = `
			UPDATE categories
			SET updated_at = now(), deleted_at = now()
			WHERE id = $1
			`
	)
	_, err := cr.DB.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	return nil
}
