package handler

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/Arkosh744/chat-client/internal/log"
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
		return fmt.Errorf("refresh token does not exist. please re-login and get new")
	}

	ctx = context.WithValue(ctx, model.UserNameKey, username)

	if _, err = h.chatClient.GetChat(ctx, chatID); err != nil {
		return err
	}

	stream, err := h.chatClient.ConnectToChat(ctx, chatID, username)
	if err != nil {
		return err
	}

	go func() {
		retries := 0

		for {
			message, errRecv := stream.Recv()
			if errRecv != nil {
				if errRecv == io.EOF {
					return
				}

				retries++
				if retries > 5 {
					log.Errorf("Close connect to the stream because too much fails: %s", errRecv)
					return
				}

				log.Errorf("failed to receive message from stream: %s", errRecv)
				continue
			}

			retries = 0

			log.Infof("[%v] %s: %s",
				message.GetCreatedAt().AsTime().Format("2006-01-02 15:04:05"),
				message.GetFrom(),
				message.GetText(),
			)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Errorf("failed to read from stdin: %s", err)
			continue
		}

		input = strings.Replace(input, "\n", "", -1)

		if err = h.chatClient.SendMessage(ctx, chatID, &model.Message{
			From:      username,
			Text:      input,
			CreatedAt: time.Now(),
		}); err != nil {
			log.Errorf("failed to send message: %s", err)

			return err
		}
	}
}
