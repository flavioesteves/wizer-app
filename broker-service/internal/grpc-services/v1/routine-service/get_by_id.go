package routineservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "github.com/flavioesteves/wizer-app/broker/proto"
	"github.com/gin-gonic/gin"
)

func GetRoutine(c *gin.Context, sc pb.RoutineServiceClient) {
	routineId := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.GetRoutine(ctx, &pb.GetRoutineRequest{Id: routineId})
	if err != nil {
		fmt.Printf("Unecpected error %v\n", err)
	}
	fmt.Printf("Routine was retrieved: %v\n", res)

	c.JSON(http.StatusOK, res)
}
