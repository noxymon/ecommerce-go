package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(
		logger.New(),
		recover.New(),
		compress.New(compress.Config{
			Level: compress.LevelBestCompression,
		}),
	)
	app.Static("/assets", "view/assets")
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("view/index.html", fiber.Map{})
	})
	app.Get("/product", func(c *fiber.Ctx) error {
		return c.Render("view/product/index.html", fiber.Map{})
	})
	app.Get("/checkout", func(c *fiber.Ctx) error {
		return c.Render("view/checkout/index.html", fiber.Map{})
	})

	log.Fatal(app.Listen(":8080"))
}
