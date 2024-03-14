package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/noxymon/ecomm-go/domain"
	"github.com/noxymon/ecomm-go/domain/category"
	"github.com/noxymon/ecomm-go/pkg/commonfx"
	"github.com/noxymon/ecomm-go/pkg/dbfx"
	"github.com/noxymon/ecomm-go/pkg/httpfx"
	"go.uber.org/fx"
)

func main() {
	container := fx.New(
		commonfx.Module,
		httpfx.Module,
		dbfx.Module,

		category.Module,

		fx.Invoke(
			fx.Annotate(
				func(handler *category.Handler) {
					categoryRoutes := domain.RegisterCategory(handler)
					route := appendRoute(*httpfx.Routes, *categoryRoutes)
					httpfx.Routes = &route
				},
				fx.ParamTags(`name:"categoryHandler"`),
			),
		),
	)
	container.Run()
}

func appendRoute(maps ...map[string]fiber.Handler) map[string]fiber.Handler {
	var out = make(map[string]fiber.Handler)

	for _, mapV := range maps {
		for k, v := range mapV {
			out[k] = v
		}
	}

	return out
}
