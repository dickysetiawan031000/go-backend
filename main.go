package main

import (
	"time"

	"github.com/dickysetiawan031000/go-backend/handler"
	"github.com/dickysetiawan031000/go-backend/middleware"
	"github.com/dickysetiawan031000/go-backend/repository"
	"github.com/dickysetiawan031000/go-backend/usecase"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"https://frontend-app-dicky.vercel.app",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authUsecase := usecase.NewAuthUseCase()
	itemRepo := repository.NewItemRepository()
	itemUsecase := usecase.NewItemUseCase(itemRepo)

	// Group /api
	api := r.Group("/api")

	// Auth
	authGroup := api.Group("/auth")
	authHandler := handler.NewAuthHandler(authGroup, authUsecase)

	// Protected routes
	protected := api.Group("/")
	protected.Use(middleware.JWTAuthMiddleware())
	protected.GET("/profile", authHandler.Profile)
	protected.PUT("/profile", authHandler.UpdateProfile)
	protected.POST("/logout", authHandler.Logout)

	// Item
	itemGroup := protected.Group("")
	handler.NewItemHandler(itemGroup, itemUsecase)

	r.Run(":8080")
}
