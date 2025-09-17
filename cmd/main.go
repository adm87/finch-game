package main

import (
	"os"
	"path"
	"path/filepath"

	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/fsys"
	"github.com/adm87/finch-game/cmd/build"
	"github.com/adm87/finch-game/data"
	"github.com/adm87/finch-game/game"
	"github.com/adm87/finch-game/module"
	"github.com/adm87/finch-resources/resources"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/cobra"
)

var version = "0.0.0-unreleased"

func main() {
	var (
		rootPath      string
		setFullscreen bool
	)

	f := finch.NewApp().
		WithWindow(&finch.Window{
			Title:        "Finch Game v" + version,
			ResizingMode: ebiten.WindowResizingModeDisabled,
			Width:        1280,
			Height:       720,
			Fullscreen:   setFullscreen,
			RenderScale:  0.5,
		}).
		WithUpdate(game.Update).
		WithDraw(game.Draw).
		WithStartup(game.Startup)

	cmd := &cobra.Command{
		Use:     "finch",
		Short:   "A sample game using Finch and Ebitengine",
		Version: version,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			module.Register(f.Context())

			rootPath = fsys.MakeAbsolute(filepath.Clean(rootPath))
			fsys.DirectoryMustExist(rootPath)

			resourcePath := path.Join(rootPath, "data")
			fsys.DirectoryMustExist(resourcePath)

			resources.LoadManifest(f.Context(), resourcePath)
			f.Context().Set("resource_path", resourcePath)

			resources.AddFilesystem("assets", os.DirFS(path.Join(resourcePath, "assets")))
			resources.AddFilesystem("embedded", data.EmbeddedFS)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return finch.Run(f)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.AddCommand(build.Command(f.Context(), &rootPath))

	cmd.PersistentFlags().StringVar(&rootPath, "root", rootPath, "Sets the root of the application")
	cmd.Flags().BoolVar(&setFullscreen, "fullscreen", setFullscreen, "Set to run in fullscreen mode")

	if err := cmd.ExecuteContext(f.Context().Context()); err != nil {
		panic(err)
	}
}
