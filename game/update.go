package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-game/data"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func Update(ctx finch.Context) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		finch.Exit()
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF1) {
		selectedMap = data.TilemapExampleA
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF2) {
		selectedMap = data.TilemapExampleB
	}

	x, y := ebiten.CursorPosition()

	camera.X = float64(x)
	camera.Y = float64(y)

	viewport := camera.Viewport()

	minx, miny := viewport.Min()
	maxx, maxy := viewport.Max()

	if minx < 0 {
		camera.X = viewport.Width() / 2
	}
	if miny < 0 {
		camera.Y = viewport.Height() / 2
	}
	if maxx > float64(ctx.Screen().Width()) {
		camera.X = float64(ctx.Screen().Width()) - (viewport.Width() / 2)
	}
	if maxy > float64(ctx.Screen().Height()) {
		camera.Y = float64(ctx.Screen().Height()) - (viewport.Height() / 2)
	}
}
