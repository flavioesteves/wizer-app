package routineservice

import "github.com/gin-gonic/gin"

func RegisterRoutineRoutes(rg *gin.RouterGroup) {
	routineServiceClient := Connect()

	rg.GET("/", func(ctx *gin.Context) { GetAllRoutines(ctx, routineServiceClient) })
	rg.POST("/", func(ctx *gin.Context) { CreateRoutine(ctx, routineServiceClient) })
	rg.PUT("/", func(ctx *gin.Context) { UpdateRoutine(ctx, routineServiceClient) })

	rg.GET("/:id", func(ctx *gin.Context) { GetRoutine(ctx, routineServiceClient) })
	rg.DELETE("/:id", func(ctx *gin.Context) { DeleteRoutine(ctx, routineServiceClient) })

	// m2m routines <-> exercises
	rg.GET("/:id/exercises", func(ctx *gin.Context) { GetExercisesByRoutineId(ctx, routineServiceClient) })
	rg.POST("/:id/exercises", func(ctx *gin.Context) { AddExerciseToRoutine(ctx, routineServiceClient) })
	rg.DELETE("/:id/exercises/:exerciseId", func(ctx *gin.Context) { RemoveExerciseFromRoutine(ctx, routineServiceClient) })
}
