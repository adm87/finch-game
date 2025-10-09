package game

import "github.com/adm87/finch-collision/collision"

var (
	PlayerCollisionLayer = collision.NewCollisionLayer("Player")
	StaticCollisionLayer = collision.NewCollisionLayer("Static")
)
