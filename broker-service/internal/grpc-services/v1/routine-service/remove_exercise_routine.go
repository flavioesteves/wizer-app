package routineservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	pbR "github.com/flavioesteves/wizer-app/broker/proto/routine"
)

func RemoveExerciseFromRoutine(c *gin.Context, sc pbR.RoutineServiceClient) {
	routineId := c.Param("id")
	exerciseId := c.Param("exerciseId")

	var exerciseIds []string

	exerciseIds = append(exerciseIds, exerciseId)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.RemoveExerciseFromRoutine(ctx, &pbR.RemoveExerciseFromRoutineRequest{
		RoutineId:   routineId,
		ExerciseIds: exerciseIds,
	})

	if err != nil {
		fmt.Errorf("error in deleting exercises: %w from routine %s\n", exerciseIds, routineId)
	}

	c.JSON(http.StatusOK, res)

}
