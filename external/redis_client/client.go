package redisclient

import (
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

type Client struct {
	RedisClient *cache.Cache
}

func InitRedisCacheClient(config RedisConfig) *Client {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"localhost":      ":7379",
			config.REDISHost: config.REDISPort,
		},
	})

	mycache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(2000, time.Minute),
	})
	return &Client{
		RedisClient: mycache,
	}
}
