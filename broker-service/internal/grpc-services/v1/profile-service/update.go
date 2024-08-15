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

func UpdateProfile(c *gin.Context, sc pb.ProfileServiceClient) {
	var requestProfile models.RequestProfile

	if err := c.BindJSON(&requestProfile); err != nil {
		//TODO: Replace those logs with a proper log service
		fmt.Printf("Error to create a profile in profile_handler %v\n", err)
		return
	}

	updateProfile := &pb.Profile{
		ProfileId:         requestProfile.ProfileID,
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

	res, err := sc.UpdateProfile(ctx, &pb.UpdateProfileRequest{Profile: updateProfile})
	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
	fmt.Printf("Profile was been updated: %v\n", res)

	c.JSON(http.StatusOK, res)
}
