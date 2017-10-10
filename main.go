package main

import (
	_ "image/png"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/nicklanng/pixel-play/engine"
	"github.com/nicklanng/pixel-play/game"
	"golang.org/x/image/colornames"
)

func loadSprites() (*engine.SpriteSheet, error) {
	spriteSheetConf := engine.SpriteSheetConfig{
		ImagePath:    "tileset.png",
		MetadataPath: "spritesheet.csv",
		TileWidth:    32,
		TileHeight:   32,
	}
	sprites, err := engine.CreateSpriteSheet(spriteSheetConf)
	if err != nil {
		return nil, err
	}

	return sprites, nil
}

func run() {
	var components []engine.Component

	cfg := pixelgl.WindowConfig{
		Title:       "Pixel-Play",
		Bounds:      pixel.R(0, 0, 1024, 768),
		Undecorated: false,
		Resizable:   false,
		VSync:       false,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	sprites, err := loadSprites()
	if err != nil {
		panic(err)
	}

	camera := engine.NewCamera(
		pixel.V(640, 480), // viewport size
		win.Bounds(),      // render bounds
		pixel.V(64, 64),   // initial position
		colornames.Black,  // clear color
	)

	worldManager := game.NewWorldManager(camera, sprites)
	components = append(components, worldManager)

	last := time.Now()
	for !win.Closed() {
		delta := time.Since(last)
		last = time.Now()

		for _, c := range components {
			c.Update(delta)
		}

		for _, c := range components {
			c.Draw()
		}

		camera.Draw(win)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
