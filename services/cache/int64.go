package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/lily-lee/blog-server/config"
)

func GetInt64(key string, fetcher func() (int64, error)) (int64, error) {
	rc := config.RC
	val, err := rc.Get(context.Background(), key).Int64()
	if err != nil && err != redis.Nil {
		return 0, err
	}

	if err == nil {
		return val, nil
	}

	val, err = fetcher()
	if err != nil {
		return 0, err
	}

	_, err = rc.Set(context.Background(), key, val, time.Hour*24*20).Result()

	return val, err

}
