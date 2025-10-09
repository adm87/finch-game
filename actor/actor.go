package actor

import (
	"github.com/adm87/finch-collision/collision"
	"github.com/adm87/finch-core/enum"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
)

type ActorType int

const (
	ActorTypeUnknown ActorType = iota
	ActorTypePlayer
)

func (at ActorType) String() string {
	switch at {
	case ActorTypeUnknown:
		return "Unknown"
	case ActorTypePlayer:
		return "Player"
	default:
		return "Unknown"
	}
}

func (at ActorType) IsValid() bool {
	return at >= ActorTypeUnknown && at <= ActorTypePlayer
}

func (at ActorType) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnum(at)
}

func (at *ActorType) UnmarshalJSON(data []byte) error {
	val, err := enum.UnmarshalEnum[ActorType](data)
	if err != nil {
		return err
	}
	*at = val
	return nil
}

type Actor interface {
	collision.Collider
	Drawable
}

type Drawable interface {
	Transform() ebiten.GeoM

	Object() *tiled.Object
	SetObject(obj *tiled.Object)
}
