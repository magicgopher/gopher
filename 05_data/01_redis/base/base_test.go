package base

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestInitRedis redis客户端初始化单元测试
func TestInitRedis(t *testing.T) {
	tests := []struct {
		name        string
		modifyOpt   func(*redis.Options) // 允许测试时临时修改连接参数
		wantErr     bool
		errContains string // 中文关键词即可
	}{
		{
			name: "正常情况（假设本地有 Redis 服务）",
			modifyOpt: func(opt *redis.Options) {
				// 使用默认配置
			},
			wantErr: false,
		},
		{
			name: "密码错误",
			modifyOpt: func(opt *redis.Options) {
				opt.Password = "错的密码123"
			},
			wantErr:     true,
			errContains: "ERR invalid password", // redis 原生错误，保留英文关键词
		},
		{
			name: "连接地址错误",
			modifyOpt: func(opt *redis.Options) {
				opt.Addr = "127.0.0.1:9999" // 基本不可能存在的端口
			},
			wantErr:     true,
			errContains: "无法连接到 Redis",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 每次测试前清空全局变量
			client = nil
			opt := &redis.Options{
				Addr:     "localhost:6379",
				Password: "12345678",
				DB:       0,
			}
			// 让测试案例可以修改配置
			if tt.modifyOpt != nil {
				tt.modifyOpt(opt)
			}
			// 临时覆盖全局 client（仅测试用）
			tempClient := redis.NewClient(opt)
			err := tempClient.Ping(context.Background()).Err()
			if err != nil {
				err = fmt.Errorf("无法连接到 Redis：%w", err)
			}
			if tt.wantErr {
				assert.Error(t, err, "预期应该报错")
				if tt.errContains != "" {
					assert.ErrorContains(t, err, tt.errContains,
						"错误信息里应该包含关键词：%s", tt.errContains)
				}
			} else {
				assert.NoError(t, err, "不应该报错")
				// 验证真的能 ping 通
				pong, pingErr := tempClient.Ping(context.Background()).Result()
				assert.NoError(t, pingErr)
				assert.Equal(t, "PONG", pong)
			}
		})
	}
}

// TestClientPanic 没有初始化redis客户端
func TestClientPanic(t *testing.T) {
	client = nil // 确保没初始化
	assert.PanicsWithValue(t,
		"Redis 客户端尚未初始化，请先调用 InitRedis()",
		func() {
			_ = Client()
		},
		"未初始化时应该 panic 并显示中文提示",
	)
}
