package handler

import (
	"fmt"
	"toko-kami/database"
	"toko-kami/model/entity"
	"toko-kami/model/request"

	"github.com/gofiber/fiber/v2"
)

func GetProduct(c *fiber.Ctx) error {
	var product []entity.Product

	result := database.DB.Find(&product)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	var ProductResponses []request.ProductResponse

	for _, product := range product {
		ProductResponse := request.ProductResponse{
			ID:          product.ID,
			Title:       product.Title,
			Description: product.Description,
			Price:       product.Price,
			Image:       product.Image,
			CreatedAt:   product.CreatedAt,
		}
		ProductResponses = append(ProductResponses, ProductResponse)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": ProductResponses,
	})
}

func GetByIdProduct(c *fiber.Ctx) error {
	productId := c.Params("id")

	var product entity.Product

	result := database.DB.First(&product, productId)

	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "product not found",
		})
	}

	ProductResponse := request.ProductResponse{
		ID:          product.ID,
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		Image:       product.Image,
		CreatedAt:   product.CreatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": ProductResponse,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	// productRequest := new(request.ProductRequest)
	//
	// productRequest.Title = c.FormValue("title")
	// productRequest.Description = c.FormValue("description")
	// productRequest.Price = c.FormValue("price")
	//
	// ImageCover, err := c.FormFile("image")
	// if err == nil {
	//   filename := ImageCover.Filename
	//   productRequest.Image = fmt.Sprintf("/image/product/%s", filename)
	//   if err := c.SaveFile(ImageCover, fmt.Sprint("./public/%s"), productRequest);
	// }

	return nil
}

func DeleteProduct(c *fiber.Ctx) error {
	productId := c.Params("id")

	result := database.DB.Delete(&entity.Product{}, productId)

	if result.Error != nil || result.RowsAffected == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "success delete product",
	})
}
