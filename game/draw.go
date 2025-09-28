package game

import (
	"image/color"

	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/geom"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var rectImg = ebiten.NewImage(1, 1)

func Draw(ctx finch.Context, screen *ebiten.Image) {
	viewMatrix := camera.ViewMatrix()
	viewport := camera.Viewport()

	tiled.DrawScene(ctx, screen, selectedMap, viewport, viewMatrix)

	if drawColliders {
		for _, rect := range staticColliders {
			draw_rect(ctx, screen, rect, color.RGBA{R: 255, A: 255}, viewport, viewMatrix)
		}
		draw_rect(ctx, screen, dynamicCollider, color.RGBA{G: 255, A: 255}, viewport, viewMatrix)
	}
}

func draw_rect(ctx finch.Context, screen *ebiten.Image, rect geom.Rect64, color color.RGBA, viewport geom.Rect64, viewMatrix ebiten.GeoM) {
	rectImg.Fill(color)

	minx, miny := rect.Min()
	maxx, maxy := rect.Max()

	vminx, vminy := viewport.Min()
	vmaxx, vmaxy := viewport.Max()

	if maxx < vminx || maxy < vminy || minx > vmaxx || miny > vmaxy {
		return
	}

	minx, miny = viewMatrix.Apply(minx, miny)
	maxx, maxy = viewMatrix.Apply(maxx, maxy)

	path := vector.Path{}
	path.MoveTo(float32(minx), float32(miny))
	path.LineTo(float32(maxx), float32(miny))
	path.LineTo(float32(maxx), float32(maxy))
	path.LineTo(float32(minx), float32(maxy))
	path.Close()

	vs, ic := path.AppendVerticesAndIndicesForStroke(nil, nil, &vector.StrokeOptions{
		Width: 1,
	})

	screen.DrawTriangles(vs, ic, rectImg, nil)

	vs, ic = path.AppendVerticesAndIndicesForFilling(nil, nil)

	for i := range vs {
		vs[i].ColorA = (float32(color.A) / 4) / 255.0
	}

	screen.DrawTriangles(vs, ic, rectImg, nil)
}
