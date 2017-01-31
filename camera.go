package main

import "github.com/veandco/go-sdl2/sdl"

// Commento a caso
type Camera struct {
	tex             *sdl.Texture
	cX              int32
	backgroundWidth int32
	speed           int32
}

func (c *Camera) draw() {
	renderer.Copy(c.tex, &sdl.Rect{X: c.cX, W: int32(windowWidth), H: int32(windowHeight)}, nil)
}

func (c *Camera) tick(elapsed int32) {
	c.cX += c.speed * elapsed

	if c.cX+int32(windowWidth) >= c.backgroundWidth {
		c.cX = c.backgroundWidth - (c.cX + int32(windowWidth))
	}
}
