package module

import (
	"github.com/adm87/finch-core/finch"
	resources "github.com/adm87/finch-resources/module"
)

func Register(ctx finch.Context) {
	resources.Register(ctx)

	ctx.Logger().Info("game module registered")
}
