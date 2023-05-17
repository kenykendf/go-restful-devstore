package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kenykendf/go-restful/internal/app/model"
	log "github.com/sirupsen/logrus"
)

type AuthRepo struct {
	DB *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{DB: db}
}

func (ar *AuthRepo) Find(userID int, refreshToken string) (model.Auth, error) {
	var (
		sqlStatement = `
			SELECT id, token, auth_type, user_id, expired_at
			FROM auth
			WHERE user_id = $1 AND token = $2
		`
		auth model.Auth
	)

	err := ar.DB.QueryRowx(sqlStatement, userID, refreshToken).StructScan(&auth)
	if err != nil {
		log.Error(fmt.Errorf("error AuthRepository - Find : %w", err))
		return auth, err
	}

	return auth, nil
}

func (ar *AuthRepo) Create(auth model.Auth) error {
	var (
		sqlStatement = `
			INSERT INTO auth (token, auth_type, user_id)
			VALUES($1, $2, $3)
		`
	)

	_, err := ar.DB.Exec(sqlStatement, auth.Token, auth.AuthType, auth.UserID)
	if err != nil {
		log.Error(fmt.Errorf("error AuthRepo - Create : %w", err))
		return err
	}
	return nil
}

func (ar *AuthRepo) DeleteAllByUserID(userID int) error {
	var (
		sqlStatement = `
			DELETE FROM auth
			WHERE user_id = $1
		`
	)

	_, err := ar.DB.Exec(sqlStatement, userID)
	if err != nil {
		log.Error(fmt.Errorf("error AuthRepository - DeleteAllByUserID : %w", err))
		return err
	}

	return nil
}
