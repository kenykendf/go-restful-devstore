package repository

import (
	"errors"
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

// get list user
func (ur *UserRepository) Browse() ([]model.User, error) {
	var (
		users        []model.User
		sqlStatement = `
			SELECT id, username, email
			FROM users
		`
	)

	rows, err := ur.DB.Queryx(sqlStatement)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - Browse : %w", err))
		return users, err
	}

	for rows.Next() {
		var user model.User
		err := rows.StructScan(&user)
		if err != nil {
			log.Error(fmt.Errorf("error UserRepository - Browse : %w", err))
		}
		users = append(users, user)
	}

	return users, nil
}

// get detail user
func (ur *UserRepository) GetByEmailAndUsername(email string, username string) (model.User, error) {
	var (
		sqlStatement = `
			SELECT id, username, email
			FROM users
			WHERE email = $1 AND username = $2
			LIMIT 1
		`
		user model.User
	)
	err := ur.DB.QueryRowx(sqlStatement, email, username).StructScan(&user)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - GetByEmailAndUsername : %w", err))
		return user, err
	}

	return user, nil
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

func (ur *UserRepository) Detail(userID int) (model.User, error) {
	var (
		sqlStatement = `
			SELECT id, email, hashed_password, username
			FROM users
			WHERE id = $1
			LIMIT 1
		`
		user model.User
	)
	err := ur.DB.QueryRowx(sqlStatement, userID).StructScan(&user)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - GetByEmail : %w", err))
		return user, err
	}

	return user, nil
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

// update user by id
func (ur *UserRepository) Update(user model.User) error {
	var (
		sqlStatement = `
			UPDATE users
			SET updated_at = NOW(),
				username = $2,
				email = $3
			WHERE id = $1
		`
	)

	result, err := ur.DB.Exec(sqlStatement, user.ID, user.Username, user.Email)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - UpdateByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}

// delete user by id
func (ur *UserRepository) DeleteByID(id string) error {
	var (
		sqlStatement = `
			DELETE FROM users
			WHERE id = $1
		`
	)

	result, err := ur.DB.Exec(sqlStatement, id)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - DeleteByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}
