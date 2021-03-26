package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		c.SendFile("./public/html/index.html")
		return nil
	})

	app.Static("/js", "./public/js")
	app.Static("/css", "./public/css")
	app.Listen(":3000")
}
