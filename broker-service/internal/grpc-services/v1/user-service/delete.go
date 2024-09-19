package userservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	pb "github.com/flavioesteves/wizer-app/broker/proto/user"
)

func DeleteUser(c *gin.Context, sc pb.UserServiceClient) {
	userId := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.DeleteUser(ctx, &pb.DeleteUserRequest{UserId: userId})

	if err != nil {
		fmt.Printf("Unexpected Error %v\n", err)
	}
	fmt.Printf("User was deleted with id: %s", userId)

	c.JSON(http.StatusOK, res)
}
