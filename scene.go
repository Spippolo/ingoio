package main

import (
	"path/filepath"

	sdl_image "github.com/veandco/go-sdl2/sdl_image"
)

// Scene has all Scene objects
type Scene struct {
	cam *Camera
}

func (s *Scene) calculateFrame(elapsed int32) {
	s.cam.tick(elapsed)
}

func (s *Scene) drawFrame() {
	s.cam.draw()
}

func newScene() *Scene {
	s := new(Scene)
	sur, err := sdl_image.Load(filepath.Join(mediaPath, "imgs/background.png"))
	checkErr(err)

	s.cam = &Camera{}

	s.cam.tex, err = renderer.CreateTextureFromSurface(sur)
	checkErr(err)
	sur.Free()
	s.cam.speed = 10
	s.cam.backgroundWidth = 2400

	return s
}
