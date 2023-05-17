package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kenykendf/go-restful/internal/app/schema"
	"github.com/kenykendf/go-restful/internal/pkg/handler"
	"github.com/kenykendf/go-restful/internal/pkg/reason"
)

type SessionController struct {
	service    SessionService
	tokenMaker RefreshTokenVerifier
}
type RefreshTokenVerifier interface {
	VerifyRefreshToken(tokenString string) (string, error)
}

type SessionService interface {
	Login(req *schema.LoginReq) (schema.LoginResp, error)
	Logout(UserID int) error
	Refresh(req *schema.RefreshTokenReq) (schema.RefreshTokenResp, error)
}

func NewSessionController(service SessionService, tokenMaker RefreshTokenVerifier) *SessionController {
	return &SessionController{
		service:    service,
		tokenMaker: tokenMaker,
	}
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

func (sc *SessionController) Refresh(ctx *gin.Context) {
	refreshToken := ctx.GetHeader("refresh_token")
	if refreshToken == "" {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "cannot refresh token")
	}

	sub, err := sc.tokenMaker.VerifyRefreshToken(refreshToken)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, reason.FailedRefreshToken)
		return
	}

	intSub, _ := strconv.Atoi(sub)
	req := &schema.RefreshTokenReq{}
	req.RefreshToken = refreshToken
	req.UserID = intSub

	resp, err := sc.service.Refresh(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.FailedRefreshToken)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success refresh", resp)
}

// logout
func (sc *SessionController) Logout(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.GetString("user_id"))
	err := sc.service.Logout(userID)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success logout", nil)
}
