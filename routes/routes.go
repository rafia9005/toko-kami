package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Get("/")
}

func AutoMigrate() {
}
