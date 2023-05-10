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
