package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/fsys"
	"github.com/adm87/finch-core/geom"
	"github.com/adm87/finch-game/data"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Update(ctx finch.Context) {
	poll_input(ctx)

	movement := geom.NewPoint64(panX, panY).
		Normalized()

	camera.X += movement.X * 100 * ctx.Time().DeltaSeconds()
	camera.Y += movement.Y * 100 * ctx.Time().DeltaSeconds()

	viewport := camera.Viewport()
	viewMatrix := camera.ViewMatrix()

	hw := viewport.Width / 2
	hh := viewport.Height / 2

	maxx := float64(selectedMap.Width())*float64(selectedMap.TileWidth()) - hw
	maxy := float64(selectedMap.Height())*float64(selectedMap.TileHeight()) - hh

	camera.X = fsys.Clamp(camera.X, hw, maxx)
	camera.Y = fsys.Clamp(camera.Y, hh, maxy)

	invt := viewMatrix
	invt.Invert()

	sx, sy := ebiten.CursorPosition()
	wx, wy := invt.Apply(float64(sx), float64(sy))

	hw = dynamicCollider.Width / 2.0
	hh = dynamicCollider.Height / 2.0

	dynamicCollider.X = float64(wx) - hw
	dynamicCollider.Y = float64(wy) - hh
}

func poll_input(ctx finch.Context) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		finch.Exit()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF9) {
		drawColliders = !drawColliders
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		set_map(tiled.MustGetTMX(data.TilemapExampleA))
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		set_map(tiled.MustGetTMX(data.TilemapExampleB))
	}

	panX, panY = 0, 0

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		panX -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		panX += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		panY -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		panY += 1
	}

	if _, y := ebiten.Wheel(); y != 0 {
		newZoom := fsys.Clamp(camera.Zoom()+y*0.1, 1, 3)
		camera.SetZoom(newZoom)
	}
}
