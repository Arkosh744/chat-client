package handler

import (
	"context"
	"fmt"

	"github.com/Arkosh744/chat-client/internal/model"
)

func (h *Handler) CreateChat(ctx context.Context, user string, usernames []string, withHistory bool) (string, error) {
	refresh, err := h.redisClient.Get(model.BuildRedisRefreshKey(user))
	if err != nil {
		return "", fmt.Errorf("failed to get refresh token: %w", err)
	}

	if len(refresh) < 1 {
		return "", fmt.Errorf("please re-login")
	}

	if !containsString(usernames, user) {
		usernames = append(usernames, user)
	}

	ctx = context.WithValue(ctx, model.UserNameKey, user)
	chatID, err := h.chatClient.CreateChat(ctx, usernames, withHistory)
	if err != nil {
		return "", err
	}

	return chatID, nil
}

func containsString(slice []string, res string) bool {
	for _, cur := range slice {
		if cur == res {
			return true
		}
	}

	return false
}
