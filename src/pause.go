// Pause screen
// (c) Jani Nykänen

package main

// Pause type
type pause struct {
	pauseMenu *menu
	active    bool
	gameRef   *game
	bmpFont   *bitmap
}

// Update
func (p *pause) update(input *inputManager) {

	if !p.active {
		return
	}

	// Update pause menu
	p.pauseMenu.update(input)

	// Check if escape pressed
	if input.getButton("cancel") == statePressed {
		p.active = false
	}
}

// Draw
func (p *pause) draw(g *graphics) {

	if !p.active {
		return
	}

	w := int32(96)
	h := int32(56)

	menuX := int32(8)
	menuY := int32(4)
	yoff := int32(16)

	g.translate(0, 0)

	// Draw box
	drawBox(g, w, h)

	// Draw menu
	p.pauseMenu.drawMenu(g, p.bmpFont, 128-w/2+menuX, 120-h/2+menuY, yoff)
}

// Activate
func (p *pause) activate() {

	p.active = true
	p.pauseMenu.cursorPos = 0
}

// Create a pause screen object
func createPause(gameRef *game, ass *assetPack) *pause {

	p := new(pause)

	// Store game reference
	p.gameRef = gameRef

	// Create menu
	str := []string{
		"RESUME",
		"RESTART",
		"QUIT",
	}
	cbs := make([]cbfun, 3)
	cbs[0] = func() {
		p.active = false
	}
	cbs[1] = func() {
		p.active = false
		p.gameRef.readyReset()
	}
	cbs[2] = func() {
		panic("Not yet implemented!")
	}
	p.pauseMenu = createMenu(str, cbs)

	// Get assets
	p.bmpFont = ass.getBitmap("font")

	p.active = false
	return p
}
