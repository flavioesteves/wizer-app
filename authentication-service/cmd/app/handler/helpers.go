package handler

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	pb "github.com/flavioesteves/wizer-app/authentication/proto"
	"github.com/redis/go-redis/v9"
)

type ServerConfig struct {
	redisClient *redis.Client
	db          *sql.DB
	pb.AuthenticationServiceServer
}

func NewServerConfig(redisClient *redis.Client, db *sql.DB) *ServerConfig {
	return &ServerConfig{redisClient: redisClient, db: db}
}

func DecryptPassword(hashedPassword string, password string) error {

	fmt.Printf("hashedPassword: %v\n", hashedPassword)
	fmt.Printf("Password: %v\n", password)

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	fmt.Printf("Hashed error: %v\n", err)
	return err
}
