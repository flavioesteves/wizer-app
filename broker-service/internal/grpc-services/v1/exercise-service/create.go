package exerciseservice

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/flavioesteves/wizer-app/broker/api/v1/models"
	pbe "github.com/flavioesteves/wizer-app/broker/proto/exercise"
	pb "github.com/flavioesteves/wizer-app/broker/proto/model"
)

func CreateExercise(c *gin.Context, sc pbe.ExerciseServiceClient) {
	var requestExercise models.RequestExercise

	if err := c.BindJSON(&requestExercise); err != nil {
		fmt.Printf("Error to create a exercise in exercise_handler %v\n", err)
		return
	}

	newSteps := make([]*pb.Step, 0, len(requestExercise.Steps))

	for _, step := range requestExercise.Steps {
		newSteps = append(newSteps, &pb.Step{
			Description: step.Description,
			ImageUrl:    step.ImageUrl,
		})
	}

	newExercise := &pb.Exercise{
		Id:                   requestExercise.Id,
		Name:                 requestExercise.Name,
		MuscleGroup:          requestExercise.MuscleGroup,
		Description:          requestExercise.Description,
		Steps:                newSteps,
		VideoUrl:             requestExercise.VideoUrl,
		VideoDurationSeconds: requestExercise.VideoDurationSeconds,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.CreateExercise(ctx, &pbe.CreateExerciseRequest{Exercise: newExercise})
	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
	fmt.Printf("Exercise was been created: %v\n", res)

	c.JSON(http.StatusOK, res)
}
