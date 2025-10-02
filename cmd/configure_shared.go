package main

import (
	"os"
	"path"

	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/spf13/cobra"
)

func RegisterModulesShared(ctx finch.Context) {
	finch.RegisterImageAssetTypes()
	tiled.RegisterTiledAssetImporter()
}

func RegisterAssetDirectoriesShared(ctx finch.Context, resourcePath string) {
	finch.RegisterAssetFilesystem(finch.AssetRoot("assets"), os.DirFS(path.Join(resourcePath, "assets")))
}

func SetupCommandShared(cmd *cobra.Command, ctx finch.Context) {
	cmd.PersistentFlags().StringVar(&rootPath, "root", ".", "Set the project root path")
	cmd.Flags().BoolVar(&setFullscreen, "fullscreen", false, "Set fullscreen mode")
}

func MustLoadShared(ctx finch.Context) {
	// TASK: Load embedded assets here
}
