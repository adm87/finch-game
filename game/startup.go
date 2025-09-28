package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/geom"
	"github.com/adm87/finch-game/data"
	"github.com/adm87/finch-tiled/tiled"
)

var selectedMap *tiled.TMX
var panX, panY float64
var camera *Camera
var staticColliders []geom.Rect64
var dynamicCollider geom.Rect64
var drawColliders = true

func Startup(ctx finch.Context) {
	finch.MustLoadAssets(
		data.TilemapCharactersPacked,
		data.TilemapExampleA,
		data.TilemapExampleB,
		data.TilemapPacked,
		data.TilesetCharacters,
		data.TilesetTiles,
	)

	set_map(tiled.MustGetTMX(data.TilemapExampleA))

	camera = NewCamera(float64(ctx.Screen().Width()), float64(ctx.Screen().Height()))
	camera.X = camera.width / 2
	camera.Y = camera.height / 2

	dynamicCollider = geom.NewRect64(0, 0, 16, 16)
}

func set_map(tmx *tiled.TMX) {
	selectedMap = tmx

	collisionLayer, ok := selectedMap.GetObjectGroupByName("collision")
	if !ok {
		staticColliders = nil
		return
	}

	staticColliders = make([]geom.Rect64, len(collisionLayer.Objects))
	for i, obj := range collisionLayer.Objects {
		staticColliders[i] = geom.NewRect64(
			float64(obj.X()),
			float64(obj.Y()),
			float64(obj.Width()),
			float64(obj.Height()),
		)
	}
}
