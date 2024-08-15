package userservice

import "github.com/gin-gonic/gin"

func RegisterUserRoutes(rg *gin.RouterGroup) {
	userServiceClient := Connect()

	rg.GET("", func(ctx *gin.Context) { GetAllUsers(ctx, userServiceClient) })
	rg.POST("", func(ctx *gin.Context) { CreateUser(ctx, userServiceClient) })
	rg.PUT("", func(ctx *gin.Context) { UpdateUser(ctx, userServiceClient) })

	rg.GET("/:id", func(ctx *gin.Context) { GetUser(ctx, userServiceClient) })
	rg.DELETE("/:id", func(ctx *gin.Context) { DeleteUser(ctx, userServiceClient) })
}
