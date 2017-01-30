package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

func handleKeyEvent(event *sdl.KeyDownEvent) {
	fmt.Println("pressed")
}
