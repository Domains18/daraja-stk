package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("welcome to the api")
}

func main() {
	app := fiber.New()
	app.Get("/api", welcome)
	log.Fatal(app.Listen(":3000"))
}
