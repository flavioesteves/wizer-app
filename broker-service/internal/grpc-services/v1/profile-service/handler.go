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

func CreateUser(c *gin.Context, sc pb.ProfileServiceClient) {
	var requestProfile models.RequestProfile

	if err := c.BindJSON(&requestProfile); err != nil {
		//TODO: Replace this logs with a proper log service
		fmt.Printf("Error to create a user in user_handler %v\n", err)
		return
	}

	newProfile := &pb.Profile{
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

	res, err := sc.CreateProfile(ctx, &pb.CreateProfileRequest{Profile: newProfile})
	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
	fmt.Printf("User was been created: %v\n", res)

	c.IndentedJSON(http.StatusOK, res)
}
