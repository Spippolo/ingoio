package main

import (
	"github.com/veandco/go-sdl2/sdl"
	sdl_ttf "github.com/veandco/go-sdl2/sdl_ttf"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

var (
	window       *sdl.Window
	renderer     *sdl.Renderer
	windowHeight = 600
	windowWidth  = 800
)

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	checkErr(err)
	err = sdl_ttf.Init()
	checkErr(err)

	window, renderer, err = sdl.CreateWindowAndRenderer(
		windowWidth, windowHeight, sdl.WINDOW_SHOWN)

	checkErr(err)
	defer window.Destroy()

	window.SetTitle("InGOio")
	drawTitle("InGO io")
	sdl.Delay(3000)
	// scene := newScene()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.WaitEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.KeyDownEvent:
				handleKeyEvent(event.(*sdl.KeyDownEvent))
			}
		}
	}

	sdl.Quit()
	// for {
	// 	handleEvents()
	// 	updateStates()
	// 	display()
	// }
}
