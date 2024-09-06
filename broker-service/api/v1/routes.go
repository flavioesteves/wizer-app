package v1

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	authservice "github.com/flavioesteves/wizer-app/broker/internal/grpc-services/v1/auth-service"
	profileservice "github.com/flavioesteves/wizer-app/broker/internal/grpc-services/v1/profile-service"
	userservice "github.com/flavioesteves/wizer-app/broker/internal/grpc-services/v1/user-service"
)

func Routes() http.Handler {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		//AllowOrigins: []string{"https://*", "http://*"},
		AllowOrigins:     []string{"http://localhost:5173"}, //TODO: replace for production domain
		AllowMethods:     []string{"GET", "PUT", "DELETE", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Accept", "Content-Length", "Authorization", "Content-Type", "X-CSRF-Token", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Link", "Content-Length"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// TODO: remove just for dev
	router.GET("/healthcheck", GetAppStatus)

	//Auth Service
	authRoutingGroup := router.Group("v1/auth")
	authservice.RegisterAuthRoutes(authRoutingGroup)

	// User Service
	userRoutingGroup := router.Group("v1/users")
	userservice.RegisterUserRoutes(userRoutingGroup)

	// Profile Service
	profileRoutingGroup := router.Group("/v1/profiles")
	profileservice.RegisterProfileRoutes(profileRoutingGroup)

	return router
}
