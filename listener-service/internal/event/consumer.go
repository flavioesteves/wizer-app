package internal

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type Consumer struct {
	redisClient *redis.Client
	queueName   string
}

// TODO: Replace by gRPC
type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func NewConsumer(redisClient *redis.Client, queueName string) (Consumer, error) {
	consumer := Consumer{
		redisClient: redisClient,
		queueName:   queueName,
	}
	return consumer, nil
}

func (consumer *Consumer) Subscribe() error {
	var ctx = context.Background()
	subscriber := consumer.redisClient.Subscribe(ctx, consumer.queueName)

	payload := Payload{}

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			return err
		}
		if err := json.Unmarshal([]byte(msg.Payload), &payload); err != nil {
			return err
		}
	}
}
