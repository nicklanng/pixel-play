package engine

import (
	"errors"
	"fmt"
	"os"

	"github.com/faiface/pixel"
	"github.com/gocarina/gocsv"
)

type SpriteSheetConfig struct {
	ImagePath    string
	MetadataPath string
	TileWidth    float64
	TileHeight   float64
}

func CreateSpriteSheet(conf SpriteSheetConfig) (*SpriteSheet, error) {
	picture, err := loadPictureFromDisk(conf.ImagePath)
	if err != nil {
		return nil, err
	}

	clientsFile, err := os.OpenFile(conf.MetadataPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer clientsFile.Close()

	f := []*spriteSheetFrame{}
	if err := gocsv.UnmarshalFile(clientsFile, &f); err != nil {
		return nil, err
	}

	frames := make(map[string]*pixel.Sprite, len(f))
	for _, frame := range f {
		fmt.Println("Loaded sprite", frame.Name)
		minX := conf.TileWidth * float64(frame.Column)
		maxX := minX + conf.TileWidth
		minY := picture.Bounds().H() - conf.TileHeight*float64(frame.Row)
		maxY := minY - conf.TileHeight
		frames[frame.Name] = pixel.NewSprite(picture, pixel.R(minX, minY, maxX, maxY))
	}

	return &SpriteSheet{
		Picture: picture,
		frames:  frames,
	}, nil
}

type spriteSheetFrame struct {
	Name   string `csv:"name"`
	Row    int    `csv:"row"`
	Column int    `csv:"col"`
}

type SpriteSheet struct {
	Picture pixel.Picture
	frames  map[string]*pixel.Sprite
}

func (ss *SpriteSheet) GetFrame(name string) (*pixel.Sprite, error) {
	frame, ok := ss.frames[name]
	if !ok {
		return nil, errors.New("No frame with name: " + name)
	}

	return frame, nil
}
