package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/linq"
	"github.com/adm87/finch-resources/resources"
)

func Startup(ctx finch.Context) {
	// Load the full manifest for testing purposes
	resources.Load(ctx, linq.Keys(resources.Manifest())...)
}
