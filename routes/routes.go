package routes

import (
	"toko-kami/handler"
	"toko-kami/model/entity"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
}

func AutoMigrate() {
	RunMigrate(&entity.Users{})
}
