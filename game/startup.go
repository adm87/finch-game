package game

import (
	"github.com/adm87/finch-collision/colliders"
	"github.com/adm87/finch-collision/collision"
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-game/data"
	"github.com/adm87/finch-tiled/tiled"
)

const (
	PlayerCollisionLayer collision.CollisionLayer = 1 << iota
	PlatformCollisionLayer
)

type TestCollisionHandler struct{}

var selectedMap *tiled.TMX
var camera *Camera

var panX, panY float64

var drawColliders = true
var drawCollisionGrid = false

var collisionWorld *collision.CollisionWorld
var debugCollider *colliders.BoxCollider

var targetX, targetY float64

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
	collisionWorld.RegisterHandlers(
		PlayerCollisionLayer|PlatformCollisionLayer,
		&TestCollisionHandler{},
	)

	debugCollider = colliders.NewBoxCollider(0, 0, 16, 16)
	debugCollider.AddToLayer(PlayerCollisionLayer | PlatformCollisionLayer)
	debugCollider.SetType(collision.ColliderDynamic)
	debugCollider.SetCollisionDetection(collision.CollisionDetectionContinuous)

	setupLevel(tiled.MustGetTMX(data.TilemapExampleA))
}

func setupLevel(tmx *tiled.TMX) {
	collisionWorld.Clear()

	selectedMap = tmx

	collisionObjects, ok := tmx.GetObjectGroupByName("collision")
	if !ok {
		return
	}

	for _, obj := range collisionObjects.Objects {
		boxCollider := colliders.NewBoxCollider(
			float64(obj.X()),
			float64(obj.Y()),
			float64(obj.Width()),
			float64(obj.Height()),
		)
		boxCollider.AddToLayer(PlatformCollisionLayer)

		collisionWorld.AddCollider(boxCollider)
	}

	collisionWorld.AddCollider(debugCollider)
}
