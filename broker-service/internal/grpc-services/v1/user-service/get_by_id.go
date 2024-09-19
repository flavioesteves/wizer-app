package userservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "github.com/flavioesteves/wizer-app/broker/proto/user"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context, sc pb.UserServiceClient) {
	userId := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.GetUser(ctx, &pb.GetUserRequest{UserId: userId})
	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
	fmt.Printf("Profile was retrieved: %v\n", res)

	c.JSON(http.StatusOK, res)
}
