package userservice

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var USER_SERVICE_HOST = "user-service:50051"

func RegisterUserRoutes(rg *gin.RouterGroup) {

	sc := Connect(USER_SERVICE_HOST)

	fmt.Printf("Connect: %v", sc)

	rg.POST("", func(c *gin.Context) { CreateUser(c, sc) })
}
