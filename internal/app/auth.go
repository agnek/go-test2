package app

import "github.com/gin-gonic/gin"

func NewAuthFunc(token string) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("auth-token")

		if authHeader != token {
			context.AbortWithStatus(400)
			return
		}

		context.Next()
	}
}
