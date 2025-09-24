package module

import (
	"github.com/adm87/finch-core/finch"

	resources "github.com/adm87/finch-resources/module"
	tiled "github.com/adm87/finch-tiled/module"
)

func Register(ctx finch.Context) {
	resources.Register(ctx)
	tiled.Register(ctx)

	ctx.Logger().Info("game module registered")
}
