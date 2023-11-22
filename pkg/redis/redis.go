package redis

import (
	"context"
	"encoding/json"
	"github.com/app-dictionary/pkg/env"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
	"time"
)

type CacheManager struct {
	client *redis.Client
}

var (
	once     sync.Once
	instance *redis.Client
)

func SetupRedisClient() *redis.Client {
	once.Do(func() {
		instance = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: env.GetEnv("REDIS_PASSWORD", ""),
			DB:       0,
		})

		// Проверка соединения с Redis
		_, err := instance.Ping(context.Background()).Result()
		if err != nil {
			log.Printf("Redis not init - %v", err)
		}
	})

	return instance
}

func NewCacheManager() *CacheManager {
	return &CacheManager{
		client: SetupRedisClient(),
	}
}

func (cm *CacheManager) Get(key string, target interface{}) error {
	data, err := cm.client.Get(context.Background(), key).Result()
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(data), target); err != nil {
		return err
	}

	return nil
}

func (cm *CacheManager) Set(key string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := cm.client.Set(context.Background(), key, jsonData, 24*time.Hour).Err(); err != nil {
		return err
	}

	return nil
}
