package main

import (
	"os"

	"fmt"

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
	scene        *Scene
	mediaPath    string
	freq         uint64
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

	window, err = sdl.CreateWindow("InGOio", 0, 0, windowWidth, windowHeight, sdl.WINDOW_OPENGL)
	checkErr(err)
	glContext, _ := sdl.GL_CreateContext(window)
	defer sdl.GL_DeleteContext(glContext)
	sdl.GL_SetSwapInterval(1)
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	checkErr(err)

	defer window.Destroy()

	drawTitle("InGO io")
	sdl.Delay(2000)
	scene = newScene()
	// freq = sdl.GetPerformanceFrequency()
	prevTime := sdl.GetTicks()
	var tmpTime uint32
	go func(fps uint32) {
		for {
			tmpTime = prevTime
			// prevTime = sdl.GetPerformanceCounter()
			prevTime = sdl.GetTicks()
			fmt.Println(tmpTime)
			fmt.Println(prevTime)
			delta := (prevTime - tmpTime)
			// fmt.Println(delta)
			fmt.Println("")
			scene.calculateFrame(int32(delta))
			scene.drawFrame()
			renderer.Present()
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
}
