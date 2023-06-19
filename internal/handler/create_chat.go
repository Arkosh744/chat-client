package handler

import (
	"context"
	"fmt"

	"github.com/Arkosh744/chat-client/internal/model"
)

func (h *Handler) CreateChat(ctx context.Context, usernames []string, refreshToken string) (string, error) {
	username, err := getNameFromRefreshToken(refreshToken)
	if err != nil {
		return "", fmt.Errorf("failed to handle refresh token: %w", err)
	}

	exist, err := h.redisClient.RefreshTokenExist(username, refreshToken)
	if err != nil {
		return "", fmt.Errorf("failed to check if refresh token exists: %w", err)
	}

	if !exist {
		return "", fmt.Errorf("refresh token does not exist. please re-login")
	}

	if !containsString(usernames, username) {
		usernames = append(usernames, username)
	}

	ctx = context.WithValue(ctx, model.UserNameKey, username)
	chatID, err := h.chatClient.CreateChat(ctx, usernames)
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
