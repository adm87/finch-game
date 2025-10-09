package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Update(ctx finch.Context) {
	pollInput(ctx)

	gameScene.Update(ctx.Time().DeltaSeconds())
}

func FixedUpdate(ctx finch.Context) {
}

func LateUpdate(ctx finch.Context) {
}

func pollInput(ctx finch.Context) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		finch.Exit()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {

	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF2) {

	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF5) {
		debugDrawTiled = !debugDrawTiled
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF6) {
		debugDrawColliders = !debugDrawColliders
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF7) {
		debugDrawCollisionGrid = !debugDrawCollisionGrid
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF8) {
		debugDrawSceneTree = !debugDrawSceneTree
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	camera := gameScene.Camera()
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		camera.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		camera.Y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		camera.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		camera.X += 1
	}
}
