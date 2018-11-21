// Stage
// (c) Jani Nykänen

package main

import (
	"strconv"
)

const (
	stageYOff = 0
)

// Stage type
type stage struct {
	baseMap      *tilemap
	width        int32
	height       int32
	index        int
	bmpTestTiles *bitmap
	xpos         int32
	ypos         int32
	objects      []object
}

// Update stage
func (s *stage) update(input *inputManager, tm float32) {

	// Update objects
	for i := 0; i < len(s.objects); i++ {

		s.objects[i].update(input, tm)
	}
}

// Draw stage
func (s *stage) draw(g *graphics) {

	// "Shadow"
	g.setGlobalColor(0, 0, 85, 255)
	g.fillRect(s.xpos-2+12, s.ypos-2+12, s.width*16+4, s.height*16+4)

	// Borders
	g.setGlobalColor(0, 0, 0, 255)
	g.fillRect(s.xpos-2, s.ypos-2, s.width*16+4, s.height*16+4)
	g.setGlobalColor(255, 255, 255, 255)
	g.fillRect(s.xpos-1, s.ypos-1, s.width*16+2, s.height*16+2)

	// Background
	g.setGlobalColor(0, 0, 0, 255)
	g.fillRect(s.xpos, s.ypos, s.width*16, s.height*16)

	// Draw tiles (temp)
	var tileID, sx, sy int32
	for y := int32(0); y < s.height; y++ {

		for x := int32(0); x < s.width; x++ {

			// Get tileID
			tileID = int32(s.baseMap.getTile(x, y))
			if tileID <= 0 {
				continue
			}

			tileID--

			// Draw tile
			sx = tileID % 16
			sy = tileID / 16
			g.drawBitmapRegion(s.bmpTestTiles, sx*16, sy*16, 16, 16,
				int32(s.xpos+x*16), int32(s.ypos+y*16), flipNone)
		}
	}

	// Draw objects
	for i := 0; i < len(s.objects); i++ {

		s.objects[i].draw(g)
	}
}

// Add an object
func (s *stage) addObject(o object) {
	s.objects = append(s.objects, o)
}

// Create a new stage
func createStage(index int, ass *assetPack) *stage {

	s := new(stage)

	// Load base map
	s.baseMap = ass.getTilemap(strconv.Itoa(index))
	// Get assets
	s.bmpTestTiles = ass.getBitmap("testTiles")
	// Get data
	s.width = int32(s.baseMap.width)
	s.height = int32(s.baseMap.height)

	// Calculate position
	s.xpos = 128 - s.width*16/2
	s.ypos = stageYOff + (240-stageYOff)/2 - s.height*16/2

	// Create an empty object list
	s.objects = make([]object, 0)

	s.index = index

	return s
}
