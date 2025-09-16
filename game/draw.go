package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func Draw(ctx finch.Context, screen *ebiten.Image) {
	tmb := tiled.Buffer(ctx, "test_tilemap")
	if tmb == nil {
		ebitenutil.DebugPrint(screen, "Tilemap not loaded")
		return
	}

	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(tmb, op)

	ebitenutil.DebugPrint(screen, "Hello, Finch Game!")
}
