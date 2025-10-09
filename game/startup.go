package game

import (
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-game/actor"
	"github.com/adm87/finch-game/data"
	"github.com/adm87/finch-game/game/actors"
	"github.com/adm87/finch-game/scene"
	"github.com/adm87/finch-tiled/tiled"
)

var gameScene *scene.Scene

var (
	debugDrawTiled         = true
	debugDrawColliders     = false
	debugDrawCollisionGrid = false
	debugDrawSceneTree     = false
)

func Startup(ctx finch.Context) {
	finch.MustLoadAssets(
		data.BoxCollider,
		data.Player,
		data.TilemapCharactersPacked,
		data.TilemapExampleA,
		data.TilemapPacked,
		data.TilesetCharacters,
		data.TilesetTiles,
	)
	gameScene = scene.New(ctx, tiled.MustGetTMX(data.TilemapExampleA))
	gameScene.Factories().
		AddActorFactory(actor.ActorTypePlayer, actors.PlayerTiledObjectFactory)
	gameScene.Build()
	gameScene.Start()
}
