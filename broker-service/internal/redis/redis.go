package internal

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

const REDIS_PORT = "6379" //TODO: yml config file

// Create ane initialize redis client
func InitRedisServer() *redis.Client {

	redisOptions := redis.Options{
		Addr:     fmt.Sprintf("redis:%s", REDIS_PORT),
		Password: "", //TODO: yml config file
		DB:       0,
	}

	redisClient := redis.NewClient(&redisOptions)

	return redisClient
}
