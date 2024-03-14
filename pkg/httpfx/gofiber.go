package httpfx

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/noxymon/ecomm-go/pkg/commonfx"
	"github.com/rs/zerolog"
	"go.uber.org/fx"
	"strconv"
	"strings"
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

func useGofiber(lifecycle fx.Lifecycle, app *fiber.App) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			registerPathWith(app)

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

func registerPathWith(app *fiber.App) {
	for path, handler := range *Routes {
		method, endpointPath := getMethodAndEndpointPathFrom(path)
		log.Info().Msg(fmt.Sprintf("Register handler %s with %s", endpointPath, method))

		switch method {
		case "GET":
			app.Get(endpointPath, handler)
		case "POST":
			app.Post(endpointPath, handler)
		case "DELETE":
			app.Delete(endpointPath, handler)
		case "PATCH":
			app.Patch(endpointPath, handler)
		default:
			app.Get(endpointPath, handler)
		}
	}
}

func getMethodAndEndpointPathFrom(path string) (string, string) {
	splitString := strings.Split(path, " ")
	method := getHttpMethodFrom(splitString)
	endpointPath := splitString[1]
	return method, endpointPath
}

func getHttpMethodFrom(splitString []string) string {
	method := "GET"
	if len(splitString[0]) > 0 {
		method = splitString[0]
	}
	return method
}

func getPortFrom(portFromConfig int) string {
	return ":" + strconv.Itoa(portFromConfig)
}
