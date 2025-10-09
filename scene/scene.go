package scene

import (
	"image/color"

	"github.com/adm87/finch-collision/collision"
	"github.com/adm87/finch-core/finch"
	"github.com/adm87/finch-core/hashset"
	"github.com/adm87/finch-core/partition/hashgrid"
	"github.com/adm87/finch-core/partition/quadtree"
	"github.com/adm87/finch-game/actor"
	"github.com/adm87/finch-game/camera"
	"github.com/adm87/finch-tiled/tiled"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	debugCollision "github.com/adm87/finch-collision/debug"
)

type Scene struct {
	ctx finch.Context

	camera         *camera.Camera
	collisionWorld *collision.CollisionWorld
	sceneTree      *quadtree.QuadTree[actor.Actor]
	tiledMap       *tiled.TMX

	factories *SceneFactories
}

func New(ctx finch.Context, tiledMap *tiled.TMX) *Scene {
	return &Scene{
		ctx: ctx,
		camera: camera.NewCamera(
			float64(ctx.Screen().Width()),
			float64(ctx.Screen().Height()),
		),
		collisionWorld: collision.NewCollisionWorld(
			hashgrid.New[collision.Collider](
				float64(tiledMap.TileWidth()+tiledMap.TileHeight()) / 4),
		),
		sceneTree: quadtree.New[actor.Actor](
			tiledMap.Bounds(),
			4,
			6,
		),
		tiledMap: tiledMap,
		factories: &SceneFactories{
			actors: make(map[actor.ActorType]tiled.TiledObjectFactory[actor.Actor]),
		},
	}
}

func (s *Scene) Factories() *SceneFactories {
	return s.factories
}

func (s *Scene) Camera() *camera.Camera {
	return s.camera
}

func (s *Scene) Start() {
	bounds := s.tiledMap.Bounds()
	s.camera.X = bounds.X + bounds.Width/2
	s.camera.Y = bounds.Y + bounds.Height/2
}

func (s *Scene) UpdateActor(a actor.Actor) {
	s.collisionWorld.AddCollider(a)
	s.sceneTree.Update(a)
}

func (s *Scene) Update(dt float64) {

}

func (s *Scene) FixedUpdate(dt float64) {

}

func (s *Scene) LateUpdate(dt float64) {

}

func (s *Scene) Draw(screen *ebiten.Image) {
	viewport := s.camera.Viewport()
	view := s.camera.ViewMatrix()

	tiled.DrawScene(s.ctx, screen, s.tiledMap, viewport, view)

	drawables := s.sceneTree.Query(viewport)
	for drawable := range drawables {
		tiled.DrawObject(s.ctx, screen, s.tiledMap, drawable.Object(), drawable.Transform(), view)
	}

}

func (s *Scene) Build() {
	staticCollision, actors := buildSceneComponents(s.ctx, s.tiledMap, s.factories)
	for collider := range staticCollision {
		s.collisionWorld.AddCollider(collider)
	}
	for a := range actors {
		s.sceneTree.Insert(a)
		s.collisionWorld.AddCollider(a)
	}
}

func (s *Scene) DebugDrawColliders(screen *ebiten.Image) {
	debugCollision.DrawColliders(s.collisionWorld, screen, s.camera.Viewport(), s.camera.ViewMatrix())
}

func (s *Scene) DebugDrawCollisionGrid(screen *ebiten.Image) {
	debugCollision.DrawCollisionGrid(s.collisionWorld, screen, s.camera.Viewport(), s.camera.ViewMatrix(), color.RGBA{G: 255, A: 64})
}

func (s *Scene) DebugDrawSceneTree(screen *ebiten.Image) {
	nodes := s.sceneTree.QueryNodes(s.camera.Viewport())
	s.drawRects(screen, nodes, color.RGBA{R: 255, A: 64})
}

func (s *Scene) drawRects(screen *ebiten.Image, nodes hashset.Set[*quadtree.QuadTree[actor.Actor]], color color.RGBA) {
	debugRectImg.Fill(color)

	path := vector.Path{}
	view := s.camera.ViewMatrix()

	for node := range nodes {
		minX, minY := node.Min()
		maxX, maxY := node.Max()
		sminX, sminY := view.Apply(minX, minY)
		smaxX, smaxY := view.Apply(maxX, maxY)

		path.MoveTo(float32(sminX), float32(sminY))
		path.LineTo(float32(smaxX), float32(sminY))
		path.LineTo(float32(smaxX), float32(smaxY))
		path.LineTo(float32(sminX), float32(smaxY))
		path.Close()
	}

	vs, is := path.AppendVerticesAndIndicesForStroke(nil, nil, &vector.StrokeOptions{
		Width: 2,
	})

	screen.DrawTriangles(vs, is, debugRectImg, nil)
}

var debugRectImg = ebiten.NewImage(3, 3)
