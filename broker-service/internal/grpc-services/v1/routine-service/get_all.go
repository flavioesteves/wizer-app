package routineservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "github.com/flavioesteves/wizer-app/broker/proto"
	"github.com/gin-gonic/gin"
)

func GetAllRoutines(c *gin.Context, sc pb.RoutineServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.GetAllRoutines(ctx, &pb.GetAllRoutinesRequest{})
	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
	fmt.Printf("All Routines were retrieved: %v\n", res)

	c.JSON(http.StatusOK, res)
}
