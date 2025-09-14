package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func Draw(ctx finch.Context, screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, Finch Game!")
}
