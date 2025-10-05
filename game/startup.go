package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-game/data"
	"github.com/adm87/finch-game/level"
	"github.com/adm87/finch-tiled/tiled"
)

var camera *Camera
var (
	debugColliders     = true
	debugCollisionGrid = false
)

func Startup(ctx finch.Context) {
	finch.MustLoadAssets(
		data.Tile0000,
		data.TilemapCharactersPacked,
		data.TilemapExampleA,
		data.TilemapExampleB,
		data.TilemapPacked,
		data.TilesetCharacters,
		data.TilesetTiles,
		data.Player,
		data.Coin,
		data.MapCollider,
	)

	camera = NewCamera(float64(ctx.Screen().Width()), float64(ctx.Screen().Height()))
	camera.X = camera.width / 2
	camera.Y = camera.height / 2

	level.SetupLevel(tiled.MustGetTMX(data.TilemapExampleA))
}
