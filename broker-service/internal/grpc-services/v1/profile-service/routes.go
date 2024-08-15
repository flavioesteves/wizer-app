package profileservice

import "github.com/gin-gonic/gin"

func RegisterProfileRoutes(rg *gin.RouterGroup) {
	profileServiceClient := Connect()

	rg.GET("", func(ctx *gin.Context) { GetAllProfiles(ctx, profileServiceClient) })
	rg.POST("", func(ctx *gin.Context) { CreateProfile(ctx, profileServiceClient) })
	rg.PUT("", func(ctx *gin.Context) { UpdateProfile(ctx, profileServiceClient) })

	rg.GET("/:id", func(ctx *gin.Context) { GetProfile(ctx, profileServiceClient) })
	rg.DELETE("/:id", func(ctx *gin.Context) { DeleteProfile(ctx, profileServiceClient) })
}
