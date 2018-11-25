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

	// Get assets
	sm.bmpFont = ass.getBitmap("font")
	sm.bmpNumbers = ass.getBitmap("numbers")
	sm.bmpButton = ass.getBitmap("button")

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

// Update
func (sm *stageMenu) update(input *inputManager, tm float32) {

	flickerSpeed := float32(15.0)

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

	// Set flickering
	old := sm.cursor
	sm.cursor = sm.cy*4 + sm.cx
	var b *stageButton
	for i := 0; i < len(sm.buttons); i++ {

		b = &sm.buttons[i]
		// Animate
		if i == sm.cursor {

			b.spr.animate(0, 0, 1, flickerSpeed, tm)
			if old != sm.cursor {
				b.spr.frame = 1
				b.spr.count = 0.0
			}

		} else {

			b.spr.frame = 0
		}
	}

	// Check button press
	if input.getButton("start") == statePressed {

		fn := func() {
			sm.evMan.changeScene(sm.cursor+1, "game")
		}
		sm.trans.activate(fadeIn, 2, fn)
	}
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
