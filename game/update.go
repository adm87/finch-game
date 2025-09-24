package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/fsys"
	"github.com/adm87/finch-game/data"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var dx, dy float64

func Update(ctx finch.Context) {
	poll_input()

	screenWidth := float64(ctx.Screen().Width())
	screenHeight := float64(ctx.Screen().Height())

	update_camera(ctx.Time().DeltaSeconds(), screenWidth, screenHeight)
}

func poll_input() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		finch.Exit()
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		selectedMap, _ = tiled.GetTmx(data.TilemapExampleA)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF2) {
		selectedMap, _ = tiled.GetTmx(data.TilemapExampleB)
	}
	if _, scrollY := ebiten.Wheel(); scrollY != 0 {
		camera.SetZoom(fsys.Clamp(camera.Zoom()*(1+scrollY*0.1), 1, 2))
	}

	dx, dy = 0, 0

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
}

func update_camera(dt float64, screenWidth, screenHeight float64) {
	camera.X += dx * 100 * dt
	camera.Y += dy * 100 * dt

	viewport := camera.Viewport()

	minx, miny := viewport.Min()
	if minx < 0 {
		camera.X = viewport.Width * 0.5
	}
	if miny < 0 {
		camera.Y = viewport.Height * 0.5
	}

	maxx, maxy := viewport.Max()
	if maxx > screenWidth {
		camera.X = screenWidth - viewport.Width*0.5
	}
	if maxy > screenHeight {
		camera.Y = screenHeight - viewport.Height*0.5
	}
}
