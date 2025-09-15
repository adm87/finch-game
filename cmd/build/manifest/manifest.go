package manifest

import (
	"path/filepath"

	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/fsys"
	"github.com/adm87/finch-resources/resources"
	"github.com/spf13/cobra"
)

func Command(ctx finch.Context, rootPath *string) *cobra.Command {
	var indent bool

	cmd := &cobra.Command{
		Use:   "manifest",
		Short: "Generate resource manifests",
		RunE: func(cmd *cobra.Command, args []string) error {
			resourcePath := ctx.Get("resource_path").(string)

			m := resources.GenerateManifest(ctx, resourcePath)
			if indent {
				return fsys.WriteJsonIndent(filepath.Join(resourcePath, resources.JsonName), m)
			}

			return fsys.WriteJson(filepath.Join(resourcePath, resources.JsonName), m)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.Flags().BoolVar(&indent, "indent", false, "Indent the JSON output")

	return cmd
}
