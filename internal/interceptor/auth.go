package interceptor

import (
	"context"
	"errors"
	"time"

	"github.com/Arkosh744/chat-client/internal/client/grpc/auth"
	"github.com/Arkosh744/chat-client/internal/client/redis"
	"github.com/Arkosh744/chat-client/internal/log"
	"github.com/Arkosh744/chat-client/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthInterceptor struct {
	authClient  auth.Client
	redisClient redis.Client
}

func NewAuthInterceptor(authClient auth.Client, redisClient redis.Client) *AuthInterceptor {
	return &AuthInterceptor{
		authClient:  authClient,
		redisClient: redisClient,
	}
}

// Run - refreshing all logged-in users' access tokens periodically until refresh token is valid
func (i *AuthInterceptor) Run(accessTokenPeriod time.Duration) {
	go func() {
		t := time.NewTicker(accessTokenPeriod)
		ctx := context.Background()

		for range t.C {
			users, err := i.redisClient.GetAllUsers()
			if err != nil {
				log.Errorf("failed to get users from redis: %v", err)
				continue
			}

			for _, user := range users {
				refreshTokenKey := model.BuildRedisRefreshKey(user)
				accessTokenKey := model.BuildRedisAccessKey(user)

				refreshToken, err := i.redisClient.Get(refreshTokenKey)
				if err != nil {
					log.Errorf("failed to get refresh token for user %s from redis: %v", user, err)
					continue
				}

				accessToken, err := i.authClient.GetAccessToken(ctx, refreshToken)
				if err != nil {
					log.Errorf("failed to get access token for user %s: %v", user, err)
					continue
				}

				if err = i.redisClient.Set(accessTokenKey, accessToken, 0); err != nil {
					log.Errorf("failed to set access token for user %s to redis: %v", user, err)
					continue
				}

				log.Infof("access token for user %s has been updated", user)
			}
		}
	}()
}

func (i *AuthInterceptor) Unary(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	user, ok := ctx.Value(model.UserNameKey).(string)
	if !ok {
		return errors.New("failed to get username from context")
	}

	accessTokenKey := model.BuildRedisAccessKey(user)

	accessToken, err := i.redisClient.Get(accessTokenKey)
	if err != nil {
		return err
	}

	md := metadata.New(map[string]string{"Authorization": "Bearer " + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	return invoker(ctx, method, req, reply, cc, opts...)
}
