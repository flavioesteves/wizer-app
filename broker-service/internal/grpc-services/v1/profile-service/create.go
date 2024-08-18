package profileservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/flavioesteves/wizer-app/broker/api/v1/models"
	pb "github.com/flavioesteves/wizer-app/broker/proto"
)

func CreateProfile(c *gin.Context, sc pb.ProfileServiceClient) {
	var requestProfile models.RequestProfile

	if err := c.BindJSON(&requestProfile); err != nil {
		//TODO: Replace those logs with a proper log service
		fmt.Printf("Error to create a profile in profile_handler %v\n", err)
		return
	}

	newProfile := &pb.Profile{
		ProfileId:         requestProfile.Id,
		UserId:            requestProfile.UserID,
		Gender:            requestProfile.Gender,
		BirthYear:         requestProfile.BirthYear,
		HeightCm:          requestProfile.HeightCm,
		WeightKg:          requestProfile.WeightKg,
		BodyFatPercentage: requestProfile.BodyFatPercentage,
		Goal:              requestProfile.Goal,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.CreateProfile(ctx, &pb.CreateProfileRequest{Profile: newProfile})
	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
	fmt.Printf("Profile was been created: %v\n", res)

	c.JSON(http.StatusOK, res)
}
