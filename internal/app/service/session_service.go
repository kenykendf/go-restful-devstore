package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/kenykendf/go-restful/internal/app/model"
	"github.com/kenykendf/go-restful/internal/app/schema"
	"github.com/kenykendf/go-restful/internal/pkg/reason"
	"golang.org/x/crypto/bcrypt"

	log "github.com/sirupsen/logrus"
)

type UserRepo interface {
	GetByEmail(email string) (model.User, error)
}

type AuthRepo interface {
	Create(auth model.Auth) error
}

type TokenGenerator interface {
	GenerateAccessToken(userID int) (string, time.Time, error)
	GenerateRefreshToken(userID int) (string, time.Time, error)
}

type SessionService struct {
	userRepo   UserRepo
	authRepo   AuthRepo
	tokenMaker TokenGenerator
}

func NewSessionService(
	userRepo UserRepo,
	authRepo AuthRepo,
	tokenMaker TokenGenerator,
) *SessionService {
	return &SessionService{
		userRepo:   userRepo,
		authRepo:   authRepo,
		tokenMaker: tokenMaker,
	}
}

func (ss *SessionService) Login(req *schema.LoginReq) (schema.LoginResp, error) {
	var resp schema.LoginResp

	existingUser, _ := ss.userRepo.GetByEmail(req.Email)
	if existingUser.ID < 0 {
		log.Error(fmt.Errorf("unable to get user by email"))
		return resp, errors.New(reason.FailedLogin)
	}

	// verify password
	isVerified := ss.verifyPassword(existingUser.Password, req.Password)
	if !isVerified {
		log.Error(fmt.Errorf("failed to verify password"))

		return resp, errors.New(reason.FailedLogin)
	}

	// generate access token
	accessToken, _, err := ss.tokenMaker.GenerateAccessToken(existingUser.ID)
	if err != nil {
		log.Error(fmt.Errorf("access token err : %w", err))
		return resp, errors.New(reason.FailedLogin)
	}

	// generate refresh token
	refreshToken, _, err := ss.tokenMaker.GenerateRefreshToken(existingUser.ID)
	if err != nil {
		log.Error(fmt.Errorf("refresh token err : %w", err))
		return resp, errors.New(reason.FailedLogin)
	}

	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	// save refresh token
	authPayload := model.Auth{
		UserID:   existingUser.ID,
		Token:    refreshToken,
		AuthType: "refresh_token",
	}
	err = ss.authRepo.Create(authPayload)
	if err != nil {
		log.Error(fmt.Errorf("refresh token saving : %w", err))
		return resp, errors.New(reason.FailedLogin)
	}

	return resp, nil
}

func (ss *SessionService) Logout() {

}

func (ss *SessionService) Refresh() {

}

func (ss *SessionService) verifyPassword(hashPass string, plainPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(plainPass))
	if err != nil {
		log.Error(fmt.Errorf("err : %w", err))
	}
	return err == nil
}