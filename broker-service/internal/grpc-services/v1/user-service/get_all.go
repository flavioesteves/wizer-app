package userservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	pb "github.com/flavioesteves/wizer-app/broker/proto/user"
)

func GetAllUsers(c *gin.Context, sc pb.UserServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := sc.GetAllUsers(ctx, &pb.GetAllUsersRequest{})

	for i, user := range res.Users {
		userRole, err := mapRoleToProtoBuf(user.Role.String())
		if err != nil {
			continue
		}
		fmt.Println(userRole)
		res.Users[i].Role = userRole
	}

	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
	}
	fmt.Printf("All Users were retrieved: %v\n", res)

	c.JSON(http.StatusOK, res)
}
