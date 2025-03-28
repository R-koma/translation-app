package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/R-koma/translation-app/backend/services"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService services.IAuthService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		log.Printf("Authorization header: %s", header)
		if header == "" {
			log.Println("Authorization header is missing")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(header, "Bearer ") {
			log.Println("Authorization header does not have 'Bearer ' prefix")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(header, "Bearer ")
		log.Printf("Extracted token: %s", tokenString)
		user, err := authService.GetUserFromToken(tokenString)
		if err != nil {
			log.Printf("Error in GetUserFromToken: %v", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		log.Printf("Authenticated user: %+v", user)
		ctx.Set("user", user)

		ctx.Next()
	}
}
