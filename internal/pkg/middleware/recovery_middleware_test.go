package middleware

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRecoveryMiddleware(t *testing.T) {
	// Test Cases
	type Testcase struct {
		Name            string
		ResponseCode    int
		ResponseMessage string
		FuncHandler     func(c *gin.Context)
	}

	// Cases
	cases := []Testcase{
		{
			Name:            "with error panic",
			ResponseCode:    http.StatusInternalServerError,
			ResponseMessage: "internal server error",
			FuncHandler: func(c *gin.Context) {
				panic("panic error")
			},
		},
		{
			Name:            "with error panic",
			ResponseCode:    http.StatusOK,
			ResponseMessage: "success",
			FuncHandler: func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "success"})
			},
		},
	}

	// Exec
	for _, v := range cases {
		t.Run(v.Name, func(t *testing.T) {
			// SETUP
			gin.SetMode(gin.TestMode)
			router := gin.New()
			router.Use(RecoveryMiddleware())
			router.GET("/recovery", v.FuncHandler)

			// PERFORM REQUEST
			response := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", "/recovery", nil)
			router.ServeHTTP(response, request)

			// PARSE RESPONSE
			var bodyJSON map[string]interface{}
			respBody, _ := io.ReadAll(response.Body)
			_ = json.Unmarshal(respBody, &bodyJSON)

			// CHECK ASSERT
			assert.Equal(t, v.ResponseCode, response.Code)
			assert.Equal(t, v.ResponseMessage, bodyJSON["message"])

		})
	}
}
