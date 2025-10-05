package level

import (
	"github.com/adm87/finch-collision/colliders"
	"github.com/adm87/finch-collision/collision"
	"github.com/adm87/finch-core/partition/hashgrid"
	"github.com/adm87/finch-tiled/tiled"
)

var activeLevel *tiled.TMX
var debugCollider *colliders.BoxCollider
var tx, ty float64

var (
	PlayerCollisionLayer   = collision.NewCollisionLayer("Player")
	PlatformCollisionLayer = collision.NewCollisionLayer("Platform")
	PickupCollisionLayer   = collision.NewCollisionLayer("Pickup")
)

func SetupLevel(lvl *tiled.TMX) {
	TearDownLevel()

	lvlBounds := lvl.Bounds()

	collisionWorld = collision.NewCollisionWorld(
		hashgrid.New[collision.Collider](9),
	)

	collisionWorld.SetProfiles(collision.CollisionProfile{
		PlayerCollisionLayer: {
			PlatformCollisionLayer: func(contact *collision.ContactInfo) {
				boxA := contact.ColliderA.(*colliders.BoxCollider)
				boxB := contact.ColliderB.(*colliders.BoxCollider)

				minxB, minyB := boxB.Bounds().Min()
				maxxB, maxyB := boxB.Bounds().Max()

				if contact.Normal.X != 0 {
					if contact.Normal.X < 0 {
						boxA.X = minxB - boxA.Bounds().Width

					} else {
						boxA.X = maxxB
					}
				} else if contact.Normal.Y != 0 {
					if contact.Normal.Y < 0 {
						boxA.Y = minyB - boxA.Bounds().Height
					} else {
						boxA.Y = maxyB
					}
				}

				collisionWorld.UpdateCollider(boxA)
			},
		},
	})

	tx := lvlBounds.X + lvlBounds.Width/2
	ty := lvlBounds.Y + lvlBounds.Height/2

	debugCollider = colliders.NewBoxCollider(tx, ty, 18, 18)
	debugCollider.SetLayer(PlayerCollisionLayer)
	debugCollider.SetDetectionType(collision.CollisionDetectionContinuous)
	debugCollider.SetType(collision.ColliderDynamic)
	collisionWorld.AddCollider(debugCollider)

	createMapCollision(lvl)
	createActors(lvl)

	activeLevel = lvl
}

func TearDownLevel() {
	CleanupMapCollision()
	if collisionWorld != nil {
		collisionWorld.RemoveCollider(debugCollider)
	}
	activeLevel = nil
}
