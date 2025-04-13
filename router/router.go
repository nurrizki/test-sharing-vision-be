package router

import (
	"test-sharing-vision/internal/config"
	"test-sharing-vision/router/middleware"
	"test-sharing-vision/service/handler"
	"time"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() error {
	router := gin.Default()
	router.Use(middleware.HandlerOptions)
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Authorization", "Content-Type", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	article := router.Group("/article")
	article.Use(middleware.TokenAuthentication())
	{
		article.POST("", handler.InsertData)
		article.GET("", handler.GetData)
		article.GET("/:id", handler.GetDataById)
		article.DELETE("/:id", handler.DeleteDataById)
		article.PUT("/:id", handler.UpdateDataById)
	}

	return router.Run(":" + config.Cfg.App.Port)
}
