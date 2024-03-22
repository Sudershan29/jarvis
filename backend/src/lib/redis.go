package lib

import (
	"fmt"
	"time"
    "context"
	helper "backend/src/helpers"
    "github.com/go-redis/redis/v8"
)

type redisStore struct {
	client  *redis.Client
	context context.Context
}

func ConnectToRedis() *redisStore{
	redisURL := fmt.Sprintf("%s:%s", helper.GetEnv("REDIS_URL"), helper.GetEnv("REDIS_PORT"))
	rdb := redis.NewClient(&redis.Options{
			Addr:     redisURL,
			Password: "",
			DB:       0,
		})

	return &redisStore{rdb, context.Background()}
}

func (store *redisStore) Store(key string, value string) error {
	return store.StoreWithExpiry(key, value, 0)
}

func (store *redisStore) StoreWithExpiry(key string, value string, expiration time.Duration)  error {
	return store.client.Set(store.context, key, value, expiration).Err()
}


func (store *redisStore) Get(key string) (string, error) {
	val, err := store.client.Get(store.context, key).Result()
    if err != nil {
        return "", err
    }
	return val, nil
}

var RedisClient *redisStore

func init() {
	RedisClient = ConnectToRedis()
}