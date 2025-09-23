package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-game/data"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Update(ctx finch.Context) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		finch.Exit()
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		selectedMap, _ = tiled.GetTmx(data.TilemapExampleA)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF2) {
		selectedMap, _ = tiled.GetTmx(data.TilemapExampleB)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF3) {
		selectedMap, _ = tiled.GetTmx(data.TilemapInfinite)
	}

	if _, scrollY := ebiten.Wheel(); scrollY != 0 {
		camera.SetZoom(camera.Zoom() * (1 + scrollY*0.1))
	}

	dx, dy := 0.0, 0.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		dx = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		dx = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		dy = -1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		dy = 1
	}

	camera.X += dx * 100 * ctx.Time().DeltaSeconds()
	camera.Y += dy * 100 * ctx.Time().DeltaSeconds()
}
