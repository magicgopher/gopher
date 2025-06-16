package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

// 请求获取携带的参数示例

func main() {
	// 创建 app
	app := fiber.New()

	// GET http://localhost:8080/hello%20world
	app.Get("/:value", func(c fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
		// => Response: "value: hello world"
	})

	// GET http://localhost:3000/john
	app.Get("/:name?", func(c fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
			// => Response: "Hello john"
		}
		return c.SendString("Where is john?")
		// => Response: "Where is john?"
	})

	// GET http://localhost:3000/api/user/john
	app.Get("/api/*", func(c fiber.Ctx) error {
		return c.SendString("API path: " + c.Params("*"))
		// => Response: "API path: user/john"
	})

	log.Fatal(app.Listen(":3000"))
}
