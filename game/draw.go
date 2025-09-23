package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
)

func Draw(ctx finch.Context, screen *ebiten.Image) {
	tiled.DrawScene(ctx, screen, selectedMap, camera.Viewport(), camera.ViewMatrix())
}

// func draw_rect(screen *ebiten.Image, rect geom.Rect64, color color.RGBA) {
// 	path := vector.Path{}

// 	minx, miny := rect.Min()
// 	maxx, maxy := rect.Max()

// 	path.MoveTo(float32(minx), float32(miny))
// 	path.LineTo(float32(maxx), float32(miny))
// 	path.LineTo(float32(maxx), float32(maxy))
// 	path.LineTo(float32(minx), float32(maxy))
// 	path.Close()

// 	vs, is := path.AppendVerticesAndIndicesForStroke(nil, nil, &vector.StrokeOptions{
// 		Width: 2,
// 	})

// 	for i := range vs {
// 		vs[i].ColorR = float32(color.R) / 255
// 		vs[i].ColorG = float32(color.G) / 255
// 		vs[i].ColorB = float32(color.B) / 255
// 		vs[i].ColorA = float32(color.A) / 255
// 	}

// 	screen.DrawTriangles(vs, is, rectImg, nil)
// }
