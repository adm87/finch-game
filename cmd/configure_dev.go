//go:build development

package main

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-game/cmd/generate"
	"github.com/spf13/cobra"
)

func addSubCommands(cmd *cobra.Command, ctx finch.Context) {
	cmd.AddCommand(generate.Generate(ctx))
}

func addPersistentFlags(cmd *cobra.Command, ctx finch.Context) {
	cmd.PersistentFlags().StringVar(&rootPath, "root", ".", "Set the project root path")
}

func addFlags(cmd *cobra.Command, ctx finch.Context) {
	cmd.Flags().BoolVar(&setFullscreen, "fullscreen", false, "Set fullscreen mode")
}
