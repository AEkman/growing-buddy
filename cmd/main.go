package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

var version = "dev"

func main() {
	fmt.Printf("Version: %s\n", version)

	fmt.Println(printApplicationStarted())

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Growing Buddy says hello!")
	})

	app.Listen(":8333")
}

func printApplicationStarted() string {
	return "Growing Buddy application started!"
}
