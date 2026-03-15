package base

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

// InitRedis 初始化 Redis 客户端
func InitRedis() error {
	opt := &redis.Options{
		Addr:     "localhost:6379",
		Password: "12345678",
		DB:       0,
	}
	cli := redis.NewClient(opt)
	// 做一次 ping 确认连接是否正常
	ctx := context.Background()
	if err := cli.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("无法连接到 Redis：%w", err)
	}
	client = cli // 赋值给包变量
	return nil
}

// Client 获取已初始化的 redis 客户端
// 如果没有初始化会 panic（或者你也可以返回 error 版本）
func Client() *redis.Client {
	if client == nil {
		panic("Redis 客户端尚未初始化，请先调用 InitRedis()")
	}
	return client
}
