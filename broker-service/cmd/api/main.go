package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

const WEB_PORT = "8080"   //TODO: yml config file
const REDIS_PORT = "6379" //TODO: yml config file

type Config struct {
	RedisClient *redis.Client
}

func main() {

	app := Config{
		RedisClient: initRedisServer(),
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", WEB_PORT),
		// Handler: app.routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}

}

func initRedisServer() *redis.Client {

	redisOptions := redis.Options{
		Addr:     fmt.Sprintf("redis:%s", REDIS_PORT),
		Password: "", //TODO: yml config file
		DB:       0,
	}

	redisClient := redis.NewClient(&redisOptions)

	return redisClient
}
