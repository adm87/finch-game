package properties

import (
	"github.com/adm87/finch-collision/collision"
	"github.com/adm87/finch-core/enum"
	"github.com/adm87/finch-tiled/tiled"
)

// ======================================================
// Collision Properties
// ======================================================

type CollisionProperties struct {
	ColliderType   collision.ColliderType
	CollisionLayer collision.CollisionLayer
}

func NewCollisionProperties() *CollisionProperties {
	return &CollisionProperties{}
}

func (c *CollisionProperties) FromTMXProperty(props *tiled.Property) (*CollisionProperties, error) {
	if prop, ok := props.PropertyOfType("ColliderType"); ok {
		if val, ok := prop.Attrs[tiled.ValueAttr]; ok {
			colliderType, err := enum.Value[collision.ColliderType](val.String())
			if err != nil {
				return nil, err
			}
			c.ColliderType = colliderType
		}
	}

	if prop, ok := props.PropertyOfType("CollisionLayer"); ok {
		if val, ok := prop.Attrs[tiled.ValueAttr]; ok {
			collisionLayer, err := enum.Value[collision.CollisionLayer](val.String())
			if err != nil {
				return nil, err
			}
			c.CollisionLayer = collisionLayer
		}
	}

	return c, nil
}

// ======================================================
// Collision Behaviours
// ======================================================

type CollisionBehaviour int

const (
	CollisionBehaviourNone CollisionBehaviour = iota
)

func (cb CollisionBehaviour) String() string {
	switch cb {
	case CollisionBehaviourNone:
		return "None"
	default:
		return "Unknown"
	}
}

func (cb CollisionBehaviour) IsValid() bool {
	return cb >= CollisionBehaviourNone && cb <= CollisionBehaviourNone
}

func (cb CollisionBehaviour) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnum(cb)
}

func (cb *CollisionBehaviour) UnmarshalJSON(data []byte) error {
	val, err := enum.UnmarshalEnum[CollisionBehaviour](data)
	if err != nil {
		return err
	}
	*cb = val
	return nil
}
