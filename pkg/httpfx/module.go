package httpfx

import "go.uber.org/fx"

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
