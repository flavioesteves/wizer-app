package handler

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	pb "github.com/flavioesteves/wizer-app/authentication/proto"
	"github.com/redis/go-redis/v9"
)

type ServerConfig struct {
	redisClient *redis.Client
	db          *sql.DB
	pb.AuthenticationServiceServer
}

type Claims struct {
	Username string `json:"email"`
	jwt.StandardClaims
}

func NewServerConfig(redisClient *redis.Client, db *sql.DB) *ServerConfig {
	return &ServerConfig{redisClient: redisClient, db: db}
}

func DecryptPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

func GenerateToken(username string) (string, error) {
	claimsKey := os.Getenv("CLAIMS_KEY")
	claims := Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "Wizer",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(claimsKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (s *ServerConfig) GetSessionData(ctx context.Context, token string) (string, error) {
	sessionData, err := s.redisClient.Get(ctx, token).Result()
	if err != nil {
		return "", err
	}

	fmt.Println(sessionData)
	return sessionData, nil
}
