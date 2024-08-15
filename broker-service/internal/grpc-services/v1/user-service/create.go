package userservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/flavioesteves/wizer-app/broker/api/v1/models"
	pb "github.com/flavioesteves/wizer-app/broker/proto"
)

func CreateUser(c *gin.Context, sc pb.UserServiceClient) {
	var requestUser models.RequestUser

	if err := c.BindJSON(&requestUser); err != nil {
		//TODO: Replace those logs with a proper log service
		fmt.Printf("Error to create a User in user_handler %v\n", err)
		return
	}

	userRole, err := mapRoleToProtoBuf(requestUser.Role)
	if err != nil {
		fmt.Printf("Role was incorretly filled with a bad argument %s", userRole)
		return
	}

	newUser := &pb.User{
		UserId:       requestUser.UserID,
		Email:        requestUser.Email,
		PasswordHasg: requestUser.PasswordHash,
		Role:         userRole,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.CreateUser(ctx, &pb.CreateUserRequest{User: newUser})
	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
	fmt.Printf("User was been created: %v\n", res)

	c.JSON(http.StatusOK, res)

}
