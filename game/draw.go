package game

import (
	collisionDebug "github.com/adm87/finch-collision/debug"
	"github.com/adm87/finch-game/data"

	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw(ctx finch.Context, screen *ebiten.Image) {
	viewMatrix := camera.ViewMatrix()
	viewport := camera.Viewport()

	if drawTiledMap {
		tiled.DrawScene(ctx, screen, selectedMap, viewport, viewMatrix)
	}

	drawPlayer(screen, viewMatrix)

	if drawColliders {
		collisionDebug.DrawColliders(collisionWorld, screen, debugCollider.AABB(), viewMatrix)
	}
	if drawCollisionGrid {
		collisionDebug.DrawCollisionGrid(collisionWorld, screen, viewport, viewMatrix)
	}
}

func drawPlayer(screen *ebiten.Image, viewMatrix ebiten.GeoM) {
	img := finch.MustGetImage(data.Tile0000)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(-img.Bounds().Dx()/2),
		float64(-img.Bounds().Dy()),
	)
	op.GeoM.Translate(debugCollider.X+debugCollider.Width/2, debugCollider.Y+debugCollider.Height)
	op.GeoM.Concat(viewMatrix)

	screen.DrawImage(img, op)
}
