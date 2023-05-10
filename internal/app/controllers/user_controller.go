package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenykendf/go-restful/internal/app/schema"
	"github.com/kenykendf/go-restful/internal/app/service"
	"github.com/kenykendf/go-restful/internal/pkg/handler"
)

type UserController struct {
	service service.IUserService
}

func NewUserController(service service.IUserService) *UserController {
	return &UserController{service: service}
}

func (uc *UserController) Create(ctx *gin.Context) {
	req := &schema.CreateUser{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := uc.service.Create(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "user created", nil)
}
