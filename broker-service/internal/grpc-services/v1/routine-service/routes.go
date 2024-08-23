package routineservice

import "github.com/gin-gonic/gin"

func RegisterRoutineRoutes(rg *gin.RouterGroup) {
	routineServiceClient := Connect()

	rg.GET("", func(ctx *gin.Context) { GetAllRoutines(ctx, routineServiceClient) })
	rg.POST("", func(ctx *gin.Context) { CreateRoutine(ctx, routineServiceClient) })
	rg.PUT("", func(ctx *gin.Context) { UpdateRoutine(ctx, routineServiceClient) })

	rg.GET("/:id", func(ctx *gin.Context) { GetRoutine(ctx, routineServiceClient) })
	rg.DELETE("/:id", func(ctx *gin.Context) { DeleteProfile(ctx, routineServiceClient) })
}
