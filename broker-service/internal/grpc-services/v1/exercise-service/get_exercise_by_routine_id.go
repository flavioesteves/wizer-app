package exerciseservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "github.com/flavioesteves/wizer-app/broker/proto/exercise"
	"github.com/gin-gonic/gin"
)

func GetRoutinesByExerciseId(c *gin.Context, sc pb.ExerciseServiceClient) {
	exerciseId := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.GetRoutinesByExerciseId(ctx, &pb.GetRoutinesByExerciseIdRequest{
		ExerciseId: exerciseId,
	})

	if err != nil {
		fmt.Println("error retrieving routines for exerciseId: %w", exerciseId)
	}

	c.JSON(http.StatusOK, res)
}
