package game

import (
	"math"

	"github.com/adm87/finch-collision/colliders"
	"github.com/adm87/finch-collision/collision"
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-game/data"
	"github.com/adm87/finch-tiled/tiled"
)

var (
	PlayerLayer   = collision.NewCollisionLayer("Player")
	PlatformLayer = collision.NewCollisionLayer("Platform")
)

type TestCollisionHandler struct{}

var selectedMap *tiled.TMX
var camera *Camera

var panX, panY float64

var drawColliders = true
var drawCollisionGrid = false
var drawTiledMap = true

var collisionWorld *collision.CollisionWorld
var debugCollider *colliders.BoxCollider

var targetX, targetY float64

func Startup(ctx finch.Context) {
	finch.MustLoadAssets(
		data.Tile0000,
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

	collisionWorld.AddCollisionRules(
		PlayerLayer, collision.CollisionRules{
			PlatformLayer: BlockMovement,
		},
	)

	targetX = camera.X
	targetY = camera.Y

	debugCollider = colliders.NewBoxCollider(targetX, targetY, 12, 14)
	debugCollider.SetLayer(PlayerLayer)
	debugCollider.SetType(collision.ColliderDynamic)
	debugCollider.SetDetectionType(collision.CollisionDetectionContinuous)

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
		boxCollider.SetLayer(PlatformLayer)

		collisionWorld.AddCollider(boxCollider)
	}

	collisionWorld.AddCollider(debugCollider)
}

func BlockMovement(contact *collision.ContactInfo) {
	boxA := contact.ColliderA.(*colliders.BoxCollider)
	boxB := contact.ColliderB.(*colliders.BoxCollider)

	minxB, minyB := boxB.AABB().Min()
	maxxB, maxyB := boxB.AABB().Max()

	if contact.Normal.X != 0 {
		if contact.Normal.X < 0 {
			boxA.X = minxB - boxA.AABB().Width
			targetX = math.Min(targetX, boxA.X)
		} else {
			boxA.X = maxxB
			targetX = math.Max(targetX, boxA.X)
		}
	} else if contact.Normal.Y != 0 {
		if contact.Normal.Y < 0 {
			boxA.Y = minyB - boxA.AABB().Height
			targetY = math.Min(targetY, boxA.Y)
		} else {
			boxA.Y = maxyB
			targetY = math.Max(targetY, boxA.Y)
		}
	}

	collisionWorld.UpdateCollider(boxA)
}
