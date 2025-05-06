package main

import (
	"github.com/R-koma/translation-app/backend/controllers"
	"github.com/R-koma/translation-app/backend/db"
	"github.com/R-koma/translation-app/backend/middlewares"
	"github.com/R-koma/translation-app/backend/models"
	"github.com/R-koma/translation-app/backend/repositories"
	"github.com/R-koma/translation-app/backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db := db.SetupDB()
	db.AutoMigrate(&models.User{}, &models.FriendRequest{})

	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	friendRequestRepository := repositories.NewFriendRequestRepository(db)
	friendRequestService := services.NewFriendRequestService(friendRequestRepository)
	friendRequestController := controllers.NewFriendRequestController(friendRequestService)

	r := gin.Default()
	r.Use(cors.Default())
	authRouter := r.Group("/auth")
	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)
	authRouter.GET("/profile", authController.Profile)
	authRouter.GET("/users", authController.GetAllUsers)

	friendsRouterWithAuth := r.Group("/friend", middlewares.AuthMiddleware(authService))

	friendsRouterWithAuth.POST("/requests", friendRequestController.CreateFriendRequest)
	friendsRouterWithAuth.GET("/requests", friendRequestController.GetFriendRequests)
	friendsRouterWithAuth.PATCH("/requests/:id", friendRequestController.UpdateFriendRequestStatus)
	friendsRouterWithAuth.GET("/friends", friendRequestController.GetMyFriends)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World! This is a backend server!")
	})

	r.Run("0.0.0.0:8080")
}
