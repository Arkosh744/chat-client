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
	chatV1 "github.com/Arkosh744/chat-server/pkg/chat_v1"
)

func (h *Handler) ConnectChat(ctx context.Context, chatID string, user string) error {
	refresh, err := h.redisClient.Get(model.BuildRedisRefreshKey(user))
	if err != nil {
		return fmt.Errorf("failed to get refresh token: %w", err)
	}

	if len(refresh) < 1 {
		return fmt.Errorf("please re-login")
	}

	ctx = context.WithValue(ctx, model.UserNameKey, user)

	if _, err = h.chatClient.GetChat(ctx, chatID); err != nil {
		return err
	}

	stream, err := h.chatClient.ConnectToChat(ctx, chatID, user)
	if err != nil {
		return err
	}

	go getMessages(stream)

	return h.loopSendMessage(ctx, chatID, user)
}

func (h *Handler) loopSendMessage(ctx context.Context, chatID string, user string) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Errorf("failed to read from stdin: %s", err)
			continue
		}

		input = strings.Replace(input, "\n", "", -1)

		if err = h.chatClient.SendMessage(ctx, chatID, &model.Message{
			From:      user,
			Text:      input,
			CreatedAt: time.Now(),
		}); err != nil {
			log.Errorf("failed to send message: %s", err)

			return err
		}
	}
}

func getMessages(stream chatV1.ChatV1_ConnectToChatClient) {
	var retries int

	for {
		message, errRecv := stream.Recv()
		if errRecv != nil {
			if errRecv == io.EOF {
				return
			}

			retries++
			if retries > 5 {
				log.Errorf("close connect to the stream because too much fails: %s", errRecv)
				return
			}

			log.Errorf("failed to receive message from stream: %s", errRecv)
			continue
		}

		retries = 0

		log.Infof("[%v] %s: %s",
			message.GetCreatedAt().AsTime().Format("01.02.2006 15:04:05"),
			message.GetFrom(),
			message.GetText(),
		)
	}
}
