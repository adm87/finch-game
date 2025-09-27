package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-game/data"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw(ctx finch.Context, screen *ebiten.Image) {
	tiled.Draw(ctx, screen, tiled.MustGetTMX(data.TilemapExampleA))
}
