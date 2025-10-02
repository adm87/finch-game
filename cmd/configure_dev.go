//go:build development

package main

import (
	"os"
	"path"

	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-game/cmd/generate"
	"github.com/adm87/finch-game/cmd/tiled"
	"github.com/adm87/finch-game/data"
	"github.com/spf13/cobra"

	tiledproject "github.com/adm87/finch-tiled/project"
)

const BuildTag = "Development"

func RegisterModules(ctx finch.Context) {
	RegisterModulesShared(ctx)

	tiledproject.RegisterAssetImporter()
}

func RegisterAssetDirectories(ctx finch.Context, resourcePath string) {
	RegisterAssetDirectoriesShared(ctx, resourcePath)

	finch.RegisterAssetFilesystem(finch.AssetRoot("project"), os.DirFS(path.Join(resourcePath, "project")))
}

func SetupCommand(cmd *cobra.Command, ctx finch.Context) {
	SetupCommandShared(cmd, ctx)

	cmd.AddCommand(
		generate.Generate(ctx),
		tiled.Tiled(ctx),
	)
}

func MustLoad(ctx finch.Context) {
	MustLoadShared(ctx)

	finch.MustLoadAssets(data.Finch)
}
