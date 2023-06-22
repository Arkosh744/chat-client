package handler

import (
	"context"
	"fmt"

	"github.com/Arkosh744/chat-client/internal/model"
)

func (h *Handler) AddToChat(ctx context.Context, chatID string, user string, addUser string) error {
	refresh, err := h.redisClient.Get(model.BuildRedisRefreshKey(user))
	if err != nil {
		return fmt.Errorf("failed to get refresh token: %w", err)
	}

	if len(refresh) < 1 {
		return fmt.Errorf("please re-login")
	}

	ctx = context.WithValue(ctx, model.UserNameKey, user)
	if err = h.chatClient.AddUserToChat(ctx, chatID, addUser); err != nil {
		return err
	}

	return nil
}
