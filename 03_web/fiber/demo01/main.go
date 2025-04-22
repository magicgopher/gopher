package main

import (
	"github.com/gofiber/fiber/v3"
	"log"
)

func main() {
	// fiber是一个基于fasthttp构建的go web框架

	// 初始化新的 Fiber 应用
	app := fiber.New()

	// 为 /hello 的URI路径上的GET方法定义路由
	app.Get("/hello", func(ctx fiber.Ctx) error {
		return ctx.SendString("Hello, Fiber!")
	})

	// 启动服务并监听3000端口
	log.Fatal(app.Listen(":3000"))
}
