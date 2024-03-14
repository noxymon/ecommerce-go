package homepage

import "github.com/gofiber/fiber/v2"

type HomePage struct {
}

func (h HomePage) Show(ctx *fiber.Ctx) error {
	return nil
}
