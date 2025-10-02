package imports

import (
	"path/filepath"

	"github.com/adm87/finch-collision/collision"
	"github.com/adm87/finch-core/enum"
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/fsys"
	"github.com/adm87/finch-game/data"
	"github.com/adm87/finch-tiled/project"
	"github.com/spf13/cobra"
)

func Configs(ctx finch.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import-configs",
		Short: "Commands for working with Tiled Project Configurations",
		RunE: func(cmd *cobra.Command, args []string) error {
			tiledProj := finch.MustGetAsset[*project.TiledProject](data.Finch)
			if err := importCollisionEnums(tiledProj); err != nil {
				return err
			}
			// TASK: Since we're using Finch's asset framework to load and get data, it should also be used to save it.
			//       This will require some changes to the framework to allow writing to the filesystem.
			resourcePath := ctx.Get("resource_path").(string)
			return fsys.WriteJsonIndent(filepath.Join(resourcePath, data.Finch.Path()), tiledProj)
		},
	}
	return cmd
}

func importCollisionEnums(proj *project.TiledProject) error {
	return project.InsertOrUpdateEnumType(proj,
		project.TiledEnumPropertyType{
			TiledPropertyType: project.TiledPropertyType{
				Name: "ColliderType",
				Type: "enum",
			},
			StorageType:   "string",
			Values:        enum.Names[collision.ColliderType](),
			ValuesAsFlags: false,
		},
		project.TiledEnumPropertyType{
			TiledPropertyType: project.TiledPropertyType{
				Name: "CollisionDetectionType",
				Type: "enum",
			},
			StorageType:   "string",
			Values:        enum.Names[collision.CollisionDetectionType](),
			ValuesAsFlags: false,
		},
		project.TiledEnumPropertyType{
			TiledPropertyType: project.TiledPropertyType{
				Name: "CollisionLayer",
				Type: "enum",
			},
			StorageType:   "string",
			Values:        enum.Names[collision.CollisionLayer](),
			ValuesAsFlags: false,
		},
	)
}
