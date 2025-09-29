package game

import (
	collisionDebug "github.com/adm87/finch-collision/debug"

	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw(ctx finch.Context, screen *ebiten.Image) {
	viewMatrix := camera.ViewMatrix()
	viewport := camera.Viewport()

	tiled.DrawScene(ctx, screen, selectedMap, viewport, viewMatrix)

	if drawColliders {
		collisionDebug.DrawColliders(collisionWorld, screen, debugCollider, viewMatrix)
	}
	if drawCollisionGrid {
		collisionDebug.DrawCollisionGrid(collisionWorld, screen, viewport, viewMatrix)
	}
}
