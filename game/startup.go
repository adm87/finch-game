package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-game/data"
	"github.com/adm87/finch-tiled/tiled"
)

var selectedMap *tiled.TMX
var panX, panY float64
var camera *Camera

func Startup(ctx finch.Context) {
	finch.MustLoadAssets(
		data.TilemapCharactersPacked,
		data.TilemapExampleA,
		data.TilemapExampleB,
		data.TilemapPacked,
		data.TilesetCharacters,
		data.TilesetTiles,
	)

	selectedMap = tiled.MustGetTMX(data.TilemapExampleB)

	camera = NewCamera(float64(ctx.Screen().Width()), float64(ctx.Screen().Height()))
	camera.X = camera.width / 2
	camera.Y = camera.height / 2
}
