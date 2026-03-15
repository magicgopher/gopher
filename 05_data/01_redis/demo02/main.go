package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

// ctx 上下文
var ctx = context.Background()

// redis client
var rdb *redis.Client

func main() {
	// 初始化redis客户端
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "12345678",
		DB:       0,
	})
	// 关闭redis连接
	defer func() {
		if closeErr := rdb.Close(); closeErr != nil {
			log.Printf("redis连接关闭失败: %v\n", closeErr)
		}
		log.Printf("redis连接关闭成功!")
	}()
	// 清空所有数据库的所有数据
	flushDB()
}

// flushDB 清理所有数据库的数据
func flushDB() {
	if err := rdb.FlushDB(ctx).Err(); err != nil {
		log.Fatalf("flushdb执行失败: %v\n", err)
	}
	log.Printf("flushdb执行成功，清理所有数据库的数据。")
}
