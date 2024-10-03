package routineservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	pbR "github.com/flavioesteves/wizer-app/broker/proto/routine"
)

func AddExerciseToRoutine(c *gin.Context, sc pbR.RoutineServiceClient) {
	var exerciceIds []string

	if err := c.BindJSON(&exerciceIds); err != nil {
		fmt.Println("error to retrieve exerciseIds %w\n", err)
		return
	}

	routineId := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.AddExerciseToRoutine(ctx, &pbR.AddExerciseToRoutineRequest{
		RoutineId:   routineId,
		ExerciseIds: exerciceIds,
	})

	if err != nil {
		fmt.Println("unexpected error %w\n", err)
	}

	c.JSON(http.StatusOK, res)
}
