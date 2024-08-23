package v1

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	profileservice "github.com/flavioesteves/wizer-app/broker/internal/grpc-services/v1/profile-service"
	userservice "github.com/flavioesteves/wizer-app/broker/internal/grpc-services/v1/user-service"
)

func Routes() http.Handler {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"https://*", "http://*"},
		AllowOrigins:     []string{"http://localhost:5173"}, //TODO: replace for production domain
		AllowMethods:     []string{"GET", "PUT", "DELETE", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// TODO: remove just for dev
	router.GET("/healthcheck", GetAppStatus)

	// Profile Service
	profileRoutingGroup := router.Group("/v1/profiles")
	profileservice.RegisterProfileRoutes(profileRoutingGroup)

	// User Service
	userRoutingGroup := router.Group("v1/users")
	userservice.RegisterUserRoutes(userRoutingGroup)

	return router
}
