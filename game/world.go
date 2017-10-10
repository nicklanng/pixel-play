package game

import (
	"time"

	"github.com/faiface/pixel"
	"github.com/nicklanng/pixel-play/engine"
)

type WorldManager struct {
	camera   *engine.Camera
	tilemap  *engine.TileMap
	entities []*engine.Entity
}

func NewWorldManager(c *engine.Camera, s *engine.SpriteSheet) *WorldManager {
	wm := new(WorldManager)

	wm.camera = c

	tilemap, err := engine.NewTileMap(s, 8, 8, 32, 32)
	if err != nil {
		panic(err)
	}

	wm.tilemap = tilemap

	primitiveSprite, err := s.GetFrame("c_primitive")
	if err != nil {
		panic(err)
	}
	primitive := engine.NewEntity(pixel.V(32, 32), primitiveSprite)
	wm.entities = append(wm.entities, primitive)

	hellknightSprite, err := s.GetFrame("c_hellknight")
	if err != nil {
		panic(err)
	}
	hellknight := engine.NewEntity(pixel.V(0, 0), hellknightSprite)
	wm.entities = append(wm.entities, hellknight)

	firegriffonSprite, err := s.GetFrame("c_firegriffon")
	if err != nil {
		panic(err)
	}
	firegriffon := engine.NewEntity(pixel.V(64, 64), firegriffonSprite)
	wm.entities = append(wm.entities, firegriffon)

	hobgoblinSprite, err := s.GetFrame("c_hobgoblin")
	if err != nil {
		panic(err)
	}
	hobgoblin := engine.NewEntity(pixel.V(96, 96), hobgoblinSprite)
	wm.entities = append(wm.entities, hobgoblin)

	return wm
}

func (wm *WorldManager) Update(delta time.Duration) {

}

func (wm *WorldManager) Draw() {
	wm.tilemap.Draw(wm.camera)
	for _, e := range wm.entities {
		e.Draw(wm.camera)
	}
}
