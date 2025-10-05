package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/fsys"
	"github.com/adm87/finch-game/data"
	"github.com/adm87/finch-game/level"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Update(ctx finch.Context) {
	pollInput(ctx)

	level.Update(ctx, camera.Viewport(), camera.ViewMatrix())
}

func FixedUpdate(ctx finch.Context) {
	level.FixedUpdate(ctx, camera.Viewport(), camera.ViewMatrix())
}

func LateUpdate(ctx finch.Context) {
	level.LateUpdate(ctx, camera.Viewport(), camera.ViewMatrix())
}

func pollInput(ctx finch.Context) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		finch.Exit()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		level.SetupLevel(tiled.MustGetTMX(data.TilemapExampleA))
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF2) {
		level.SetupLevel(tiled.MustGetTMX(data.TilemapExampleB))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF5) {
		debugColliders = !debugColliders
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF6) {
		debugCollisionGrid = !debugCollisionGrid
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if _, y := ebiten.Wheel(); y != 0 {
		newZoom := fsys.Clamp(camera.Zoom()+y*0.1, 1, 2)
		camera.SetZoom(newZoom)
	}
}
