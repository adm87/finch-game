package actor

import "github.com/adm87/finch-core/enum"

type ActorType int

const (
	ActorTypePlayer ActorType = iota
	ActorTypeEnemy
	ActorTypePickup
)

func (at ActorType) String() string {
	switch at {
	case ActorTypePlayer:
		return "Player"
	case ActorTypeEnemy:
		return "Enemy"
	case ActorTypePickup:
		return "Pickup"
	default:
		return "Unknown"
	}
}

func (at ActorType) IsValid() bool {
	return at >= ActorTypePlayer && at <= ActorTypePickup
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
