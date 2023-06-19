package chat_server

import (
	"context"

	"github.com/Arkosh744/chat-client/internal/model"
	chatV1 "github.com/Arkosh744/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ Client = (*client)(nil)

type Client interface {
	CreateChat(ctx context.Context, usernames []string, withHistory bool) (string, error)
	ConnectToChat(ctx context.Context, chatID string, username string) (chatV1.ChatV1_ConnectToChatClient, error)
	SendMessage(ctx context.Context, chatID string, message *model.Message) error
	AddUserToChat(ctx context.Context, chatID string, username string) error
}

type client struct {
	client chatV1.ChatV1Client
}

func NewClient(cc *grpc.ClientConn) *client {
	return &client{
		client: chatV1.NewChatV1Client(cc),
	}
}

func (c *client) CreateChat(ctx context.Context, usernames []string, withHistory bool) (string, error) {
	res, err := c.client.CreateChat(ctx, &chatV1.CreateChatRequest{
		Usernames:   usernames,
		SaveHistory: withHistory,
	})
	if err != nil {
		return "", err
	}

	return res.GetChatId(), nil
}

func (c *client) ConnectToChat(ctx context.Context, chatID string, username string) (chatV1.ChatV1_ConnectToChatClient, error) {
	return c.client.ConnectToChat(ctx, &chatV1.ConnectChatRequest{
		ChatId:   chatID,
		Username: username,
	})
}

func (c *client) SendMessage(ctx context.Context, chatID string, message *model.Message) error {
	_, err := c.client.SendMessage(ctx, &chatV1.SendMessageRequest{
		ChatId: chatID,
		Message: &chatV1.Message{
			From:      message.From,
			Text:      message.Text,
			CreatedAt: timestamppb.New(message.CreatedAt),
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *client) AddUserToChat(ctx context.Context, chatID string, username string) error {
	_, err := c.client.AddUserToChat(ctx, &chatV1.AddUserToChatRequest{
		ChatId:   chatID,
		Username: username,
	})
	if err != nil {
		return err
	}

	return nil
}
