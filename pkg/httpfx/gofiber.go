package httpfx

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/noxymon/ecomm-go/pkg/commonfx"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"strconv"
)

var log zerolog.Logger
var internalCommon *commonfx.Common

func newGofiber(common *commonfx.Common) *fiber.App {
	internalCommon = common
	log = internalCommon.Logger

	app := fiber.New()
	app.Use(
		logger.New(),
		recover.New(),
		compress.New(compress.Config{
			Level: compress.LevelBestCompression,
		}),
	)
	return app
}

func useGofiber(
	lifecycle fx.Lifecycle,
	app *fiber.App) {

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				portFromConfig := getPortFrom(internalCommon.Viper.GetInt("server.port"))
				if err := app.Listen(portFromConfig); err != nil {
					log.Error().Msg(err.Error())
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := app.Shutdown(); err != nil {
				log.Fatal().Msg(err.Error())
			}
			return nil
		},
	})
}

func getPortFrom(portFromConfig int) string {
	return ":" + strconv.Itoa(portFromConfig)
}
