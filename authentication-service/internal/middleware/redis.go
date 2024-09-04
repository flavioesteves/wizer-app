package middleware

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
var REDIS_DB = os.Getenv("REDIS_DB")
var REDIS_ADDRESS = os.Getenv("REDIS_ADDRESS")

func RedisConn() (*redis.Client, error) {
	redisDB, err := strconv.Atoi(REDIS_DB)
	if err != nil {
		fmt.Println("Error converting redisDB to integer:", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     REDIS_ADDRESS,
		Password: REDIS_PASSWORD,
		DB:       redisDB,
	})

	status := redisClient.Ping(context.Background())
	fmt.Println(status)

	return redisClient, nil
}
