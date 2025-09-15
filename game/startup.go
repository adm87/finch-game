package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-resources/resources"
)

func Startup(ctx finch.Context) {
	resources.Load(ctx, "test_tilemap_infinite")
}
