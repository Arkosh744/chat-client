package redis

import (
	"fmt"
	"strings"
	"time"

	"github.com/Arkosh744/chat-client/internal/model"
	"github.com/go-redis/redis"
)

var _ Client = (*client)(nil)

type Client interface {
	Ping() error
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) error
	GetAllUsers() ([]string, error)
	RefreshTokenExist(username, refresh string) (bool, error)

	Close() error
}

type client struct {
	client *redis.Client
}

func NewClient(addr string) *client {
	return &client{
		client: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
}

func (c *client) Ping() error {
	return c.client.Ping().Err()
}

func (c *client) Get(key string) (string, error) {
	res, err := c.client.Get(key).Result()
	if err != nil {
		return "", err
	}

	return res, nil
}

func (c *client) Set(key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(key, value, expiration).Err()
}

func (c *client) Close() error {
	return c.client.Close()
}

func (c *client) GetAllUsers() ([]string, error) {
	keys, err := c.client.Keys("user:*:refresh").Result()
	if err != nil {
		return nil, err
	}

	users := make([]string, len(keys))
	for i, key := range keys {
		parts := strings.Split(key, ":")
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid key: %s", key)
		}

		users[i] = parts[1]
	}

	return users, nil
}

func (c *client) RefreshTokenExist(username, refresh string) (bool, error) {
	refreshToken, err := c.Get(model.BuildRedisRefreshKey(username))
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}

		return false, err
	}

	return refreshToken == refresh, nil

}
