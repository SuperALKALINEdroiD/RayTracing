package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// midpoint circle algorithm
func DrawCircle(renderer *sdl.Renderer, x0, y0, radius int32) {
	x := radius
	y := int32(0)
	err := int32(1 - x)

	for x >= y {
		renderer.DrawPoint(x0+x, y0+y)
		renderer.DrawPoint(x0+y, y0+x)
		renderer.DrawPoint(x0-y, y0+x)
		renderer.DrawPoint(x0-x, y0+y)
		renderer.DrawPoint(x0-x, y0-y)
		renderer.DrawPoint(x0-y, y0-x)
		renderer.DrawPoint(x0+y, y0-x)
		renderer.DrawPoint(x0+x, y0-y)

		y++
		if err <= 0 {
			// inside circle permiter
			err += 2*y + 1
		} else {
			// outside circle perimeter
			x--
			err += 2*(y-x) + 1
		}
	}
}

func CheckCollision(balls []Ball) {
	n := len(balls)
	// https://www.youtube.com/watch?v=rtBCVe3j_24

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			b1, b2 := &balls[i], &balls[j]

			dx := b2.X - b1.X
			dy := b2.Y - b1.Y
			distSquared := dx*dx + dy*dy
			radiusSum := b1.Radius + b2.Radius

			if distSquared < float64(radiusSum*radiusSum) {
				resolveCollision(b1, b2)
			}
		}
	}
}

func resolveCollision(b1, b2 *Ball) {
	// swap speed on collison
	b1.DX, b2.DX = b2.DX, b1.DX
	b1.DY, b2.DY = b2.DY, b1.DY
}
