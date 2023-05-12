package middleware

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/kenykendf/go-restful/internal/pkg/handler"
	"github.com/kenykendf/go-restful/internal/pkg/reason"
)

type Authorization struct{}
type AuthorizationMid interface{}

func AuthorizationMiddleware(
	sub, obj, act string,
	enforcer *casbin.Enforcer,
) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// sub := ctx.GetString("user_id")
		// enforcer
		ok, err := enforcer.Enforce(sub, obj, act)
		// check err
		if err != nil {
			handler.ResponseError(ctx, http.StatusUnauthorized, reason.InvalidAccess)
			ctx.Abort()
			return
		}
		if !ok {
			handler.ResponseError(ctx, http.StatusUnauthorized, reason.InvalidAccess)
			ctx.Abort()
			return
		}
		// check permission
	}
}
