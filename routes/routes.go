package routes

import (
	"toko-kami/handler"
	"toko-kami/model/entity"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// Product
	app.Get("/product", handler.GetProduct)
	app.Get("/product/:id", handler.GetByIdProduct)
	app.Delete("/product/:id", handler.DeleteProduct)

	// static folder
	app.Static("/", "./public")
}

func AutoMigrate() {
	RunMigrate(&entity.Users{})
	RunMigrate(&entity.Product{})
}
