package v1

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/flavioesteves/wizer-app/broker/internal/grpc-services/v1/user-service"
)

func Routes() http.Handler {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "PUT", "DELETE", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// TODO: remove just for dev
	router.GET("/healthcheck", GetAppStatus)

	// User Service
	userRoutingGroup := router.Group("/v1/users")
	userservice.RegisterUserRoutes(userRoutingGroup)

	return router
}
