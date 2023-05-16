package repository

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/kenykendf/go-restful/internal/app/model"
	"github.com/sirupsen/logrus"
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

func (cr *CategoryRepo) Browse() ([]model.Category, error) {
	var (
		categories   []model.Category
		sqlStatement = `
			SELECT id, name, description
			FROM categories
			WHERE deleted_at IS NULL
		`
	)

	rows, err := cr.DB.Queryx(sqlStatement)
	if err != nil {
		log.Print(fmt.Errorf("error exec query statement : %w", err))
		return categories, err
	}

	for rows.Next() {
		var category model.Category
		_ = rows.StructScan(&category)
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
		fmt.Println("ERR REPO ", err)
		log.Println("Unable to fetch category detail")
		return model.Category{}, err
	}
	return category, nil
}

func (cr *CategoryRepo) Update(id string, category model.Category) error {
	var (
		sqlStatement = `
			UPDATE categories
			SET updated_at = now(), name = $2, description = $3
			WHERE id = $1
			`
	)
	_, err := cr.DB.Exec(sqlStatement, id, category.Name, category.Description)
	if err != nil {
		logrus.Error(fmt.Errorf("error updating category : %w", err))

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
		log.Print("Delete Failed , ", err)
		return err
	}

	return nil
}
