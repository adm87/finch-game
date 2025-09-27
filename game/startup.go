package game

import (
	"github.com/adm87/finch-core/finch"
)

var (
	tilemapPackedImg           = finch.AssetFile("assets/images/tilemap_packed.png")
	tilemapCharactersPackedImg = finch.AssetFile("assets/images/tilemap-characters_packed.png")
	tilemapExampleATmx         = finch.AssetFile("assets/tilemaps/tilemap-example-a.tmx")
	tilemapExampleBTmx         = finch.AssetFile("assets/tilemaps/tilemap-example-b.tmx")
	tilesetTilesTsx            = finch.AssetFile("assets/tilesets/tileset-tiles.tsx")
	tilesetCharactersTsx       = finch.AssetFile("assets/tilesets/tileset-characters.tsx")
)

func Startup(ctx finch.Context) {
	finch.MustLoadAssets(
		tilemapPackedImg,
		tilemapCharactersPackedImg,
		tilemapExampleATmx,
		tilemapExampleBTmx,
		tilesetTilesTsx,
		tilesetCharactersTsx,
	)
}
