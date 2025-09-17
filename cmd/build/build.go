package build

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-game/cmd/build/manifest"
	"github.com/spf13/cobra"
)

func Command(ctx finch.Context, rootPath *string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "build",
		Short: "Build the game for distribution",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.AddCommand(manifest.Command(ctx, rootPath))

	return cmd
}
