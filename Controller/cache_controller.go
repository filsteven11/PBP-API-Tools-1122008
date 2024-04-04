package Controller

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type CacheController struct {
	client *redis.Client
}

func NewCacheController() *CacheController {
	return &CacheController{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}),
	}
}

func (cc *CacheController) Set(key string, value interface{}) error {
	ctx := context.Background()
	err := cc.client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (cc *CacheController) Get(key string) (string, error) {
	ctx := context.Background()
	val, err := cc.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
