package game

import (
	"github.com/adm87/finch-collision/colliders"
	"github.com/adm87/finch-collision/collision"
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/fsys"
	"github.com/adm87/finch-core/geom"
	"github.com/adm87/finch-game/data"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (h *TestCollisionHandler) OnCollision(contact *collision.ContactInfo) {
	boxA := contact.ColliderA.(*colliders.BoxCollider)
	boxB := contact.ColliderB.(*colliders.BoxCollider)

	minxB, minyB := boxB.AABB().Min()
	maxxB, maxyB := boxB.AABB().Max()

	if contact.Normal.X != 0 {
		if contact.Normal.X < 0 {
			boxA.X = minxB - boxA.AABB().Width
		} else {
			boxA.X = maxxB
		}
	} else if contact.Normal.Y != 0 {
		if contact.Normal.Y < 0 {
			boxA.Y = minyB - boxA.AABB().Height
		} else {
			boxA.Y = maxyB
		}
	}

	collisionWorld.UpdateCollider(boxA)
}

func Update(ctx finch.Context) {
	pollInput(ctx)

	movement := geom.NewPoint64(panX, panY).
		Normalized()

	camera.X += movement.X * 100 * ctx.Time().DeltaSeconds()
	camera.Y += movement.Y * 100 * ctx.Time().DeltaSeconds()

	viewport := camera.Viewport()
	viewMatrix := camera.ViewMatrix()

	hw := viewport.Width / 2
	hh := viewport.Height / 2

	maxx := float64(selectedMap.Width())*float64(selectedMap.TileWidth()) - hw
	maxy := float64(selectedMap.Height())*float64(selectedMap.TileHeight()) - hh

	camera.X = fsys.Clamp(camera.X, hw, maxx)
	camera.Y = fsys.Clamp(camera.Y, hh, maxy)

	toWorld := viewMatrix
	toWorld.Invert()

	mx, my := ebiten.CursorPosition()
	wx, wy := toWorld.Apply(float64(mx), float64(my))

	targetX = wx - debugCollider.Width/2
	targetY = wy - debugCollider.Height/2
}

func FixedUpdate(ctx finch.Context) {
	debugCollider.X = fsys.Lerp(debugCollider.X, targetX, ctx.Time().FixedSeconds()*10)
	debugCollider.Y = fsys.Lerp(debugCollider.Y, targetY, ctx.Time().FixedSeconds()*10)
	collisionWorld.UpdateCollider(debugCollider)

	collisionWorld.CheckForCollisions(ctx.Time().FixedSeconds())
}

func pollInput(ctx finch.Context) {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		finch.Exit()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyF9) {
		drawColliders = !drawColliders
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF10) {
		drawCollisionGrid = !drawCollisionGrid
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF11) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
	}

	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		setupLevel(tiled.MustGetTMX(data.TilemapExampleA))
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		setupLevel(tiled.MustGetTMX(data.TilemapExampleB))
	}

	panX, panY = 0, 0

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		panX -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		panX += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		panY -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		panY += 1
	}

	if _, y := ebiten.Wheel(); y != 0 {
		newZoom := fsys.Clamp(camera.Zoom()+y*0.1, 1, 3)
		camera.SetZoom(newZoom)
	}
}
