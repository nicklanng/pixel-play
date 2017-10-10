package engine

import (
	"math/rand"
	"strconv"

	"github.com/faiface/pixel"
)

type Tile struct {
	texture  *pixel.Sprite
	position pixel.Vec
}

func (ti *Tile) Draw(t pixel.Target) {
	ti.texture.Draw(t, pixel.IM.Moved(ti.position))
}

type TileMap struct {
	width      int
	height     int
	tileWidth  int
	tileHeight int

	tiles []*Tile

	batch *pixel.Batch
}

func NewTileMap(spritesheet *SpriteSheet, width, height, tileWidth, tileHeight int) (*TileMap, error) {
	tilemap := new(TileMap)
	tilemap.width = width
	tilemap.height = height
	tilemap.tileWidth = tileWidth
	tilemap.tileHeight = tileHeight
	tilemap.batch = pixel.NewBatch(&pixel.TrianglesData{}, spritesheet.Picture)

	tilemap.tiles = make([]*Tile, width*height)

	for i := 0; i < width*height; i++ {
		frame, err := spritesheet.GetFrame("t_longgrass_" + strconv.Itoa(rand.Intn(6)))
		if err != nil {
			return nil, err
		}

		x := i % width
		y := i / width

		tilemap.tiles[i] = &Tile{
			texture:  frame,
			position: pixel.V(float64(x*tileWidth), float64(y*tileHeight)),
		}
	}

	return tilemap, nil
}

func (tm *TileMap) Draw(t pixel.Target) {
	tm.batch.Clear()

	for _, tile := range tm.tiles {
		tile.Draw(tm.batch)
	}

	tm.batch.Draw(t)
}
