package main

import (
	"fmt"
	"path"
	"path/filepath"

	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/fsys"
	"github.com/adm87/finch-game/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/spf13/cobra"
)

var (
	version       = "0.0.0-unreleased"
	rootPath      = "."
	setFullscreen = false
)

func main() {
	f := finch.NewApp().
		WithWindow(&finch.Window{
			Title:        fmt.Sprintf("Finch v%s - %s", version, BuildTag),
			ResizingMode: ebiten.WindowResizingModeDisabled,
			Width:        1170,
			Height:       675,
			Fullscreen:   setFullscreen,
			RenderScale:  0.4,
		}).
		WithUpdate(game.Update).
		WithFixedUpdate(game.FixedUpdate).
		WithLateUpdate(game.LateUpdate).
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

			// Note: The execution order of these configuration methods matter.
			RegisterModules(f.Context())
			RegisterAssetDirectories(f.Context(), resourcePath)
			MustLoad(f.Context())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return finch.Run(f)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	SetupCommand(cmd, f.Context())

	if err := cmd.ExecuteContext(f.Context().Context()); err != nil {
		panic(err)
	}
}
