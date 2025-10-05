package actor

import (
	"github.com/adm87/finch-collision/collision"
)

type Actor interface {
	Collider() collision.Collider
	Type() ActorType
}
