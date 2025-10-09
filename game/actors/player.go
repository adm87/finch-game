package actors

import (
	"github.com/adm87/finch-collision/collision"
	"github.com/adm87/finch-core/enum"
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/fsys"
	"github.com/adm87/finch-game/actor"
	"github.com/adm87/finch-tiled/tiled"
)

type Player struct {
	*GameActor
}

func NewPlayer() *Player {
	return &Player{
		GameActor: NewGameActor(),
	}
}

var PlayerTiledObjectFactory = tiled.TiledObjectFactory[actor.Actor]{
	FromTemplate: PlayerFromTemplate,
	FromObject:   PlayerFromObject,
}

func PlayerFromTemplate(instance *tiled.Object, template *tiled.TX, tmx *tiled.TMX) actor.Actor {
	player := PlayerFromObject(template.Object, tmx)
	player.SetObject(instance)

	if tileset, _ := tiled.GetTSX(finch.AssetFile(template.Tileset.Source())); tileset != nil {
		if alignAttr, ok := tileset.Attrs[tiled.ObjectAlignmentAttr]; ok {
			x, y := getTiledAlignmentOffset(fsys.MustGet(enum.Value[tiled.ObjectAlignment](alignAttr.String())))
			x *= player.Width()
			y *= player.Height()
			player.SetOffset(x+float64(tileset.TileOffset.X()), y+float64(tileset.TileOffset.Y()))
		}
	}

	if x := instance.X(); x != 0 {
		player.SetX(float64(x))
	}
	if y := instance.Y(); y != 0 {
		player.SetY(float64(y))
	}
	if w := instance.Width(); w != 0 {
		player.SetWidth(float64(w))
	}
	if h := instance.Height(); h != 0 {
		player.SetHeight(float64(h))
	}

	return player
}

func PlayerFromObject(obj *tiled.Object, tmx *tiled.TMX) actor.Actor {
	player := NewPlayer()

	player.SetX(float64(obj.X()))
	player.SetY(float64(obj.Y()))
	player.SetWidth(float64(obj.Width()))
	player.SetHeight(float64(obj.Height()))
	player.SetObject(obj)

	if colliderProps, ok := obj.PropertyOfType("BoxCollider"); ok {
		if detection, ok := colliderProps.PropertyOfType("CollisionDetectionType"); ok {
			player.Collider.SetDetectionType(fsys.MustGet(enum.Value[collision.CollisionDetectionType](detection.Value())))
		}
		if layer, ok := colliderProps.PropertyOfType("CollisionLayer"); ok {
			player.Collider.SetCollisionLayer(fsys.MustGet(enum.Value[collision.CollisionLayer](layer.Value())))
		}
		if ctype, ok := colliderProps.PropertyOfType("ColliderType"); ok {
			player.Collider.SetColliderType(fsys.MustGet(enum.Value[collision.ColliderType](ctype.Value())))
		}
	}

	return player.GameActor
}

func getTiledAlignmentOffset(align tiled.ObjectAlignment) (x, y float64) {
	switch align {
	case tiled.ObjectAlignmentUnspecified, tiled.ObjectAlignmentTopLeft:
		return 0, 0
	case tiled.ObjectAlignmentTop:
		return -0.5, 0
	case tiled.ObjectAlignmentTopRight:
		return -1, 0
	case tiled.ObjectAlignmentLeft:
		return 0, -0.5
	case tiled.ObjectAlignmentCenter:
		return -0.5, -0.5
	case tiled.ObjectAlignmentRight:
		return -1, -0.5
	case tiled.ObjectAlignmentBottomLeft:
		return 0, -1
	case tiled.ObjectAlignmentBottom:
		return -0.5, -1
	case tiled.ObjectAlignmentBottomRight:
		return -1, -1
	default:
		return 0, 0
	}
}
