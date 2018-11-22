// Stage
// (c) Jani Nykänen

package main

import (
	"strconv"
)

const (
	stageYOff = 16
)

// Stage type
type stage struct {
	baseMap      *tilemap
	width        int32
	height       int32
	index        int
	bmpTestTiles *bitmap
	bmpFont      *bitmap
	xpos         int32
	ypos         int32
	objects      []object
}

// Get difficulty string
func (s *stage) getDifficultyString() string {

	dif := s.baseMap.difficulty
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

// Update stage
func (s *stage) update(input *inputManager, tm float32) {

	// Update objects
	for i := 0; i < len(s.objects); i++ {

		s.objects[i].update(input, tm)
	}
}

// Draw map
func (s *stage) drawMap(g *graphics) {

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

}

// Draw info
func (s *stage) drawInfo(g *graphics) {

	stageIndexY := int32(8)
	nameY := int32(24)
	xoff := int32(-6)
	starXoff := int32(-3)
	bottomY := int32(16)
	bottomXOff := int32(2)
	difMinusX := int32(-4)

	// Draw stage index
	g.drawText(s.bmpFont, "Stage "+strconv.Itoa(s.index),
		128, stageIndexY, xoff, 0, true)
	// Draw stage name
	g.drawText(s.bmpFont, "\""+s.baseMap.name+"\"",
		128, nameY, xoff, 0, true)

	// Draw difficulty text
	str := "Difficulty: "
	g.drawText(s.bmpFont, "Difficulty: ",
		bottomXOff, 240-bottomY, xoff, 0, false)
	// Draw difficulty
	g.drawText(s.bmpFont, s.getDifficultyString(),
		bottomXOff+int32(len(str)*10)+difMinusX, 240-bottomY, starXoff, 0, false)

	// Draw moves
	str = "Moves: " + strconv.Itoa(s.baseMap.moveLimit)
	g.drawText(s.bmpFont, str,
		256-int32(len(str)+1)*10+bottomXOff,
		240-bottomY, xoff, 0, false)
}

// Draw stage
func (s *stage) draw(g *graphics) {

	// Draw map
	s.drawMap(g)

	// Draw objects
	for i := 0; i < len(s.objects); i++ {

		s.objects[i].draw(g)
	}

	// Draw info
	s.drawInfo(g)
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
	s.bmpFont = ass.getBitmap("font")
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
