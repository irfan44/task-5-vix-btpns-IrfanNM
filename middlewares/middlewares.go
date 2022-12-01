package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irfan44/task-5-vix-btpns-IrfanNurghiffariM/app"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token, err := app.ValidateToken(authHeader)
		if token.Valid {
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
	}
}
