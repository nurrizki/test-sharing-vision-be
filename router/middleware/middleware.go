package middleware

import (
	"net/http"
	"strings"
	"test-sharing-vision/service/model"
	"time"

	"github.com/gin-gonic/gin"
)

func TokenAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Time", time.Now().Unix())
		headerAuth := c.Request.Header.Get("Authorization")

		if headerAuth == "" {
			responseUnauthorized(c)
		}

		if strings.HasPrefix(headerAuth, "Basic ") {
			valid := isValidBasicAuth(headerAuth)
			if !valid {
				responseUnauthorized(c)
			}
		}

	}
}

func responseUnauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, model.GeneralResponse{
		Status:  "400",
		Message: "unauthorized token",
	})
	c.Abort()
}

func HandlerOptions(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Origin, Authorization, Content-Type, X-Requested-With, Rest-Api-Key, Client-Id, Client-Secret, Client-Key")
	c.Header("Content-Type", "application/json")
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
