package camera

import (
	"github.com/adm87/finch-core/geom"
	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	X, Y          float64
	width, height float64
	zoom          float64
}

func NewCamera(width, height float64) *Camera {
	return &Camera{
		X:      0,
		Y:      0,
		width:  width,
		height: height,
		zoom:   1,
	}
}

// Viewport returns the rectangle representing the area of the world currently visible through the camera.
func (c *Camera) Viewport() geom.Rect64 {
	halfW := c.width / (2 * c.zoom)
	halfH := c.height / (2 * c.zoom)
	left := c.X - halfW
	top := c.Y - halfH
	return geom.NewRect64(left, top, 2*halfW, 2*halfH)
}

// ViewMatrix returns the transformation matrix to render the scene from the camera's perspective.
func (c *Camera) ViewMatrix() ebiten.GeoM {
	m := ebiten.GeoM{}
	m.Translate(-c.X, -c.Y)
	m.Scale(c.zoom, c.zoom)
	m.Translate(c.width/2, c.height/2)
	return m
}

func (c *Camera) Zoom() float64 {
	return c.zoom
}

func (c *Camera) SetZoom(zoom float64) {
	if zoom <= 0 {
		zoom = 0.01
	}
	c.zoom = zoom
}
