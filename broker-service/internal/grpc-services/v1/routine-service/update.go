package routineservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/flavioesteves/wizer-app/broker/api/v1/models"
	pb "github.com/flavioesteves/wizer-app/broker/proto/model"
	pbr "github.com/flavioesteves/wizer-app/broker/proto/routine"
)

func UpdateRoutine(c *gin.Context, sc pbr.RoutineServiceClient) {
	var requestRoutine models.RequestRoutine

	if err := c.BindJSON(&requestRoutine); err != nil {
		fmt.Printf("Error to create a routine in routine_handler %v\n", err)
		return
	}

	updateRoutine := &pb.Routine{
		Id:        requestRoutine.Id,
		Name:      requestRoutine.Name,
		ProfileId: requestRoutine.ProfileId,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.UpdateRoutine(ctx, &pbr.UpdateRoutineRequest{Routine: updateRoutine})
	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
	fmt.Printf("Routine was been updated: %v\n", res)

	c.JSON(http.StatusOK, res)
}
