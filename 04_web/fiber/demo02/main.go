package main

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func main() {
	// 创建一个app
	app := fiber.New()

	// 路由处理GET请求
	app.Get("/user/list", func(ctx fiber.Ctx) error {
		return ctx.SendString("get请求1")
	})

	// 路由处理GET请求URL路径携带参数
	app.Get("/user/:id", func(ctx fiber.Ctx) error {
		id := ctx.Params("id")
		fmt.Println(id)
		return ctx.SendString("get请求2")
	})

	// 启动服务并监听3000端口
	log.Fatal(app.Listen(":3000"))
}
