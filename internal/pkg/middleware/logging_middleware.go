package middleware

import (
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		requestMethod := ctx.Request.Method
		reqURI := ctx.Request.RequestURI
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()

		log.WithFields(log.Fields{
			"latency_time":   latencyTime,
			"reqeust_method": requestMethod,
			"req_uri":        reqURI,
			"status_code":    statusCode,
			"client_ip":      clientIP,
		}).Info("http request")

		ctx.Next()
	}
}
