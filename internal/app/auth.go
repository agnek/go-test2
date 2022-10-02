package app

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func NewAuthFunc(token string) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("auth-token")
		log.Info().Str("token", authHeader).Msg("auth check")
		if authHeader != token {
			context.AbortWithStatus(400)
			return
		}

		context.Next()
	}
}
