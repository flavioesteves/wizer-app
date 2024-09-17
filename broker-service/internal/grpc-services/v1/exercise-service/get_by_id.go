package exerciseservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "github.com/flavioesteves/wizer-app/broker/proto"
	"github.com/gin-gonic/gin"
)

func GetExercise(c *gin.Context, sc pb.ExerciseServiceClient) {
	exerciseId := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.GetExercise(ctx, &pb.GetExerciseRequest{Id: exerciseId})
	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
	fmt.Printf("Exercise was retrieved: %v\n", res)

	c.JSON(http.StatusOK, res)
}
