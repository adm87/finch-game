package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-game/data"
)

func Startup(ctx finch.Context) {
	finch.MustLoadAssets(
		data.TilemapCharactersPacked,
		data.TilemapExampleA,
		data.TilemapExampleB,
		data.TilemapPacked,
		data.TilesetCharacters,
		data.TilesetTiles,
	)
}
