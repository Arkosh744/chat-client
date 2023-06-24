package redis

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

var _ Client = (*client)(nil)

type Client interface {
	Ping() error
	Get(key string) (string, error)
	Set(key string, value interface{}, expiration time.Duration) error
	GetAllUsers() ([]string, error)

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

// Get - if key does not found - return empty string without error
func (c *client) Get(key string) (string, error) {
	res, err := c.client.Get(key).Result()
	if err != nil {
		if errors.Is(redis.Nil, err) {
			return "", nil
		}

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
