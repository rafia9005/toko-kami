package main

import (
	"toko-kami/database"
	"toko-kami/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	godotenv.Load()
	database.Connect()
	routes.AutoMigrate()
	routes.SetupRouter(app)
	app.Listen(":3000")
}
