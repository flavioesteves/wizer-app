package exerciseservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "github.com/flavioesteves/wizer-app/broker/proto/exercise"
	"github.com/gin-gonic/gin"
)

func GetAllExercises(c *gin.Context, sc pb.ExerciseServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("Broker: All GetAllExercises was called")

	res, err := sc.GetAllExercises(ctx, &pb.GetAllExercisesRequest{})
	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
	fmt.Printf("All Exercises were retrieved: %v\n", res)

	c.JSON(http.StatusOK, res)
}
