package authservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/flavioesteves/wizer-app/broker/api/v1/models"
	pb "github.com/flavioesteves/wizer-app/broker/proto"
)

func ValidateToken(c *gin.Context, sc pb.AuthenticationServiceClient) {
	var reqToken models.RequestAuthToken

	if err := c.BindJSON(&reqToken); err != nil {
		fmt.Printf("Error to create a validate_token in the handler: %v\n", err)
		return
	}

	reqValidateToken := &pb.ValidateTokenRequest{
		Token: reqToken.Token,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.ValidateToken(ctx, reqValidateToken)
	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
	fmt.Printf("New validation was retrieved: %v\n", res)

	c.JSON(http.StatusOK, res)

}
