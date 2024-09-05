package authservice

import (
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup) {
	authServiceClient := Connect()
	rg.POST("", func(ctx *gin.Context) { LoginUser(ctx, authServiceClient) })
	rg.GET("/validate", func(ctx *gin.Context) { ValidateToken(ctx, authServiceClient) })
}
