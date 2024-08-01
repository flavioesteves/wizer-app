package event

import "github.com/redis/go-redis/v9"

type Consumer struct {
	redisClient *redis.Client
	queueName   string
}

func NewPublish()
