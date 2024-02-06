package lib

import (
	"context"
	"errors"
	"fmt"
	"github.com/agung96tm/miblog/constants"
	appErrors "github.com/agung96tm/miblog/errors"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
	"time"
)

type IRedis interface {
	Set(key string, value any, expiration time.Duration) error
	Get(key string, value any) error
	Delete(keys ...string) (bool, error)
	Check(keys ...string) (bool, error)
}

type Redis struct {
	cache  *cache.Cache
	client *redis.Client
	prefix string
}

func NewRedis(config Config) Redis {
	addr := config.Redis.Addr()

	client := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   constants.RedisMainDB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatalf("Error to open redis[%s] connection: %v", addr, err)
	}

	//log.Info("Redis connection established")
	return Redis{
		client: client,
		prefix: config.Redis.KeyPrefix,
		cache: cache.New(&cache.Options{
			Redis:      client,
			LocalCache: cache.NewTinyLFU(1000, time.Minute),
		}),
	}
}

func (a Redis) wrapperKey(key string) string {
	return fmt.Sprintf("%s:%s", a.prefix, key)
}

func (a Redis) Set(key string, value any, expiration time.Duration) error {
	return a.cache.Set(&cache.Item{
		Ctx:            context.TODO(),
		Key:            a.wrapperKey(key),
		Value:          value,
		TTL:            expiration,
		SkipLocalCache: true,
	})
}

func (a Redis) Get(key string, value any) error {
	err := a.cache.Get(context.TODO(), a.wrapperKey(key), value)
	if errors.Is(err, cache.ErrCacheMiss) {
		err = appErrors.ErrRedisKeyNoExist
	}

	return err
}

func (a Redis) Delete(keys ...string) (bool, error) {
	wrapperKeys := make([]string, len(keys))
	for index, key := range keys {
		wrapperKeys[index] = a.wrapperKey(key)
	}

	cmd := a.client.Del(context.TODO(), wrapperKeys...)
	if err := cmd.Err(); err != nil {
		return false, err
	}

	return cmd.Val() > 0, nil
}

func (a Redis) Check(keys ...string) (bool, error) {
	wrapperKeys := make([]string, len(keys))
	for index, key := range keys {
		wrapperKeys[index] = a.wrapperKey(key)
	}

	cmd := a.client.Exists(context.TODO(), wrapperKeys...)
	if err := cmd.Err(); err != nil {
		return false, err
	}
	return cmd.Val() > 0, nil
}

func (a Redis) Close() error {
	return a.client.Close()
}

func (a Redis) GetClient() *redis.Client {
	return a.client
}
