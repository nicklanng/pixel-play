package engine

import (
	"github.com/faiface/pixel"
)

func NewEntity(worldPosition pixel.Vec, sprite *pixel.Sprite) (e *Entity) {
	e = new(Entity)
	e.WorldPosition = worldPosition
	e.Sprite = sprite

	return
}

type Entity struct {
	WorldPosition pixel.Vec
	Sprite        *pixel.Sprite
}

func (e *Entity) Draw(t pixel.Target) {
	mat := pixel.IM
	mat = mat.Moved(e.WorldPosition)
	e.Sprite.Draw(t, mat)
}
