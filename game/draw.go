package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-resources/images"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func Draw(ctx finch.Context, screen *ebiten.Image) {
	s := ctx.Screen()
	i, _ := images.Get("tile_0000")

	shw := float64(s.Width()) / 2.0
	shh := float64(s.Height()) / 2.0

	ihw := float64(i.Bounds().Dx()) / 2.0
	ihh := float64(i.Bounds().Dy()) / 2.0

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-ihw, -ihh)
	op.GeoM.Translate(shw, shh)

	screen.DrawImage(i, op)

	ebitenutil.DebugPrint(screen, "Hello, Finch Game!")
}
