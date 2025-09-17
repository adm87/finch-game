package enums

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-resources/enums"
	"github.com/adm87/finch-resources/resources"
	"github.com/spf13/cobra"
)

func Command(ctx finch.Context) *cobra.Command {
	var (
		dest string
		pkg  string = "enums"
	)

	cmd := &cobra.Command{
		Use:   "enums",
		Short: "Generates enums for accessing resources from the manifest",
		Long:  "Generates an enums.go file that contains an enum for each resource in the manifest.",
		Run: func(cmd *cobra.Command, args []string) {
			enums.Generate(ctx, resources.GetManifest(), pkg, dest)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.Flags().StringVarP(&dest, "dest", "d", dest, "Output directory to put the generated enums.go")
	cmd.Flags().StringVarP(&pkg, "package", "p", pkg, "Package name to use for the generated enums.go")

	cmd.MarkFlagRequired("dest")

	return cmd
}
