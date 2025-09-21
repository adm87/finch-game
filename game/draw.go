package game

import (
	"image/color"

	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/geom"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func Draw(ctx finch.Context, screen *ebiten.Image) {
	renderTarget.Fill(color.RGBA{R: 100, G: 149, B: 237, A: 255})
	tiled.DrawRegion(ctx, renderTarget, selectedMap, camera.Viewport())

	sw, sh := ctx.Screen().Width(), ctx.Screen().Height()
	hw, hh := sw/2, sh/2

	x := float64(sw-hw) / 2
	y := float64(sh-hh) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x, y)

	draw_rect(screen, cameraOutline, camera.Viewport(), color.RGBA{R: 255, A: 255})
	screen.DrawImage(renderTarget, op)
	draw_rect(screen, imageOutline, geom.NewRect64(x, y, float64(renderTarget.Bounds().Dx()), float64(renderTarget.Bounds().Dy())), color.RGBA{G: 255, A: 255})

	ebitenutil.DebugPrintAt(screen, "Move the camera with the mouse cursor.\nPress F1 or F2 to switch tilemaps.\nPress ESC to exit.", 10, 10)
}

func draw_rect(screen *ebiten.Image, img *ebiten.Image, rect geom.Rect64, col color.Color) {
	path := vector.Path{}

	minx, miny := rect.Min()
	maxx, maxy := rect.Max()

	path.MoveTo(float32(minx), float32(miny))
	path.LineTo(float32(maxx), float32(miny))
	path.LineTo(float32(maxx), float32(maxy))
	path.LineTo(float32(minx), float32(maxy))
	path.Close()

	vs, is := path.AppendVerticesAndIndicesForStroke(nil, nil, &vector.StrokeOptions{
		Width: 1,
	})

	for i := range vs {
		vs[i].ColorR = float32(col.(color.RGBA).R) / 255
		vs[i].ColorG = float32(col.(color.RGBA).G) / 255
		vs[i].ColorB = float32(col.(color.RGBA).B) / 255
		vs[i].ColorA = float32(col.(color.RGBA).A) / 255
	}

	screen.DrawTriangles(vs, is, img, nil)
}
