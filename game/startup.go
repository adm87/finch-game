package game

import (
	"github.com/adm87/finch-collision/collision"
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/geom"
	"github.com/adm87/finch-game/data"
	"github.com/adm87/finch-tiled/tiled"
)

var selectedMap *tiled.TMX
var camera *Camera

var panX, panY float64

var drawColliders = true
var drawCollisionGrid = true

var collisionWorld *collision.CollisionWorld
var debugCollider geom.Rect64

func Startup(ctx finch.Context) {
	finch.MustLoadAssets(
		data.TilemapCharactersPacked,
		data.TilemapExampleA,
		data.TilemapExampleB,
		data.TilemapPacked,
		data.TilesetCharacters,
		data.TilesetTiles,
	)

	camera = NewCamera(float64(ctx.Screen().Width()), float64(ctx.Screen().Height()))
	camera.X = camera.width / 2
	camera.Y = camera.height / 2

	collisionWorld = collision.NewCollisionWorld(9)

	debugCollider = geom.NewRect64(0, 0, 16, 16)

	setup_level(tiled.MustGetTMX(data.TilemapExampleA))
}

func setup_level(tmx *tiled.TMX) {
	collisionWorld.Clear()

	selectedMap = tmx

	collisionObjects, ok := tmx.GetObjectGroupByName("collision")
	if !ok {
		return
	}

	for _, obj := range collisionObjects.Objects {
		rect := geom.NewRect64(
			float64(obj.X()),
			float64(obj.Y()),
			float64(obj.Width()),
			float64(obj.Height()),
		)
		collisionWorld.AddCollider(&rect)
	}

	collisionWorld.AddCollider(&debugCollider)
}
