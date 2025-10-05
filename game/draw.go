package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-game/level"
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw(ctx finch.Context, screen *ebiten.Image) {
	level.Draw(ctx, screen, camera.Viewport(), camera.ViewMatrix())

	if debugColliders {
		level.DebugDrawColliders(screen, camera.Viewport(), camera.ViewMatrix())
	}
	if debugCollisionGrid {
		level.DebugDrawCollisionGrid(screen, camera.Viewport(), camera.ViewMatrix())
	}
}
