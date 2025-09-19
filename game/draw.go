package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-game/data"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
)

var img *ebiten.Image

func Draw(ctx finch.Context, screen *ebiten.Image) {
	if img == nil {
		img = tiled.Buffer(ctx, data.TilemapExampleA)
	}

	screen.DrawImage(img, &ebiten.DrawImageOptions{})
}
