// Game scene
// (c) Jani Nykänen

package main

// Game type
type game struct {
	ass         *assetPack
	gameStage   *stage
	trans       *transition
	evMan       *eventManager
	pauseScreen *pause
	info        *infoBox
}

// Reset
func (t *game) reset(sIndex int) {

	// Create game stage
	t.gameStage = createStage(sIndex, t.ass, t)
}

// "Ready reset"
func (t *game) readyReset() {

	// Set transition callback
	fn := func() {
		t.reset(t.gameStage.index)
	}
	// Activate transition
	t.trans.activate(fadeIn, 2.0, fn)
}

// Show info box
func (t *game) showInfoBox(victory bool) {

	fn1 := func() {
		t.readyReset()
	}
	fn2 := func() {
		t.quit()
	}

	if victory {
		t.info.activate("STAGE CLEAR!", fn2)

	} else {
		t.info.activate("OUT OF MOVES!", fn1)
	}
}

// Quit
func (t *game) quit() {

	// Set transition callback
	fn := func() {
		t.evMan.terminate()
	}
	// Activate transition
	t.trans.activate(fadeIn, 2.0, fn)
}

// Initialize
func (t *game) init(g *graphics, trans *transition, evMan *eventManager, ass *assetPack) error {

	// Store references for future use
	t.ass = ass
	t.trans = trans
	t.evMan = evMan

	// Create pause screen
	t.pauseScreen = createPause(t, ass)
	// Create info box
	t.info = createInfoBox(ass)

	// Fade out
	t.trans.activate(fadeOut, 2.0, nil)

	// Start with stage 1
	t.reset(1)

	return nil
}

// Update
func (t *game) update(input *inputManager, tm float32) {

	if !t.info.active {

		// Check if pause enabled
		if t.pauseScreen.active {
			t.pauseScreen.update(input)
			return
		}
		// Otherwise check if pause is to be enabled
		if input.getButton("start") == statePressed ||
			input.getButton("cancel") == statePressed {

			t.pauseScreen.activate()
			return
		}

	} else {

		// Update info screen
		t.info.update(input, tm)
	}

	// If info box active, cut here
	if t.info.active {
		return
	}

	// Reset
	if input.getButton("restart") == statePressed {
		// t.reset(t.gameStage.index)
		t.readyReset()
		return
	}

	// Update stage
	t.gameStage.update(input, tm)
}

// Draw
func (t *game) draw(g *graphics) {

	// Draw stage
	t.gameStage.draw(g)

	// Draw pause screen
	t.pauseScreen.draw(g)

	// Draw info box
	t.info.draw(g)
}

// Destroy
func (t *game) destroy() {

}

// Scene changed
func (t *game) onChange() {

}
