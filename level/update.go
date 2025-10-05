package level

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/fsys"
	"github.com/adm87/finch-core/geom"
	"github.com/hajimehoshi/ebiten/v2"
)

func Update(ctx finch.Context, viewport geom.Rect64, viewMatrix ebiten.GeoM) {
	toWorld := viewMatrix
	toWorld.Invert()

	sx, sy := ebiten.CursorPosition()
	tx, ty = toWorld.Apply(float64(sx), float64(sy))

	tx -= debugCollider.Width / 2
	ty -= debugCollider.Height / 2

	bounds := activeLevel.Bounds()
	minx, miny := bounds.Min()
	maxx, maxy := bounds.Max()

	tx = fsys.Clamp(tx, minx, maxx-debugCollider.Width)
	ty = fsys.Clamp(ty, miny, maxy-debugCollider.Height)

	debugCollider.X = fsys.Lerp(debugCollider.X, tx, 0.5)
	debugCollider.Y = fsys.Lerp(debugCollider.Y, ty, 0.5)

	collisionWorld.UpdateCollider(debugCollider)
	collisionWorld.CheckForCollisions(ctx.Time().FixedMilli())
}

func FixedUpdate(ctx finch.Context, viewport geom.Rect64, viewMatrix ebiten.GeoM) {

}

func LateUpdate(ctx finch.Context, viewport geom.Rect64, viewMatrix ebiten.GeoM) {

}
