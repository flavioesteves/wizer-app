package profileservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	pb "github.com/flavioesteves/wizer-app/broker/proto"
)

func DeleteProfile(c *gin.Context, sc pb.ProfileServiceClient) {
	profileId := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.DeleteProfile(ctx, &pb.DeleteProfileRequest{Id: profileId})

	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
	fmt.Printf("Profile was deleted with id: %s", profileId)

	c.JSON(http.StatusOK, res)
}
