package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenykendf/go-restful/internal/pkg/reason"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": reason.InternalServerError})
			}
		}()
		ctx.Next()
	}

}
