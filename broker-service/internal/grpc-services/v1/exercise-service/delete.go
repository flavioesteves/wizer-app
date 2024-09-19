package exerciseservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "github.com/flavioesteves/wizer-app/broker/proto/exercise"
	"github.com/gin-gonic/gin"
)

func DeleteExercise(c *gin.Context, sc pb.ExerciseServiceClient) {
	exerciseId := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.DeleteExercise(ctx, &pb.DeleteExerciseRequest{Id: exerciseId})

	if err != nil {
		fmt.Printf("Error in deleting Exercise with id: %s", exerciseId)
	}
	fmt.Printf("Exercise was deleted with id: %s", exerciseId)

	c.JSON(http.StatusOK, res)
}
