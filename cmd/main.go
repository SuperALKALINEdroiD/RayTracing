package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

func main() {
	var Speed int32 = 5

	// Initialize SDL
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	// Create Window
	window, err := sdl.CreateWindow("Ray Tracing Test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, ScreenWidth, ScreenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	window.SetResizable(false)
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	rect := sdl.Rect{X: 20, Y: 20, W: 100, H: 200}

	running := true

	for running {
		// Event handling
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		if rect.X+rect.W >= ScreenWidth || rect.X <= 0 {
			Speed = -Speed
		}
		rect.X += Speed

		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.DrawPoint(100, 100)
		renderer.FillRect(&rect)

		renderer.Present()

		sdl.Delay(16)
	}
}
