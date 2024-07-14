package main

import (
	"toko-kami/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.SetupRouter(app)

	app.Listen(":3000")
}
