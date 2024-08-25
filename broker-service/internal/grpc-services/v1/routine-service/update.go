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

func UpdateRoutine(c *gin.Context, sc pb.RoutineServiceClient) {
	var requestRoutine models.RequestRoutine

	if err := c.BindJSON(&requestRoutine); err != nil {
		fmt.Printf("Error to create a routine in routine_handler %v\n", err)
		return
	}

	updateRoutine := &pb.Routine{
		Id:        requestRoutine.Id,
		ProfileId: requestRoutine.ProfileId,
		Exercises: make([]*pb.Exercise, 0, len(requestRoutine.Exercises)), // Pre-allocate slice to populate the fiels with the logic
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.UpdateRoutine(ctx, &pb.UpdateRoutineRequest{Routine: updateRoutine})
	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
	fmt.Printf("Routine was been updated: %v\n", res)

	c.JSON(http.StatusOK, res)
}
