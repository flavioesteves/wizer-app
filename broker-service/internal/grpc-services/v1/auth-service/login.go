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

func LoginUser(c *gin.Context, sc pb.AuthenticationServiceClient) {
	var reqLogin models.RequestAuthLogin

	fmt.Printf("sc: %v\n", sc)

	if err := c.BindJSON(&reqLogin); err != nil {
		fmt.Printf("Error to login in login_handler %v\n", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := sc.Login(ctx, &pb.LoginRequest{Email: reqLogin.Email, Password: reqLogin.Password})
	if err != nil {
		fmt.Printf("Login Unexpected error %v\n", err)
		c.JSON(http.StatusOK, gin.H{"isValid": false, "token": ""})
	} else {

		fmt.Printf("Login: %v\n", res)

		c.JSON(http.StatusOK, res)
	}
}
