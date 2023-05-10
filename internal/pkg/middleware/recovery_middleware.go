package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kenykendf/go-restful/internal/pkg/reason"
	log "github.com/sirupsen/logrus"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				log.Error(fmt.Errorf("err message : %s ", err))
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": reason.InternalServerError})
			}
		}()
		ctx.Next()
	}

}
