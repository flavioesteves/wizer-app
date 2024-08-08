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
		//TODO: Replace this logs with a proper log service
		fmt.Printf("Error to create a user in user_handler %v\n", err)
		return
	}

	newUser := &pb.User{
		Id:       requestUser.Id,
		Email:    requestUser.Email,
		Password: requestUser.Password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.CreateUser(ctx, newUser)
	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}

	fmt.Printf("User was been created: %v\n", res)

	c.IndentedJSON(http.StatusOK, res)
}
