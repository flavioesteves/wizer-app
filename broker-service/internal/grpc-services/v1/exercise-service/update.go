package exerciseservice

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/flavioesteves/wizer-app/broker/api/v1/models"
	pbe "github.com/flavioesteves/wizer-app/broker/proto/exercise"
	pb "github.com/flavioesteves/wizer-app/broker/proto/model"
	"github.com/gin-gonic/gin"
)

func UpdateExercise(c *gin.Context, sc pbe.ExerciseServiceClient) {
	var requestExercise models.RequestExercise

	if err := c.BindJSON(&requestExercise); err != nil {
		fmt.Printf("Error to create a exercise in exercise_handler %v\n", err)
		return
	}
	updatedSteps := make([]*pb.Step, 0, len(requestExercise.Steps))

	for _, step := range requestExercise.Steps {
		updatedSteps = append(updatedSteps, &pb.Step{
			Description: step.Description,
			ImageUrl:    step.ImageUrl,
		})
	}

	updateExercise := &pb.Exercise{
		Id:                   requestExercise.Id,
		Name:                 requestExercise.Name,
		MuscleGroup:          requestExercise.MuscleGroup,
		Description:          requestExercise.Description,
		Steps:                updatedSteps,
		VideoUrl:             requestExercise.VideoUrl,
		VideoDurationSeconds: requestExercise.VideoDurationSeconds,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := sc.UpdateExercise(ctx, &pbe.UpdateExerciseRequest{Exercise: updateExercise})
	if err != nil {
		fmt.Printf("Unexpected error %v\n", err)
	}
	fmt.Printf("Exercise was been updated: %v\n", res)

	c.JSON(http.StatusOK, res)
}
