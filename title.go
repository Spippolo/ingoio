package main

import (
	"github.com/veandco/go-sdl2/sdl"
	sdl_ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

func drawTitle(title string) {
	font, err := sdl_ttf.OpenFont("./Roboto/Roboto-Bold.ttf", 14)
	checkErr(err)
	defer font.Close()

	txtColor := sdl.Color{R: 66, G: 0, B: 0, A: 255}
	surface, err := font.RenderUTF8_Solid(title, txtColor)
	checkErr(err)
	defer surface.Free()

	tex, err := renderer.CreateTextureFromSurface(surface)
	checkErr(err)

	renderer.Copy(tex, nil, &sdl.Rect{X: int32(windowWidth / 2), Y: int32(windowHeight / 2), W: 200, H: 100})
	renderer.Present()
}
