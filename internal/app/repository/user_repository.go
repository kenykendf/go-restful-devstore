package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kenykendf/go-restful/internal/app/model"

	log "github.com/sirupsen/logrus"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Create(user model.User) error {
	var (
		sqlStatement = `
			INSERT INTO users (username, email, password)
			VALUES ($1, $2, $3)
		`
	)
	_, err := ur.DB.Exec(sqlStatement, user.Username, user.Email, user.Password)
	fmt.Println("USER ", user)

	if err != nil {
		log.Error(fmt.Errorf("create user : %w", err))
		return err
	}

	return nil
}

func (ur *UserRepository) GetByEmail(email string) (model.User, error) {
	var (
		sqlStatement = `
			SELECT id, username, email, password
			FROM users
			WHERE email = $1
			LIMIT 1
		`
		user model.User
	)

	err := ur.DB.QueryRowx(sqlStatement, email).StructScan(&user)
	if err != nil {
		log.Error(fmt.Errorf("GetByEmail : %w", err))
		return user, err
	}

	return user, nil
}
