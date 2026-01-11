package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var ctx = context.Background()

var rdb *redis.Client

func main() {
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

	// set操作
	//setOption("name", "张三", 30*time.Second)

	// get操作
	//getOption()

	// ttl操作
	//ttlOption("name")

	// getset操作
	//getSetOption("hello", "redis!")

	// setnx操作
	//setNxOption("hello", "你好", 0)
	//setNxOption("name", "张三", 0)

	// mset操作
	//m := map[string]interface{}{
	//	"k1": "v1",
	//	"k2": "v2",
	//	"k3": "v3",
	//	"k4": "v4",
	//	"k5": "v5",
	//}
	//mSetOption(m)

	// mget操作
	//keys := []string{
	//	"k1",
	//	"k2",
	//	"k3",
	//	"k4",
	//	"k5",
	//}
	//mGetOption(keys)

	// incr操作
	//incrOption("num")

	// incrby操作
	//incrByOption("num", 5)

	// incrbyfloat操作
	//incrByFloatOption("num", 1.8)

	// decr操作
	//decrOption("num")

	// decrby操作
	//decrByOption("num", 5)

	// del操作
	//delOption("k1", "k2", "k3")

	// expire操作
	//expireOption("name", 20*time.Second)
}

// setOption 操作
func setOption(key string, value any, expire time.Duration) {
	if setErr := rdb.Set(ctx, key, value, expire).Err(); setErr != nil {
		log.Printf("redis set操作失败: %v\n", setErr)
	}
	log.Printf("redis set操作成功, key=%v, value=%v\n", key, value)
}

// getOption get操作
func getOption(key string) {
	result, err := rdb.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			log.Printf("key [%s] 不存在\n", err)
			return
		}
		fmt.Printf("redis get操作失败: %v\n", err)
		return
	}
	fmt.Println("获取成功:", result)
}

// ttlOption TTL操作
func ttlOption(key string) {
	// 执行 TTL 命令
	result, err := rdb.TTL(ctx, key).Result()

	// 错误处理
	if err != nil {
		log.Printf("redis ttl 命令执行失败: %v", err)
		return
	}

	if result < 0 {
		if result == -2*time.Nanosecond || result == -2*time.Second {
			fmt.Printf("key: [%s] 不存在.\n", key)
		} else if result == -1*time.Nanosecond || result == -1*time.Second {
			fmt.Println("key 永久有效。")
		}
		return
	}
	fmt.Printf("key [%s] 剩余有效期: %v\n", key, result)
}

// getSetOption 设置一个key的值，并返回这个key的旧值
func getSetOption(key, newValue string) {
	oldValue, err := rdb.GetSet(ctx, key, newValue).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			log.Printf("key [%s] 之前不存在，已设置为新值: %s", key, newValue)
			fmt.Printf("key: [%v]; oldValue: [不存在/已新建]\n", key)
			return
		}
		log.Printf("执行 GetSet 失败: %v", err)
		return
	}
	// 打印旧值
	fmt.Printf("key: [%v]; oldValue:[%v]\n", key, oldValue)
}

func setNxOption(key string, value any, expire time.Duration) {
	// setnx 命令
	result, err := rdb.SetNX(ctx, key, value, expire).Result()
	if err != nil {
		log.Printf("redis setnx 操作失败: key=%s, err=%v", key, err)
	}
	// 根据 bool 结果判断业务逻辑
	if result {
		// 返回 true，表示抢占成功（通常用于分布式锁获取成功）
		fmt.Printf("成功设置 Key [%s]，过期时间: %v\n", key, expire)
	} else {
		// 返回 false，表示 Key 已存在，本次操作未执行任何更改
		fmt.Printf("设置 Key [%s] 失败: Key 已存在\n", key)
	}
}

func mSetOption(pairs map[string]interface{}) {
	status, err := rdb.MSet(ctx, pairs).Result()

	// 错误处理
	if err != nil {
		log.Printf("redis mset 批量设置失败: %v", err)
		return
	}

	// 成功处理
	fmt.Printf("批量设置成功: %s, 影响 Key 数量: %d\n", status, len(pairs))
}

