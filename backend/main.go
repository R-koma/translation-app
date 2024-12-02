// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// func main() {

// 	r := gin.Default()

// 	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
// }

package main

import (
	"github.com/R-koma/translation-app/backend/controllers"
	"github.com/R-koma/translation-app/backend/infra"
	"github.com/R-koma/translation-app/backend/repositories"
	"github.com/R-koma/translation-app/backend/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()
	authRepository := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)

	r := gin.Default()
	r.Use(cors.Default())
	authRouter := r.Group("/auth")
	authRouter.POST("/signup", authController.Signup)
	authRouter.POST("/login", authController.Login)

	r.Run("localhost:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
