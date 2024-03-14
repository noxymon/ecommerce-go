package domain

import (
	"github.com/gofiber/fiber/v2"
	"github.com/noxymon/ecomm-go/domain/category"
	"github.com/noxymon/ecomm-go/domain/homepage"
	"github.com/noxymon/ecomm-go/domain/product"
)

func RegisterCategory(handler *category.Handler) *map[string]fiber.Handler {
	return &map[string]fiber.Handler{
		"GET /categories":     handler.GetAll,
		"GET /categories/:id": handler.GetOne,
	}
}

func RegisterProduct() *map[string]fiber.Handler {
	productHandler := product.Product{}

	return &map[string]fiber.Handler{
		"GET /products":     productHandler.GetAll,
		"GET /products/:id": productHandler.GetOne,
	}
}

func RegisterHomepage() *map[string]fiber.Handler {
	homepageHandler := homepage.HomePage{}

	return &map[string]fiber.Handler{
		"GET /homepage": homepageHandler.Show,
	}
}
