package app

import (
	"github.com/gin-gonic/gin"
	"go-test2/cmd"
)

func StartServer(b *Bank, config cmd.Config) error {
	r := gin.Default()
	secured := r.Use(NewAuthFunc(config.ApiToken))

	return nil
}
