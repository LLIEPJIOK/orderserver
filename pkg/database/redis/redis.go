package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
	"gitlab.crja72.ru/gospec/students/223640-nphne-et6ofbhg-course-1195/internal/models"
	"gitlab.crja72.ru/gospec/students/223640-nphne-et6ofbhg-course-1195/pkg/database"
)

type RedisConfig struct {
	RedisHost string `env:"REDIS_HOST" env-default:"localhost"`
	RedisPort string `env:"REDIS_PORT" env-default:"6379"`
}

type Cache struct {
	cache *redis.Client
}

func New(cfg RedisConfig) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return &Cache{
		cache: client,
	}, nil
}

func (c *Cache) SetOrder(ctx context.Context, order *models.Order) error {
	data, err := json.Marshal(*order)
	if err != nil {
		return fmt.Errorf("failed to marshal order: %w", err)
	}

	key := orderIDKey(order.ID)

	res := c.cache.Set(ctx, key, string(data), 0)
	if err := res.Err(); err != nil {
		return fmt.Errorf("failed to set: %w", err)
	}

	return nil
}

func (c *Cache) GetOrder(ctx context.Context, id string) (*models.Order, error) {
	key := orderIDKey(id)

	val, err := c.cache.Get(ctx, key).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to get: %w", err)
	}

	var ord models.Order

	err = json.Unmarshal([]byte(val), &ord)
	if errors.Is(err, redis.Nil) {
		return nil, database.NewErrNotExists(id)
	} else if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	return &ord, nil
}

func (c *Cache) DeleteOrder(ctx context.Context, id string) error {
	key := orderIDKey(id)

	err := c.cache.Del(ctx, key).Err()
	if errors.Is(err, redis.Nil) {
		return database.NewErrNotExists(id)
	} else if err != nil {
		return fmt.Errorf("failed to delete: %w", err)
	}

	return nil
}

func (c *Cache) Close() error {
	if err := c.cache.Close(); err != nil {
		return fmt.Errorf("failed to close redis: %w", err)
	}

	return nil
}

func orderIDKey(id string) string {
	return fmt.Sprintf("order:%s", id)
}
