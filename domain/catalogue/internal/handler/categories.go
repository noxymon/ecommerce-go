package handler

import "github.com/gofiber/fiber/v2"

type Categories struct {
}

func (c Categories) GetOne(ctx *fiber.Ctx) error {
	_ = ctx.Params("id")
	return nil
}

func (c Categories) GetAll(ctx *fiber.Ctx) error {
	query := ctx.Query("ids")
	if query == "" {
		return nil
	}
	return nil
}
