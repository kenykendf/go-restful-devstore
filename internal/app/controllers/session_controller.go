package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenykendf/go-restful/internal/app/schema"
	"github.com/kenykendf/go-restful/internal/pkg/handler"
)

type SessionController struct {
	service SessionService
}

type SessionService interface {
	Login(req *schema.LoginReq) (schema.LoginResp, error)
}

func NewSessionController(service SessionService) *SessionController {
	return &SessionController{service: service}
}

func (sc *SessionController) Login(ctx *gin.Context) {
	req := &schema.LoginReq{}
	if handler.BindAndCheck(ctx, req) {
		return
	}

	resp, err := sc.service.Login(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "login success", resp)

}
