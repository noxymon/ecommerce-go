package category

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Handler struct {
	S *ServiceCategory
}

func newHandler(category *ServiceCategory) *Handler {
	return &Handler{category}
}

func (c Handler) GetOne(ctx *fiber.Ctx) error {
	categoryId := ctx.Params("id")
	catIdInt, err := strconv.Atoi(categoryId)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	category, err := c.S.getCategoryById(ctx.Context(), catIdInt)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return ctx.JSON(category)
}

func (c Handler) GetAll(ctx *fiber.Ctx) error {
	query := ctx.Query("ids")
	if query == "" {
		return nil
	}
	return nil
}
