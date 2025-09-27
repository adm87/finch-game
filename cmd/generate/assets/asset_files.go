package assets

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/tmpl"
	"github.com/spf13/cobra"
)

const assetFileTemplate = `package {{ .PackageName }}
{{ with .AssetPaths }}
import "github.com/adm87/finch-core/finch"

// Auto-generated file, do not edit manually.

var (
	{{- range $key, $value := . }}
	{{ $key | pascal }} = finch.AssetFile("{{ $value }}") // Path: {{ $value }}
	{{- end }}
)
{{ end }}
`

func AssetFiles(ctx finch.Context) *cobra.Command {
	var (
		pkgName string = "assets"
		output  string = ""
	)

	cmd := &cobra.Command{
		Use:           "asset-files",
		Short:         "Generates a list of asset files used in the game",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			root := ctx.Get("resource_path").(string)
			assetPaths := get_asset_paths(root)

			wrapper := struct {
				PackageName string
				AssetPaths  map[string]string
			}{
				PackageName: pkgName,
				AssetPaths:  assetPaths,
			}

			data := tmpl.Render(assetFileTemplate, assetFileTemplate, wrapper)
			file, err := os.Create(output)

			if err != nil {
				return err
			}

			defer file.Close()

			if _, err := file.Write(data); err != nil {
				return err
			}

			return exec.Command("go", "fmt", output).Run()
		},
	}

	cmd.Flags().StringVar(&pkgName, "package", pkgName, "Sets the package name for the generated file")
	cmd.Flags().StringVar(&output, "output", output, "Sets the output file")

	cmd.MarkFlagRequired("package")
	cmd.MarkFlagRequired("output")

	return cmd
}

func get_asset_paths(root string) map[string]string {
	assetPaths := make(map[string]string)

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			relPath, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}

			ext := filepath.Ext(info.Name())
			if !finch.HasAssetTypeSupport(finch.AssetType(ext[1:])) {
				return nil
			}
			name := strings.TrimSuffix(info.Name(), ext)

			if _, exists := assetPaths[name]; exists {
				return fmt.Errorf("duplicate asset name: %s (path: %s)", name, relPath)
			}

			assetPaths[name] = relPath
		}

		return nil
	})

	return assetPaths
}
