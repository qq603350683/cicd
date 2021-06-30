package main

import "github.com/go-redis/redis"

var RedisClient *redis.Client

func ConnectRedis() {
	// go-redis 配置所有参数详细说明 https://blog.csdn.net/pengpengzhou/article/details/105385666
	var opts redis.Options

	opts.PoolSize = 10
	opts.Addr = "127.0.0.1:6379"
	opts.Password = ""
	opts.DB = 0

	RedisClient = redis.NewClient(&opts)
}
