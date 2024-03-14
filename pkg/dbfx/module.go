package dbfx

import "go.uber.org/fx"

var Module = fx.Module(
	"dbfx",
	fx.Provide(
		fx.Annotate(
			newGorm,
			fx.ResultTags(`name:"dbfx.query"`),
			fx.ParamTags(`name:"commonfx"`),
		),
	),
)
