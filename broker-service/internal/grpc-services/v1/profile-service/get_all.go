package profileservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	pb "github.com/flavioesteves/wizer-app/broker/proto"
)

func GetAllProfiles(c *gin.Context, sc pb.ProfileServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.GetAllProfiles(ctx, &pb.GetAllProfilesRequest{})
	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
	fmt.Printf("All Profiles were retrieved: %v\n", res)

	c.JSON(http.StatusOK, res)
}
