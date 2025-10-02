package tiled

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-game/cmd/tiled/imports"
	"github.com/spf13/cobra"
)

func Tiled(ctx finch.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tiled",
		Short: "Commands for working with Tiled Projects",
	}

	cmd.AddCommand(imports.Configs(ctx))
	return cmd
}
