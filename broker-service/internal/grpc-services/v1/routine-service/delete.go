package routineservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	pb "github.com/flavioesteves/wizer-app/broker/proto"
)

func DeleteRoutine(c *gin.Context, sc pb.RoutineServiceClient) {
	routineId := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.DeleteRoutine(ctx, &pb.DeleteRoutineRequest{Id: routineId})

	if err != nil {
		fmt.Printf("Profile was deleted with id: %s", routineId)
	}
	fmt.Printf("Profile was deleted with id: %s", routineId)

	c.JSON(http.StatusOK, res)
}