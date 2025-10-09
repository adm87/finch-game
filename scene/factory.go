package scene

import (
	"github.com/adm87/finch-game/actor"
	"github.com/adm87/finch-tiled/tiled"
)

type SceneFactories struct {
	actors map[actor.ActorType]tiled.TiledObjectFactory[actor.Actor]
}

func (sf *SceneFactories) AddActorFactory(at actor.ActorType, factory tiled.TiledObjectFactory[actor.Actor]) {
	if !at.IsValid() {
		panic("invalid actor type")
	}
	if _, exists := sf.actors[at]; exists {
		panic("actor factory already registered for type: " + at.String())
	}
	sf.actors[at] = factory
}
