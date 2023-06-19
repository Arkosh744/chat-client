package handler

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Arkosh744/chat-client/internal/model"
)

func (h *Handler) ConnectChat(ctx context.Context, chatID string, refreshToken string) error {
	username, err := getNameFromRefreshToken(refreshToken)
	if err != nil {
		return fmt.Errorf("failed to handle refresh token: %w", err)
	}

	exist, err := h.redisClient.RefreshTokenExist(username, refreshToken)
	if err != nil {
		return fmt.Errorf("failed to check if refresh token exists: %w", err)
	}

	if !exist {
		return fmt.Errorf("refresh token does not exist. please re-login")
	}

	ctx = context.WithValue(ctx, model.UserNameKey, username)
	stream, err := h.chatClient.ConnectToChat(ctx, chatID, username)
	if err != nil {
		return err
	}

	go func() {
		for {
			message, errRecv := stream.Recv()
			if errRecv != nil {
				if errRecv == io.EOF {
					return
				}

				log.Println("failed to receive message from stream: ", errRecv)
				return
			}

			log.Printf("[%v] - [from: %s]: %s\n", message.GetCreatedAt(), message.GetFrom(), message.GetText())
		}
	}()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		var lines strings.Builder

		for {
			scanner.Scan()
			line := scanner.Text()
			if len(line) == 0 {
				break
			}

			lines.WriteString(line)
			lines.WriteString("\n")
		}

		if err = scanner.Err(); err != nil {
			log.Println("failed to scan message: ", err)
		}

		if err = h.chatClient.SendMessage(ctx, chatID, &model.Message{
			From:      username,
			Text:      lines.String(),
			CreatedAt: time.Now(),
		}); err != nil {
			log.Println("failed to send message: ", err)
			return err
		}
	}
}
