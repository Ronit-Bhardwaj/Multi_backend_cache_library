package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
	mutex  sync.Mutex
	ctx    context.Context
}

func NewRedisCache() *RedisCache {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &RedisCache{
		client: client,
		ctx:    context.Background(),
	}
}

func (r *RedisCache) Set(key string, value interface{}, ttl time.Duration) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	jsonData, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal value for key %s: %w", key, err)
	}

	if err := r.client.Set(r.ctx, key, jsonData, ttl).Err(); err != nil {
		return fmt.Errorf("failed to set key %s in Redis: %w", key, err)
	}

	return nil
}

func (r *RedisCache) Get(key string) (interface{}, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	val, err := r.client.Get(r.ctx, key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("key %s not found in cache", key)
	} else if err != nil {
		return nil, fmt.Errorf("failed to get key %s from Redis: %w", key, err)
	}

	var result interface{}
	if err := json.Unmarshal([]byte(val), &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal value for key %s: %w", key, err)
	}

	return result, nil
}

func (r *RedisCache) GetAllKeys() (map[string]interface{}, error) {
    r.mutex.Lock()
    defer r.mutex.Unlock()
   
    keys, err := r.client.Keys(r.ctx, "*").Result()
    if err != nil {
        return nil, fmt.Errorf("failed to get keys from Redis: %w", err)
    }

    items := make(map[string]interface{})
    for _, key := range keys {
        val, err := r.client.Get(r.ctx, key).Result()
        if err != nil {
            return nil, fmt.Errorf("failed to get value for key %s: %w", key, err)
        }
        
        var result interface{}
        if err := json.Unmarshal([]byte(val), &result); err != nil {
            return nil, fmt.Errorf("failed to unmarshal value for key %s: %w", key, err)
        }
        items[key] = result
    }

    return items, nil
}

func (r *RedisCache) Delete(key string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if err := r.client.Del(r.ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to delete key %s from Redis: %w", key, err)
	}

	return nil
}

func (r *RedisCache) Clear() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if err := r.client.FlushDB(r.ctx).Err(); err != nil {
		return fmt.Errorf("failed to clear Redis cache: %w", err)
	}

	return nil
}
