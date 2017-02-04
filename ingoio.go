package main

import (
	"github.com/veandco/go-sdl2/sdl"
	sdl_ttf "github.com/veandco/go-sdl2/sdl_ttf"
	"os"
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
	scene        *Scene
	mediaPath    string
)

func init() {
	mediaPath = os.Getenv("INGOIO_MEDIA")
	if mediaPath == "" {
		mediaPath = "./res/"
	}
}

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	checkErr(err)
	err = sdl_ttf.Init()
	checkErr(err)

	window, err = sdl.CreateWindow("InGOio",
		0, 0, windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	checkErr(err)
	renderer, err = sdl.CreateRenderer(window, -1,
		sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	checkErr(err)
	defer window.Destroy()

	window.SetTitle("InGOio")
	drawTitle("InGO io")
	sdl.Delay(3000)
	scene = newScene()

	go func(fps uint32) {
		for {
			scene.drawFrame()
			renderer.Present()
			sdl.Delay(1000 / fps)
		}
	}(30)

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