func mGetOption(keys []string) {
	// 执行 mget 命令
	results, err := rdb.MGet(ctx, keys...).Result()
	if err != nil {
		log.Printf("redis mget 批量获取失败: %v", err)
		return
	}

	// 遍历结果
	fmt.Printf("批量获取结果 (共 %d 个):\n", len(results))
	for i, val := range results {
		key := keys[i]

		// 检查该 key 是否存在
		if val == nil {
			fmt.Printf("Key: [%s] -> 结果: [不存在/Nil]\n", key)
			continue
		}

		// 如果存在，val 是 interface{}，通常需要断言为 string
		// 或者直接打印
		fmt.Printf("Key: [%s] -> 结果: [%v]\n", key, val)
	}
}

func incrOption(key string) {
	// 先设置一个可以 Incr 的初始值
	// 我们用 SetNX 保证只有在 Key 不存在时才初始化为 0
	_, err := rdb.SetNX(ctx, key, "0", 0).Result()
	if err != nil {
		log.Printf("初始化 Key [%s] 失败: %v", key, err)
		return
	}

	// 执行自增操作
	// Incr 返回自增后的新值 (int64)
	newValue, err := rdb.Incr(ctx, key).Result()

	// 错误处理
	if err != nil {
		// 情况 A: Key 里的内容不是数字 (比如存的是 "abc")
		// Redis 会返回: ERR value is not an integer or out of range
		log.Printf("自增操作失败，Key [%s] 可能包含非数字内容: %v", key, err)
		return
	}

	// 打印结果
	fmt.Printf("Key [%s] 自增成功，当前值: %d\n", key, newValue)
}

func incrByOption(key string, increment int64) {
	// 直接执行 IncrBy，如果 key 不存在，Redis 会自动创建并从 0 开始加
	result, err := rdb.IncrBy(ctx, key, increment).Result()

	// 错误处理
	if err != nil {
		log.Printf("IncrBy 操作失败 [key: %s, 增量: %d]: %v", key, increment, err)
		return
	}

	// 打印结果
	fmt.Printf("Key [%s] 成功增加 %d，当前最新值为: %d\n", key, increment, result)
}

func incrByFloatOption(key string, increment float64) {
	// 直接执行 IncrByFloat
	// result 是增加后的最终结果 (float64)
	result, err := rdb.IncrByFloat(ctx, key, increment).Result()

	// 处理错误
	if err != nil {
		// 常见错误场景：
		// - Key 存储的内容不能转换为浮点数（例如存了非数字字符串）
		// - 结果超出了浮点数的表示范围
		log.Printf("IncrByFloat 操作失败 [key: %s, 增量: %f]: %v", key, increment, err)
		return
	}

	// 使用 %.2f 格式化输出，保留两位小数
	fmt.Printf("Key [%s] 成功增加 %.2f，当前最新值为: %.2f\n", key, increment, result)
}

func decrOption(key string) {
	result, err := rdb.Decr(ctx, key).Result()
	if err != nil {
		log.Printf("Decr 失败 [key: %s]: %v", key, err)
	}
	fmt.Printf("Decr 成功，当前值: %d\n", result)
}

func decrByOption(key string, decrement int64) {
	result, err := rdb.DecrBy(ctx, key, decrement).Result()
	if err != nil {
		log.Printf("DecrBy 失败 [key: %s]: %v", key, err)
	}
	fmt.Printf("DecrBy 成功，减去 %d 后当前值: %d\n", decrement, result)
}

func delOption(keys ...string) {
	// Del 可以接收多个 key，如 rdb.Del(ctx, "k1", "k2")
	count, err := rdb.Del(ctx, keys...).Result()

	if err != nil {
		log.Printf("Del 操作系统错误: %v", err)
		return
	}

	// count 表示真正被删除掉的 key 的个数
	if count == 0 {
		fmt.Println("没有 key 被删除（这些 key 本来就不存在）")
	} else {
		fmt.Printf("成功删除 %d 个 key\n", count)
	}
}

func expireOption(key string, expiration time.Duration) {
	success, err := rdb.Expire(ctx, key, expiration).Result()

	if err != nil {
		log.Printf("Expire 操作系统错误: %v", err)
		return
	}

	if success {
		fmt.Printf("Key [%s] 已成功设置过期时间: %v\n", key, expiration)
	} else {
		// 这里的 false 通常意味着 key 不存在
		fmt.Printf("Key [%s] 设置过期时间失败，可能该 key 不存在\n", key)
	}
}
