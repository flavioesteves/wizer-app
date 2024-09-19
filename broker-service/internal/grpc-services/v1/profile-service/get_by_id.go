package profileservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	pb "github.com/flavioesteves/wizer-app/broker/proto/profile"
	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context, sc pb.ProfileServiceClient) {
	profileId := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.GetProfile(ctx, &pb.GetProfileRequest{Id: profileId})
	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
	fmt.Printf("Profile was retrieved: %v\n", res)

	c.JSON(http.StatusOK, res)
}
