package userservice

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	profileServiceClient := Connect()

	rg.POST("", func(c *gin.Context) { CreateUser(c, profileServiceClient) })
}
