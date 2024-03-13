package catalogue

import (
	"github.com/gofiber/fiber/v2"
	"github.com/noxymon/ecomm-go/domain/catalogue/internal/handler"
)

func Register(app *fiber.App) {
	homePageHandler := handler.HomePage{}
	productHandler := handler.Product{}
	categoriesHandler := handler.Categories{}

	app.Get("/homepage", homePageHandler.Show)
	app.Get("/products", productHandler.GetAll)
	app.Get("/products/:id", productHandler.GetOne)

	app.Get("/categories", categoriesHandler.GetAll)
	app.Get("/categories/:id", categoriesHandler.GetOne)
}
