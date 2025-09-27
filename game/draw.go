package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw(ctx finch.Context, screen *ebiten.Image) {
	tiled.DrawScene(ctx, screen, selectedMap, camera.Viewport(), camera.ViewMatrix())
}
