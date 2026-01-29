package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino/schema"
	"io"
	"log"
)

func main() {
	ctx := context.Background()
	chatModel, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		BaseURL: "http://localhost:11434", // Ollama 服务地址
		Model:   "qwen3:4b",               // 模型名称
	})
	if err != nil {
		log.Fatalf("创建 ChatModel 失败:%v\n", err)
	}
	// 准备消息
	message := []*schema.Message{
		schema.SystemMessage("你是一个助手"),
		schema.UserMessage("你好"),
	}
	stream, err := chatModel.Stream(ctx, message)
	if err != nil {
		log.Fatalf("获取 Stream 失败:%v\n", err)
	}
	defer stream.Close()
	// 处理 Stream 内容
	i := 0
	for {
		recv, err := stream.Recv()
		if err == io.EOF { // 流式输出结束
			return
		}
		if err != nil {
			log.Fatalf("recv 失败: %v\n", err)
		}
		fmt.Printf("message[%d]: %+v\n", i, recv.Content)
	}
}
