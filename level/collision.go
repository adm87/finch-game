package level

import (
	"github.com/adm87/finch-collision/colliders"
	"github.com/adm87/finch-collision/collision"
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/hashset"
	"github.com/adm87/finch-game/properties"
	"github.com/adm87/finch-tiled/tiled"
)

var (
	collisionWorld *collision.CollisionWorld

	staticColliderPool    = hashset.New[*colliders.BoxCollider]()
	activeStaticColliders = hashset.New[*colliders.BoxCollider]()
)

func CleanupMapCollision() {
	for collider := range activeStaticColliders {
		collisionWorld.RemoveCollider(collider)
		staticColliderPool.AddDistinct(collider)
	}
	activeStaticColliders.Clear()
}

func createMapCollision(level *tiled.TMX) {
	if collisionGroup, exists := level.GetObjectGroupByName("MapCollision"); exists {
		for _, obj := range collisionGroup.Objects {
			var collider *colliders.BoxCollider

			if obj.HasTemplate() {
				collider = colliderFromTemplate(obj, finch.AssetFile(obj.Template()))
			} else {
				collider = colliderFromObject(obj)
			}

			if collider != nil {
				collisionWorld.AddCollider(collider)
				activeStaticColliders.AddDistinct(collider)
			}
		}
	}
}

func colliderFromTemplate(instance *tiled.Object, tmpl finch.AssetFile) *colliders.BoxCollider {
	tx := tiled.MustGetTX(tmpl)

	collider := colliderFromObject(tx.Object)
	collider.X = float64(instance.X())
	collider.Y = float64(instance.Y())

	w := instance.Width()
	h := instance.Height()

	if w != 0 {
		collider.Width = float64(instance.Width())
	}
	if h != 0 {
		collider.Height = float64(instance.Height())
	}

	return collider
}

func colliderFromObject(obj *tiled.Object) *colliders.BoxCollider {
	props, exists := obj.PropertyOfType("CollisionProperties")
	if !exists {
		return nil
	}

	collisionProps, err := properties.NewCollisionProperties().FromTMXProperty(props)
	if err != nil {
		return nil
	}

	collider := getNextStaticCollider()
	collider.X = float64(obj.X())
	collider.Y = float64(obj.Y())
	collider.Width = float64(obj.Width())
	collider.Height = float64(obj.Height())
	collider.SetType(collisionProps.ColliderType)
	collider.SetLayer(collisionProps.CollisionLayer)

	return collider
}

func getNextStaticCollider() *colliders.BoxCollider {
	var collider *colliders.BoxCollider
	if len(staticColliderPool) > 0 {
		for c := range staticColliderPool {
			collider = c
			staticColliderPool.Remove(c)
			break
		}
	} else {
		collider = colliders.NewBoxCollider(0, 0, 0, 0)
	}
	activeStaticColliders.AddDistinct(collider)
	return collider
}
