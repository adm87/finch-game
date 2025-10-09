package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw(ctx finch.Context, screen *ebiten.Image) {
	if debugDrawTiled {
		gameScene.Draw(screen)
	}
	if debugDrawColliders {
		gameScene.DebugDrawColliders(screen)
	}
	if debugDrawCollisionGrid {
		gameScene.DebugDrawCollisionGrid(screen)
	}
	if debugDrawSceneTree {
		gameScene.DebugDrawSceneTree(screen)
	}
}
