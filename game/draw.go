package game

import (
	"image/color"

	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
)

var lvlDrawOptions = &ebiten.DrawImageOptions{
	Filter: ebiten.FilterLinear,
}

func Draw(ctx finch.Context, screen *ebiten.Image) {
	lvlTexture.Fill(color.RGBA{100, 149, 237, 255})

	tiled.DrawScene(ctx, lvlTexture, selectedMap, camera.Viewport(), camera.ViewMatrix())

	screen.DrawImage(lvlTexture, lvlDrawOptions)
}
