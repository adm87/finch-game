package scene

import (
	"sync"

	"github.com/adm87/finch-collision/colliders"
	"github.com/adm87/finch-collision/collision"
	"github.com/adm87/finch-core/enum"
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/fsys"
	"github.com/adm87/finch-core/hashset"
	"github.com/adm87/finch-game/actor"
	"github.com/adm87/finch-tiled/tiled"
)

// ======================================================
// Build Scene Components
// ======================================================

// The idea is that Tiled Map scenes can be huge, so we'll build the
// components in parallel to speed things up.

func buildSceneComponents(ctx finch.Context, tmx *tiled.TMX, factories *SceneFactories) (
	staticCollision hashset.Set[collision.Collider],
	actors hashset.Set[actor.Actor],
) {
	wg := sync.WaitGroup{}

	staticCollisionCh := make(chan hashset.Set[collision.Collider], 1)
	wg.Add(1)
	go func(c finch.Context, tmx *tiled.TMX) {
		defer wg.Done()
		staticCollisionCh <- buildCollision(c, tmx)
	}(ctx, tmx)

	actorsCh := make(chan hashset.Set[actor.Actor], 1)
	wg.Add(1)
	go func(c finch.Context, tmx *tiled.TMX, f *SceneFactories) {
		defer wg.Done()
		actorsCh <- buildActors(c, tmx, f)
	}(ctx, tmx, factories)

	wg.Wait()

	staticCollision = <-staticCollisionCh
	actors = <-actorsCh

	close(staticCollisionCh)
	close(actorsCh)
	return
}

// ======================================================
// Actors
// ======================================================

func buildActors(ctx finch.Context, tmx *tiled.TMX, factories *SceneFactories) hashset.Set[actor.Actor] {
	results := hashset.New[actor.Actor]()

	if actorsGroup := tmx.ObjectGroupByProperty("SceneLayer", SceneActorLayer.String()); actorsGroup != nil {
		for _, obj := range actorsGroup.Objects {
			var a actor.Actor

			if obj.HasTemplate() {
				a = actorFromTemplate(obj, tiled.MustGetTX(finch.AssetFile(obj.Template())), tmx, factories)
			} else {
				a = actorFromObject(obj, tmx, factories)
			}

			results.Add(a)
		}
	}

	return results
}

func actorFromTemplate(instance *tiled.Object, tmpl *tiled.TX, tmx *tiled.TMX, factories *SceneFactories) actor.Actor {
	if at, ok := tmpl.Object.PropertyOfType("ActorType"); ok {
		actorType := fsys.MustGet(enum.Value[actor.ActorType](at.Value()))

		factory, exists := factories.actors[actorType]
		if !exists {
			return nil
		}

		return factory.FromTemplate(instance, tmpl, tmx)
	}
	return nil
}

func actorFromObject(obj *tiled.Object, tmx *tiled.TMX, factories *SceneFactories) actor.Actor {
	if at, ok := obj.PropertyOfType("ActorType"); ok {
		actorType := fsys.MustGet(enum.Value[actor.ActorType](at.Value()))

		factory, exists := factories.actors[actorType]
		if !exists {
			return nil
		}

		return factory.FromObject(obj, tmx)
	}
	return nil
}

// ======================================================
// Collision
// ======================================================

func buildCollision(ctx finch.Context, tmx *tiled.TMX) hashset.Set[collision.Collider] {
	result := hashset.New[collision.Collider]()

	if collisionGroup := tmx.ObjectGroupByProperty("SceneLayer", SceneCollisionLayer.String()); collisionGroup != nil {
		for _, obj := range collisionGroup.Objects {
			var collider *colliders.BoxCollider

			if obj.HasTemplate() {
				collider = colliderFromTemplate(obj, finch.AssetFile(obj.Template()))
			} else {
				collider = colliderFromObject(obj)
			}

			result.Add(collider)
		}
	}

	return result
}

func colliderFromTemplate(instance *tiled.Object, tmpl finch.AssetFile) *colliders.BoxCollider {
	collider := colliderFromObject(tiled.MustGetTX(tmpl).Object)

	if x := instance.X(); float64(x) != collider.X() {
		collider.SetX(float64(x))
	}
	if y := instance.Y(); float64(y) != collider.Y() {
		collider.SetY(float64(y))
	}
	if w := instance.Width(); float64(w) != collider.Width() {
		collider.SetWidth(float64(w))
	}
	if h := instance.Height(); float64(h) != collider.Height() {
		collider.SetHeight(float64(h))
	}

	return collider
}

func colliderFromObject(obj *tiled.Object) *colliders.BoxCollider {
	collider := colliders.NewBoxCollider(
		float64(obj.X()),
		float64(obj.Y()),
		float64(obj.Width()),
		float64(obj.Height()),
	)

	if colliderProps, ok := obj.PropertyOfType("BoxCollider"); ok {
		if detection, ok := colliderProps.PropertyOfType("CollisionDetectionType"); ok {
			collider.SetDetectionType(fsys.MustGet(enum.Value[collision.CollisionDetectionType](detection.Value())))
		}
		if layer, ok := colliderProps.PropertyOfType("CollisionLayer"); ok {
			collider.SetCollisionLayer(fsys.MustGet(enum.Value[collision.CollisionLayer](layer.Value())))
		}
		if ctype, ok := colliderProps.PropertyOfType("ColliderType"); ok {
			collider.SetColliderType(fsys.MustGet(enum.Value[collision.ColliderType](ctype.Value())))
		}
	}

	return collider
}
