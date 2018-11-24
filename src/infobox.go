// An info box
// (c) Jani Nykänen

package main

// Constants
const (
	infoBoxTime    = 240
	infoBoxFlicker = 60
)

// Info box type
type infoBox struct {
	active  bool
	timer   float32
	msg     string
	cb      cbfun
	bmpFont *bitmap
}

// Update
func (ib *infoBox) update(input *inputManager, tm float32) {

	if !ib.active {
		return
	}

	// Update timer
	ib.timer -= 1.0 * tm
	if ib.timer <= 0.0 ||
		(ib.timer < (infoBoxTime-infoBoxFlicker) &&
			input.getButton("start") == statePressed) {

		ib.active = false
		// Call callback if defined
		if ib.cb != nil {
			ib.cb()
		}
	}
}

// Draw
func (ib *infoBox) draw(g *graphics) {

	if !ib.active {
		return
	}

	xoff := int32(-7)

	// Check if flickering
	limit := float32(infoBoxTime - infoBoxFlicker)
	if ib.timer > limit {

		t := ib.timer - limit
		if int(floorFloat32(t/4))%2 == 0 {
			return
		}
	}

	// Draw box
	w := int32(len(ib.msg)+1)*(16+xoff) + 8
	h := int32(16)
	drawBox(g, w, h)

	// Draw message
	g.drawText(ib.bmpFont, ib.msg, 128, 120-h/2, xoff, 0, true)

}

// Activate
func (ib *infoBox) activate(str string, cb cbfun) {

	ib.msg = str
	ib.cb = cb
	ib.timer = infoBoxTime
	ib.active = true
}

// Create an info box
func createInfoBox(ass *assetPack) *infoBox {

	ib := new(infoBox)

	// Set defaults
	ib.active = false
	ib.timer = 0.0
	ib.cb = nil
	ib.msg = ""

	// Get assets
	ib.bmpFont = ass.getBitmap("font")

	return ib
}
