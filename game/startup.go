package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-game/data"
	"github.com/adm87/finch-resources/resources"
	"github.com/adm87/finch-tiled/tiled"
)

var camera *Camera
var selectedMap *tiled.TMX

func Startup(ctx finch.Context) {
	resources.Load(ctx,
		data.TilemapExampleA,
		data.TilemapExampleB,
		data.TilemapInfinite,
	)

	width := float64(ctx.Screen().Width())
	height := float64(ctx.Screen().Height())

	camera = NewCamera(width, height)

	camera.X = width / 2
	camera.Y = height / 2

	selectedMap, _ = tiled.GetTmx(data.TilemapExampleA)
}
