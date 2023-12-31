package handler

import (
	"github.com/Arkosh744/chat-client/internal/client/grpc/auth"
	chatServer "github.com/Arkosh744/chat-client/internal/client/grpc/chat_server"
	"github.com/Arkosh744/chat-client/internal/client/redis"
)

type Handler struct {
	redisClient redis.Client
	authClient  auth.Client
	chatClient  chatServer.Client
}

func NewHandler(redisClient redis.Client, authClient auth.Client, chatClient chatServer.Client) *Handler {
	return &Handler{
		redisClient: redisClient,
		authClient:  authClient,
		chatClient:  chatClient,
	}
}
