package engine

import (
	"image/color"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func NewCamera(size pixel.Vec, bounds pixel.Rect, position pixel.Vec, clearColor color.Color) *Camera {
	c := new(Camera)

	c.clearColor = clearColor
	c.size = size
	c.bounds = bounds
	c.canvas = pixelgl.NewCanvas(pixel.R(0, 0, size.X, size.Y))
	c.SetPosition(position)

	return c
}

type Camera struct {
	canvas *pixelgl.Canvas

	clearColor color.Color
	size       pixel.Vec
	position   pixel.Vec
	bounds     pixel.Rect
}

func (c *Camera) Draw(t pixel.Target) {
	mat := pixel.IM.Scaled(pixel.ZV,
		math.Min(
			c.bounds.W()/c.size.X,
			c.bounds.H()/c.size.Y,
		),
	).Moved(c.bounds.Center())

	c.canvas.Draw(t, mat)

	if c.clearColor != nil {
		c.canvas.Clear(c.clearColor)
	}
}

func (c *Camera) MakeTriangles(t pixel.Triangles) pixel.TargetTriangles {
	return c.canvas.MakeTriangles(t)
}

func (c *Camera) MakePicture(p pixel.Picture) pixel.TargetPicture {
	return c.canvas.MakePicture(p)
}

func (c *Camera) SetPosition(position pixel.Vec) {
	c.position = position

	camMatrix := pixel.IM.Moved(position.Scaled(-1)).Moved(c.size.Scaled(0.5))
	c.canvas.SetMatrix(camMatrix)
}
