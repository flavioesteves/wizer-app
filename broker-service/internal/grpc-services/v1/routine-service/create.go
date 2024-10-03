package routineservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/flavioesteves/wizer-app/broker/api/v1/models"
	pb "github.com/flavioesteves/wizer-app/broker/proto/model"
	pbR "github.com/flavioesteves/wizer-app/broker/proto/routine"
)

func CreateRoutine(c *gin.Context, sc pbR.RoutineServiceClient) {
	var requestRoutine models.RequestRoutine

	if err := c.BindJSON(&requestRoutine); err != nil {
		//TODO: Replace those logs with a proper log service
		fmt.Printf("Error to create a routine in routine_handler %v\n", err)
		return
	}

	newRoutine := &pb.Routine{
		ProfileId: requestRoutine.ProfileId,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.CreateRoutine(ctx, &pbR.CreateRoutineRequest{Routine: newRoutine})
	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
	fmt.Printf("Routine was been created: %v\n", res)

	c.JSON(http.StatusOK, res)
}
