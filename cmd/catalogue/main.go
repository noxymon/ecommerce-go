package main

import (
	"github.com/noxymon/ecomm-go/pkg/commonfx"
	"github.com/noxymon/ecomm-go/pkg/httpfx"
	"go.uber.org/fx"
)

func main() {
	container := fx.New(
		commonfx.Module,
		httpfx.Module,
	)
	container.Run()
}
