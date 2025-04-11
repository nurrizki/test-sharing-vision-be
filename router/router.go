package router

import (
	"net/http"
	"test-sharing-vision/internal/config"
	"test-sharing-vision/router/middleware"
	"test-sharing-vision/service/handler"
	"time"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() error {
	router := gin.Default()

	api := router.Group("")
	{
		//version
		api.GET("/version", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{
				"version": "v1.0.0",
				"time":    time.Now().Unix(),
				"status":  "active",
			}})
		})
	}

	article := router.Group("/article")
	{
		//router configuration
		article.Use(cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
			AllowedHeaders:   []string{"Origin", "Authorization", "Content-Type", "Content-Length", "X-Requested-With", "Rest-Api-Key", "Client-Id", "Client-Secret", "Client-Key"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}), middleware.TokenAuthentication())
		{
			article.POST("", handler.InsertData)
			article.GET("", handler.GetData)
			article.GET("/:id", handler.GetDataById)
			article.DELETE("/:id", handler.DeleteDataById)
			article.PUT("/:id", handler.UpdateDataById)
		}

	}

	return router.Run(":" + config.Cfg.App.Port)
}
