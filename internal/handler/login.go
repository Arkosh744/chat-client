package handler

import (
	"context"

	"github.com/Arkosh744/chat-client/internal/model"
)

func (h *Handler) Login(ctx context.Context, info *model.AuthInfo) (string, error) {
	refreshToken, err := h.authClient.GetRefreshToken(ctx, info)
	if err != nil {
		return "", err
	}

	accessToken, err := h.authClient.GetAccessToken(ctx, refreshToken)
	if err != nil {
		return "", err
	}

	err = h.redisClient.Set(model.BuildRedisAccessKey(info.Username), accessToken, 0)
	if err != nil {
		return "", err
	}

	err = h.redisClient.Set(model.BuildRedisRefreshKey(info.Username), refreshToken, 0)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
