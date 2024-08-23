package routineservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/flavioesteves/wizer-app/broker/api/v1/models"
	pb "github.com/flavioesteves/wizer-app/broker/proto"
)

func CreateRoutine(c *gin.Context, sc pb.RoutineServiceClient) {
	var requestRoutine models.RequestRoutine

	if err := c.BindJSON(&requestRoutine); err != nil {
		//TODO: Replace those logs with a proper log service
		fmt.Printf("Error to create a routine in routine_handler %v\n", err)
		return
	}

	newRoutine := &pb.Routine{
		Id:        requestRoutine.Id,
		Exercises: requestRoutine.Exercises, //TODO: update to be an array
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.CreateRoutine(ctx, &pb.CreateRoutineRequest{Routine: newRoutine})
	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
	fmt.Printf("Routine was been created: %v\n", res)

	c.JSON(http.StatusOK, res)
}
