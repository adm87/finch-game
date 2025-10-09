package scene

import "github.com/adm87/finch-core/enum"

type SceneLayer int

const (
	SceneCollisionLayer SceneLayer = iota
	SceneActorLayer
)

func (sl SceneLayer) String() string {
	switch sl {
	case SceneCollisionLayer:
		return "Collision"
	case SceneActorLayer:
		return "Actor"
	default:
		return "Unknown"
	}
}

func (sl SceneLayer) IsValid() bool {
	return sl >= SceneCollisionLayer && sl <= SceneActorLayer
}

func (sl SceneLayer) MarshalJSON() ([]byte, error) {
	return enum.MarshalEnum(sl)
}

func (sl *SceneLayer) UnmarshalJSON(data []byte) error {
	val, err := enum.UnmarshalEnum[SceneLayer](data)
	if err != nil {
		return err
	}
	*sl = val
	return nil
}
