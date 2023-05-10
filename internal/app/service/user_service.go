package service

import (
	"errors"

	"github.com/kenykendf/go-restful/internal/app/model"
	"github.com/kenykendf/go-restful/internal/app/repository"
	"github.com/kenykendf/go-restful/internal/app/schema"
	"github.com/kenykendf/go-restful/internal/pkg/reason"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repository.IUserService
}

func NewUserService(userRepo repository.IUserService) *UserService {
	return &UserService{userRepo: userRepo}
}

func (us *UserService) Create(req *schema.CreateUser) error {
	// Check user
	isExistUser, _ := us.userRepo.GetByEmail(req.Email)
	if isExistUser.ID > 0 {
		return errors.New(reason.UserAlreadyExist)
	}

	// Hash Password
	pass, _ := us.hashPassword(req.Password)
	// Create User Data

	var insertData model.User
	insertData.Email = req.Email
	insertData.Password = pass
	insertData.Username = req.Username

	err := us.userRepo.Create(insertData)
	if err != nil {
		return errors.New(reason.RegisterFailed)
	}
	return nil
}

func (us *UserService) hashPassword(password string) (string, error) {
	bytePass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytePass), nil

}
