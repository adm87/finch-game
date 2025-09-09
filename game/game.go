package game

import (
	"image/color"
	"io/fs"
	"os"
	"path"

	"github.com/adm87/finch-application/application"
	"github.com/adm87/finch-application/config"
	"github.com/adm87/finch-game/module"
	"github.com/adm87/finch-resources/images"
	"github.com/adm87/finch-resources/resources"
	"github.com/hajimehoshi/ebiten/v2"
)

var Game = application.NewApplicationWithConfig(
	&application.ApplicationConfig{
		Metadata: &config.Metadata{
			Name:      "Finch Game",
			Root:      ".",
			TargetFps: 60,
		},
		Window: &config.Window{
			Title:           "Finch Game",
			Width:           1240,
			Height:          720,
			RenderScale:     0.35,
			ClearColor:      color.RGBA{R: 100, G: 149, B: 237, A: 255}, // Cornflower Blue
			ClearBackground: true,
		},
		Resources: &config.Resources{
			Path: "data/",
		},
	}).
	WithRegistration(Register).
	WithStartup(Startup).
	WithShutdown(Shutdown).
	WithUpdate(Update).
	WithDraw(Draw)

func Register(app *application.Application) error {
	if err := module.RegisterModule(); err != nil {
		return err
	}
	if err := resources.RegisterFileSystems(map[string]fs.FS{
		"assets": os.DirFS(path.Join(app.Config().Resources.Path, "assets")),
	}); err != nil {
		return err
	}
	return nil
}

func Startup(app *application.Application) error {
	return resources.Load(
		"tileset_image_0000",
		"tileset_0000",
		"tilemap_0000",
	)
}

func Shutdown(app *application.Application) error {
	return nil
}

func Draw(app *application.Application, screen *ebiten.Image) error {
	img, err := images.Resources().Get("tileset_image_0000")
	if err != nil {
		return err
	}
	screen.DrawImage(img, &ebiten.DrawImageOptions{})
	return nil
}

func Update(app *application.Application, deltaSeconds, fixedDeltaSeconds float64, frames int) error {
	return nil
}
