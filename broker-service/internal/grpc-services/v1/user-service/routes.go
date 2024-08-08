package userservice

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userServiceClient := Connect()

	rg.POST("", func(c *gin.Context) { CreateUser(c, userServiceClient) })
}
