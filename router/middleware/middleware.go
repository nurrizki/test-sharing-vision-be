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
