package event

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Emitter struct {
	redisClient *redis.Client
	queueName   string
}

// TODO: replace by gRPC
type RMessage struct {
	Action string `json:"action"`
	Data   string `json:"data"`
}

func NewEventEmitter() {}

func NewPayLoad(c *gin.Context, redisCLient *redis.Client) {
	message := new(RMessage)

	if err := c.ShouldBindJSON(message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	payload, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	fmt.Println(payload)
}
