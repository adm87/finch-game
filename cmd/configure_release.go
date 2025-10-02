//go:build release

package main

import (
	"github.com/adm87/finch-core/finch"
	"github.com/spf13/cobra"
)

const BuildTag = "Release"

func RegisterModules(ctx finch.Context) {
	RegisterModulesShared(ctx)
}

func RegisterAssetDirectories(ctx finch.Context, resourcePath string) {
	RegisterAssetDirectoriesShared(ctx, resourcePath)
}

func SetupCommand(cmd *cobra.Command, ctx finch.Context) {
	SetupCommandShared(cmd, ctx)
}

func MustLoad(ctx finch.Context) {
	MustLoadShared(ctx)
}
