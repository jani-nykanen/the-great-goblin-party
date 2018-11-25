// Stage menu scene
// (c) Jani Nykänen

package main

import (
	"strconv"
)

// Menu button type
type stageButton struct {
	x, y     int32
	number   int
	selected bool
	spr      sprite
}

// Stage menu type
type stageMenu struct {
	bmpFont    *bitmap
	bmpButton  *bitmap
	bmpNumbers *bitmap
	maps       [12]*tilemap
	trans      *transition
	evMan      *eventManager
	buttons    [12]stageButton
	cx, cy     int
	cursor     int
}

// Initialize
func (sm *stageMenu) init(g *graphics, trans *transition, evMan *eventManager, ass *assetPack) error {

	buttonYStart := int32(32)
	buttonYOff := int32(48)
	buttonXStart := int32(32)
	buttonXOff := int32(48)

	// Store references
	sm.trans = trans
	sm.evMan = evMan

	// Get bitmaps
	sm.bmpFont = ass.getBitmap("font")
	sm.bmpNumbers = ass.getBitmap("numbers")
	sm.bmpButton = ass.getBitmap("button")

	// Get tilemaps
	for i := 0; i < 12; i++ {

		sm.maps[i] = ass.getTilemap(strconv.Itoa(i + 1))
	}

	// Set buttons
	var i int
	for y := 0; y < 3; y++ {
		for x := 0; x < 4; x++ {
			i = y*4 + x
			sm.buttons[i] = stageButton{
				buttonXStart + int32(x)*buttonXOff,
				buttonYStart + int32(y)*buttonYOff, i + 1, false,
				createSprite(48, 48),
			}
		}
	}

	// Set default cursor positions
	sm.cx = 0
	sm.cy = 0

	return nil
}

// Update cursor
func (sm *stageMenu) updateCursor(input *inputManager) {

	// Update cursor
	if input.getButton("left") == statePressed {
		sm.cx--

	} else if input.getButton("right") == statePressed {
		sm.cx++

	} else if input.getButton("up") == statePressed {
		sm.cy--

	} else if input.getButton("down") == statePressed {
		sm.cy++

	}
	// Restrict
	if sm.cy < 0 {
		sm.cy += 3
	}
	if sm.cx < 0 {
		sm.cx += 4
	}
	sm.cy %= 3
	sm.cx %= 4
}

// Update
func (sm *stageMenu) update(input *inputManager, tm float32) {

	// Update cursor
	sm.updateCursor(input)

	// Set sprite frames
	sm.cursor = sm.cy*4 + sm.cx
	var b *stageButton
	for i := 0; i < len(sm.buttons); i++ {

		b = &sm.buttons[i]
		// Animate
		if i == sm.cursor {
			b.spr.frame = 1

		} else {
			b.spr.frame = 0
		}
	}

	// Check button press
	if input.getButton("start") == statePressed &&
		sm.maps[sm.cursor] != nil {

		fn := func() {
			sm.evMan.changeScene(sm.cursor+1, "game")
		}
		sm.trans.activate(fadeIn, 2, fn)
	}

	// Check escape button
	if input.getButton("cancel") == statePressed {

		fn := func() {
			sm.evMan.terminate()
		}
		sm.trans.activate(fadeIn, 2.0, fn)
	}
}

// Draw info
func (sm *stageMenu) drawInfo(g *graphics) {

	startX := int32(8)
	xoff := int32(-7)
	difXoff := int32(-3)
	startY := int32(184)
	yoff := int32(16)

	m := sm.maps[sm.cursor]
	if m == nil {
		return
	}

	// Get info
	name := m.name
	diff := m.difficulty
	moves := m.moveLimit

	// Draw info
	g.drawText(sm.bmpFont, "NAME: "+name, startX, startY, xoff, 0, false)
	// Draw difficulty
	str := "DIFFICULTY: "
	g.drawText(sm.bmpFont, str, startX, startY+yoff, xoff, 0, false)
	g.drawText(sm.bmpFont, getDifficultyString(diff),
		startX+int32(len(str))*(16+xoff), startY+yoff, difXoff, 0, false)
	// Draw move limit
	g.drawText(sm.bmpFont, "MOVE LIMIT: "+strconv.Itoa(moves),
		startX, startY+yoff*2, xoff, 0, false)
}

// Draw
func (sm *stageMenu) draw(g *graphics) {

	xoff := int32(-7)
	titleY := int32(12)

	xplus := int32(-4)
	yplus := int32(-4)

	shadowOff := int32(3)

	// Clear screen
	g.clearScreen(30, 160, 248)

	// Draw title
	g.drawText(sm.bmpFont, "CHOOSE A STAGE", 128, titleY, xoff, 0, true)

	// Draw buttons
	var dx, dy int32
	for i := 0; i < len(sm.buttons); i++ {

		dx = sm.buttons[i].x
		dy = sm.buttons[i].y

		// Draw shadow
		sm.buttons[i].spr.drawFrame(g, sm.bmpButton, 3, 0, dx+shadowOff, dy+shadowOff, flipNone)

		// If chosen
		if sm.cursor == i {
			dx += xplus
			dy += yplus
		}
		// Draw button
		sm.buttons[i].spr.draw(g, sm.bmpButton, dx, dy, flipNone)
		// Draw stage index
		str := strconv.Itoa(i + 1)
		g.drawText(sm.bmpNumbers, str, dx+20, dy+8, -20, 0, true)
	}

	// Draw info
	sm.drawInfo(g)
}

// Destroy
func (sm *stageMenu) destroy() {

}

// Scene changed
func (sm *stageMenu) onChange(param int) {
	// ...
}

// Get name
func (sm *stageMenu) getName() string {
	return "stagemenu"
}
