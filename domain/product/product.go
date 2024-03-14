package product

import "github.com/gofiber/fiber/v2"

type Product struct {
}

func (product Product) GetOne(c *fiber.Ctx) error {
	productId := c.Params("id")
	return c.SendString(productId)
}

func (product Product) GetAll(c *fiber.Ctx) error {
	query := c.Query("ids")
	if query == "" {
		return nil
	}
	return nil
}
