package config

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var RC *redis.Client

func InitRedis() {
	opt := &redis.Options{}
	opt.Addr = Conf.RedisAddr
	opt.PoolSize = Conf.RedisPoolSize
	RC = redis.NewClient(opt)
	if RC == nil {
		panic("init redis failed.")
	}
	fmt.Println(RC.Ping(context.Background()))
}
