package actors

import (
	"github.com/adm87/finch-collision/colliders"
	"github.com/adm87/finch-collision/collision"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameActor struct {
	collision.Collider

	obj       *tiled.Object
	transform ebiten.GeoM
}

func NewGameActor() *GameActor {
	return &GameActor{
		Collider: colliders.NewBoxCollider(0, 0, 0, 0),
	}
}

func (ga *GameActor) Object() *tiled.Object {
	return ga.obj
}

func (ga *GameActor) SetObject(obj *tiled.Object) {
	ga.obj = obj
}

func (ga *GameActor) Transform() ebiten.GeoM {
	ga.transform.Reset()
	ga.transform.Translate(ga.Offset().X, ga.Offset().Y)
	ga.transform.Translate(ga.X(), ga.Y())
	return ga.transform
}
