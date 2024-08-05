package main

import (
	event "github.com/flavioesteves/wizer-app/listener/internal/event"
	redis "github.com/flavioesteves/wizer-app/listener/internal/redis"
)

const REDIS_PORT = "6379"

type RMessage struct {
	Action string `json:"action"`
	Data   string `json:"data"`
}

func main() {

	redisClient := redis.InitRedisServer()

	subscriber, err := event.NewConsumer(redisClient, "test-JSON-data")
	if err != nil {
		panic(err)
	}

	err = subscriber.Subscribe()
	if err != nil {
		panic(err)
	}

}
