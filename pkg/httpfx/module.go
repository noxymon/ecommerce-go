package httpfx

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Routes *map[string]fiber.Handler

var Module = fx.Module(
	"httpfx",
	fx.Provide(
		fx.Annotate(
			newGofiber,
			fx.ParamTags(`name:"commonfx"`),
		),
	),
	fx.Invoke(
		useGofiber,
	),
)
