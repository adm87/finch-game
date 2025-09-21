package game

import (
	"image/color"

	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-game/data"
	"github.com/adm87/finch-resources/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

var camera *Camera

var renderTarget *ebiten.Image
var cameraOutline *ebiten.Image
var imageOutline *ebiten.Image

var selectedMap resources.ResourceHandle = data.TilemapExampleA

func Startup(ctx finch.Context) {
	width := float64(ctx.Screen().Width()) / 2.0
	height := float64(ctx.Screen().Height()) / 2.0

	camera = NewCamera(width, height)
	renderTarget = ebiten.NewImage(int(width), int(height))

	cameraOutline = ebiten.NewImage(3, 3)
	imageOutline = ebiten.NewImage(3, 3)

	cameraOutline.Fill(color.White)
	imageOutline.Fill(color.White)

	resources.Load(ctx,
		data.TilemapExampleA,
		data.TilemapExampleB,
	)
}
