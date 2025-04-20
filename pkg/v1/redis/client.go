package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
)

type RedisStore struct {
	Client  *redis.Client
	Limiter *redis_rate.Limiter
}

func NewRedisClient(addr, password string, db int) *RedisStore {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

	return &RedisStore{
		Client:  client,
		Limiter: redis_rate.NewLimiter(client),
	}
}

func (r *RedisStore) CheckHealth(ctx context.Context) error {
	_, err := r.Client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("redis is not ready: %v", err)
	}
	return nil
}

func (r *RedisStore) Allow(ctx context.Context, key string, limit redis_rate.Limit) (*redis_rate.Result, error) {
	return r.Limiter.Allow(ctx, key, limit)
}

func (r *RedisStore) Close() error {
	return r.Client.Close()
}
