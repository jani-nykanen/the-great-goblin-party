// Ending scene
// (c) Jani Nykänen

package main

import (
	"math"
)

// Ending type
type ending struct {
	bmpFont    *bitmap
	bmpGremlin *bitmap
	trans      *transition
	evMan      *eventManager
	sprStar    sprite
	jumpTimer  float32
}

// Initialize
func (e *ending) init(g *graphics, trans *transition, evMan *eventManager, ass *assetPack) error {

	// Store references
	e.trans = trans
	e.evMan = evMan

	// Get bitmaps
	e.bmpFont = ass.getBitmap("font")
	e.bmpGremlin = ass.getBitmap("gremlin")

	// Create star sprite
	e.sprStar = createSprite(16, 16)
	e.jumpTimer = 0.0

	return nil
}

// Update
func (e *ending) update(input *inputManager, tm float32) {

	animSpeed := float32(6.0)
	jumpSpeed := float32(0.05)

	// Animate star
	e.sprStar.animate(0, 0, 3, animSpeed, tm)

	// Update jump time
	e.jumpTimer += jumpSpeed * tm

	// Check enter or cancel
	if input.getButton("start") == statePressed ||
		input.getButton("cancel") == statePressed {

		e.trans.activate(fadeIn, 2.0, func() {

			e.evMan.changeScene(0, "stagemenu")
		})
	}
}

// Draw
func (e *ending) draw(g *graphics) {

	congrY := int32(96)
	xoff := int32(-7)
	textX := int32(16)
	textY := int32(128)
	starX := int32(88)
	starXOff := int32(32)
	starY := int32(88)
	jumpHeight := 16.0

	// Clear screen
	g.clearScreen(0, 0, 0)

	// Draw stars
	var jump float32
	for i := 0; i < 3; i++ {

		// Calculate jump height
		// (these type casts... Go is a nice language
		//  indeed)
		jump = float32(math.Abs(jumpHeight * math.Sin(math.Pi/4.0*float64(i)+float64(e.jumpTimer))))

		e.sprStar.drawFrame(g, e.bmpGremlin, e.sprStar.frame,
			2+int32(i)*5, starX+starXOff*int32(i), starY-16-int32(jump), flipNone)
	}

	// Draw "congratulations"
	g.drawText(e.bmpFont, "CONGRATULATIONS!", 128, congrY, xoff, 0, true)
	// Draw text
	g.drawText(e.bmpFont,
		"Thanks to you the goblins\n"+
			"can start a great party\n"+
			"where everyone is a star!\n"+
			"\n"+
			"Literally.",
		textX, textY, xoff, 0, false)
}

// Destroy
func (e *ending) destroy() {

}

// Scene changed
func (e *ending) onChange(param int) {

	// ...

}

// Get name
func (e *ending) getName() string {
	return "ending"
}
