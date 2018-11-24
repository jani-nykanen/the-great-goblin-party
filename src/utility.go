// Utilities
// (c) Jani Nykänen

package main

import (
	"math"
)

// Callback function type
type cbfun func()

// Int min
func minInt(a, b int) int {

	if a < b {
		return a
	}
	return b
}

// Float32 floor
func floorFloat32(a float32) float32 {

	return float32(math.Floor(float64(a)))
}

// Int max
func maxInt(a, b int) int {

	if a > b {
		return a
	}
	return b
}

// Int abs
func absInt(a int) int {

	if a < 0 {
		return -a
	}
	return a
}

// Float32 abs
func absFloat32(a float32) float32 {

	if a < 0.0 {
		return -a
	}
	return a
}

// Draw a nice box with borders. Always
// in the center of the screen!
func drawBox(g *graphics, w, h int32) {

	g.setGlobalColor(0, 0, 0, 0)
	g.fillRect(128-w/2-2, 120-h/2-2, w+4, h+4)
	g.setGlobalColor(255, 255, 255, 255)
	g.fillRect(128-w/2-1, 120-h/2-1, w+2, h+2)
	g.setGlobalColor(72, 72, 72, 72)
	g.fillRect(128-w/2, 120-h/2, w, h)
}
