package category

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"catalogue/category",
	fx.Provide(
		fx.Annotate(
			newServiceCategory,
			fx.ParamTags(`name:"dbfx.query"`),
			fx.ResultTags(`name:"serviceCategory"`),
		),
		fx.Annotate(
			newHandler,
			fx.ParamTags(`name:"serviceCategory"`),
			fx.ResultTags(`name:"categoryHandler"`),
		),
	),
)
