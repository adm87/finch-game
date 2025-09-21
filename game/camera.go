package game

import (
	"github.com/adm87/finch-core/geom"
	"github.com/hajimehoshi/ebiten/v2"
)

type Camera struct {
	X, Y             float64
	width, height    float64
	extentX, extentY float64
}

func NewCamera(width, height float64) *Camera {
	return &Camera{
		X:       0,
		Y:       0,
		width:   width,
		height:  height,
		extentX: width / 2,
		extentY: height / 2,
	}
}

func (c *Camera) Viewport() geom.Rect64 {
	return geom.NewRect64(c.X-(c.extentX), c.Y-(c.extentY), c.width, c.height)
}

func (c *Camera) ViewMatrix() ebiten.GeoM {
	m := ebiten.GeoM{}
	m.Translate(-c.extentX, -c.extentY)
	m.Translate(c.X, c.Y)
	m.Invert()
	return m
}
