package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type JWTPayload struct {
	Username string `json:"username"`
}

func getNameFromRefreshToken(refresh string) (string, error) {
	parts := strings.Split(refresh, ".")
	if len(parts) != 3 {
		return "", fmt.Errorf("invalid token: %s", refresh)
	}

	b, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("failed to decode token: %s", refresh)
	}

	var payload JWTPayload
	if err = json.Unmarshal(b, &payload); err != nil {
		return "", fmt.Errorf("failed to unmarshal token: %s", refresh)
	}

	return payload.Username, nil
}
