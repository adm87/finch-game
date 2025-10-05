package level

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/geom"
	"github.com/hajimehoshi/ebiten/v2"

	collisionDebug "github.com/adm87/finch-collision/debug"
)

var rectImage = ebiten.NewImage(1, 1)

func Draw(ctx finch.Context, screen *ebiten.Image, viewport geom.Rect64, view ebiten.GeoM) {
	// tiled.DrawScene(ctx, screen, activeLevel, viewport, view)

	// DebugDrawQuadTree(screen, viewport, view)
}

func DebugDrawColliders(screen *ebiten.Image, viewport geom.Rect64, view ebiten.GeoM) {
	collisionDebug.DrawColliders(collisionWorld, screen, debugCollider.Bounds(), view)
}

func DebugDrawCollisionGrid(screen *ebiten.Image, viewport geom.Rect64, view ebiten.GeoM) {
	collisionDebug.DrawCollisionGrid(collisionWorld, screen, debugCollider.Bounds(), view)
}
