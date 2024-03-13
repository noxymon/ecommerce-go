package handler

import "github.com/gofiber/fiber/v2"

type Product struct {
}

func (product Product) GetOne(c *fiber.Ctx) error {
	_ = c.Params("id")
	return nil
}

func (product Product) GetAll(c *fiber.Ctx) error {
	query := c.Query("ids")
	if query == "" {
		return nil
	}
	return nil
}
