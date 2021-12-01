package services

import (
    "github.com/go-redis/redis"
    "sync"
    "workplace/internal/config"
)

var (
    client *redis.Client
    once   sync.Once
)

func GetRedisClient(configuration *config.Configuration) (rc *redis.Client, err error) {
    once.Do(func() {
        client = redis.NewClient(&redis.Options{
            Addr: configuration.RedisUrl,
            Password: "",
            DB: 0,
        })
        _, err = client.Ping().Result()
    })

    return client, err
}
