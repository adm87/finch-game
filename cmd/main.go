package main

import (
	"os"
	"path"
	"path/filepath"

	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/fsys"
	"github.com/adm87/finch-game/game"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/cobra"
)

var (
	version       = "0.0.0-unreleased"
	rootPath      = "."
	setFullscreen = false
)

func main() {
	finch.RegisterImageAssetManager()
	tiled.RegisterTMXAssetManager()
	tiled.RegisterTSXAssetManager()

	f := finch.NewApp().
		WithWindow(&finch.Window{
			Title:        "Finch Game v" + version,
			ResizingMode: ebiten.WindowResizingModeDisabled,
			Width:        1170,
			Height:       675,
			Fullscreen:   setFullscreen,
			RenderScale:  0.4,
		}).
		WithUpdate(game.Update).
		WithFixedUpdate(game.FixedUpdate).
		WithDraw(game.Draw).
		WithStartup(game.Startup)

	cmd := &cobra.Command{
		Use:     "finch",
		Short:   "A sample game using Finch and Ebitengine",
		Version: version,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			rootPath = fsys.MakeAbsolute(filepath.Clean(rootPath))
			fsys.DirectoryMustExist(rootPath)

			resourcePath := path.Join(rootPath, "data")
			fsys.DirectoryMustExist(resourcePath)

			f.Context().Set("resource_path", resourcePath)

			finch.RegisterAssetFilesystem(finch.AssetRoot("assets"), os.DirFS(path.Join(resourcePath, "assets")))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return finch.Run(f)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	addSubCommands(cmd, f.Context())
	addPersistentFlags(cmd, f.Context())
	addFlags(cmd, f.Context())

	if err := cmd.ExecuteContext(f.Context().Context()); err != nil {
		panic(err)
	}
}
