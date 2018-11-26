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

// Draw a nice box with border Always
// in the center of the screen!
func drawBox(g *graphics, w, h int32) {

	g.setGlobalColor(0, 0, 0, 0)
	g.fillRect(128-w/2-2, 120-h/2-2, w+4, h+4)
	g.setGlobalColor(255, 255, 255, 255)
	g.fillRect(128-w/2-1, 120-h/2-1, w+2, h+2)
	g.setGlobalColor(72, 72, 72, 72)
	g.fillRect(128-w/2, 120-h/2, w, h)
}

// Get difficulty string
func getDifficultyString(dif int) string {

	ret := ""

	// Full stars
	for i := 0; i < int(dif/2); i++ {
		ret += "#"
	}

	// Half stars
	if dif%2 == 1 {
		ret += "$"
	}

	return ret
}

// Draw borders
func drawBorders(g *graphics, bmpBorders *bitmap, xpos, ypos, width, height int32) {

	ypos -= 8
	xpos -= 8
	xjump := width*16 + 8
	yjump := height*16 + 8

	// Horizontal
	for x := 0; x < int(width)*2; x++ {

		// Top
		g.drawBitmapRegion(bmpBorders, 8, 0, 8, 8,
			xpos+8+int32(x)*8, ypos, flipNone)
		// Bottom
		g.drawBitmapRegion(bmpBorders, 8, 16, 8, 8,
			xpos+8+int32(x)*8, ypos+yjump, flipNone)
	}

	// Vertical
	for y := 0; y < int(height)*2; y++ {

		// Left
		g.drawBitmapRegion(bmpBorders, 0, 8, 8, 8,
			xpos, ypos+int32(y)*8+8, flipNone)
		// Right
		g.drawBitmapRegion(bmpBorders, 16, 8, 8, 8,
			xpos+xjump, ypos+int32(y)*8+8, flipNone)
	}

	// Corners
	g.drawBitmapRegion(bmpBorders, 0, 0, 8, 8,
		xpos, ypos, flipNone)
	g.drawBitmapRegion(bmpBorders, 16, 0, 8, 8,
		xpos+xjump, ypos, flipNone)
	g.drawBitmapRegion(bmpBorders, 0, 16, 8, 8,
		xpos, ypos+yjump, flipNone)
	g.drawBitmapRegion(bmpBorders, 16, 16, 8, 8,
		xpos+xjump, ypos+yjump, flipNone)
}
