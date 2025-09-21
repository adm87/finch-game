package game

import (
	"github.com/adm87/finch-core/geom"
	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	X, Y                  float64
	width, height         float64
	halfWidth, halfHeight float64
}

func NewCamera(width, height float64) *Camera {
	return &Camera{
		X:          0,
		Y:          0,
		width:      width,
		height:     height,
		halfWidth:  width / 2,
		halfHeight: height / 2,
	}
}

// Viewport returns the rectangle representing the area of the world currently visible through the camera.
func (c *Camera) Viewport() geom.Rect64 {
	return geom.NewRect64(c.X-(c.halfWidth), c.Y-(c.halfHeight), c.width, c.height)
}

// ViewMatrix returns the transformation matrix to be applied to render the scene from the camera's perspective.
func (c *Camera) ViewMatrix() ebiten.GeoM {
	m := ebiten.GeoM{}
	m.Translate(-c.halfWidth, -c.halfHeight)
	m.Translate(c.X, c.Y)
	m.Invert()
	return m
}
