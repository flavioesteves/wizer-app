package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"

	"github.com/flavioesteves/wizer-app/broker/api/v1"
	"github.com/flavioesteves/wizer-app/broker/internal/redis-service"
)

const WEB_PORT = "8080" //TODO: yml config file

type Config struct {
	RedisClient *redis.Client
	Handler     http.Handler
}

func main() {
	app := Config{
		RedisClient: redisservice.InitRedisServer(),
		Handler:     v1.Routes(),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", WEB_PORT),
		Handler: app.Handler,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
