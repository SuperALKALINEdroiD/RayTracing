package main

import (
	"fmt"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
	NumBalls     = 10 // Number of bouncing balls
)

type Ball struct {
	X, Y   float64
	DX, DY float64
	Radius int32
	Color  sdl.Color
	Mass   int32
}

var balls []Ball

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Bouncing Balls", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, ScreenWidth, ScreenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	for i := 0; i < NumBalls; i++ {
		balls = append(balls, Ball{
			X:      float64(rand.Intn(ScreenWidth-100) + 50),
			Y:      float64(rand.Intn(ScreenHeight-100) + 50),
			DX:     (rand.Float64() * 4) - 2,
			DY:     (rand.Float64() * 4) - 2,
			Radius: int32(rand.Intn(30) + 10),
			Color: sdl.Color{
				R: uint8(rand.Intn(256)),
				G: uint8(rand.Intn(256)),
				B: uint8(rand.Intn(256)),
				A: 255,
			},
			Mass: int32(rand.Intn(10)),
		})
	}

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseButtonEvent:
				fmt.Println(e.X, e.Y, e.Button, "MOUSE CLICK")

			}
		}

		CheckCollision(balls)

		// check edge hit on all balls and reflect in back
		for i := range balls {
			balls[i].X += balls[i].DX
			balls[i].Y += balls[i].DY

			// hit left or right ( 0 or screen width)
			if balls[i].X-float64(balls[i].Radius) <= 0 || balls[i].X+float64(balls[i].Radius) >= ScreenWidth {
				balls[i].DX = -balls[i].DX
			}
			// hit up or down (0 or screenheight)
			if balls[i].Y-float64(balls[i].Radius) <= 0 || balls[i].Y+float64(balls[i].Radius) >= ScreenHeight {
				balls[i].DY = -balls[i].DY
			}
		}

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		for _, ball := range balls {
			renderer.SetDrawColor(ball.Color.R, ball.Color.G, ball.Color.B, ball.Color.A)
			DrawCircle(renderer, int32(ball.X), int32(ball.Y), ball.Radius)
		}

		renderer.Present()
		sdl.Delay(16)
	}
}
