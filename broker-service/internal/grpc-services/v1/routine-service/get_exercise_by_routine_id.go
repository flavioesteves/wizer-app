package routineservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "github.com/flavioesteves/wizer-app/broker/proto/routine"
	"github.com/gin-gonic/gin"
)

func GetExercisesByRoutineId(c *gin.Context, sc pb.RoutineServiceClient) {
	routineId := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.GetExercisesByRoutineId(ctx, &pb.GetExercisesByRoutineIdRequest{
		RoutineId: routineId,
	})

	if err != nil {
		fmt.Println("error in getting exercises for the routine id: %w", routineId)
	}

	c.JSON(http.StatusOK, res)
}
