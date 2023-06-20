package handler

import (
	"context"
	"fmt"

	"github.com/Arkosh744/chat-client/internal/model"
)

func (h *Handler) AddToChat(ctx context.Context, chatID string, user string, refreshToken string) error {
	username, err := getNameFromRefreshToken(refreshToken)
	if err != nil {
		return fmt.Errorf("failed to handle refresh token: %w", err)
	}

	exist, err := h.redisClient.RefreshTokenExist(username, refreshToken)
	if err != nil {
		return fmt.Errorf("failed to check if refresh token exists: %w", err)
	}

	if !exist {
		return fmt.Errorf("refresh token does not exist. please re-login and get new")
	}

	ctx = context.WithValue(ctx, model.UserNameKey, username)
	if err = h.chatClient.AddUserToChat(ctx, chatID, user); err != nil {
		return err
	}

	return nil
}
