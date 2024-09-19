package exerciseservice

import "github.com/gin-gonic/gin"

func RegisterExerciseRoutes(rg *gin.RouterGroup) {
	exerciseServiceClient := Connect()

	rg.GET("", func(ctx *gin.Context) { GetAllExercises(ctx, exerciseServiceClient) })
	rg.POST("", func(ctx *gin.Context) { CreateExercise(ctx, exerciseServiceClient) })
	rg.PUT("", func(ctx *gin.Context) { UpdateExercise(ctx, exerciseServiceClient) })

	rg.GET("/:id", func(ctx *gin.Context) { GetExercise(ctx, exerciseServiceClient) })
	rg.DELETE("/:id", func(ctx *gin.Context) { DeleteExercise(ctx, exerciseServiceClient) })

}
