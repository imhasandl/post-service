package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisClient represents a Redis client wrapper
type RedisClient struct {
	Client *redis.Client
}

// NewRedisClient creates and returns a new Redis client
func NewRedisClient(url string) (*RedisClient, error) {
	opt, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)

	// Check connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &RedisClient{
		Client: client,
	}, nil
}

// Set stores a value with an optional expiration
func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(ctx, key, value, expiration).Err()
}

// Get retrieves a value by key
func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

// Delete removes a key from Redis
func (r *RedisClient) Delete(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}

// Close closes the Redis connection
func (r *RedisClient) Close() error {
	return r.Client.Close()
}
