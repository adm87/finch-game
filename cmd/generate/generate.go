package generate

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-game/cmd/generate/assets"
	"github.com/spf13/cobra"
)

func Generate(ctx finch.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "generate",
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.AddCommand(assets.AssetFiles(ctx))

	return cmd
}
