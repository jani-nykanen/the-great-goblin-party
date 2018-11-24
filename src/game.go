// Game scene
// (c) Jani Nykänen

package main

// Game type
type game struct {
	ass       *assetPack
	gameStage *stage
	trans     *transition
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

// Initialize
func (t *game) init(g *graphics, trans *transition, ass *assetPack) error {

	// Store references for future use
	t.ass = ass
	t.trans = trans

	// Fade out
	t.trans.activate(fadeOut, 2.0, nil)

	// Start with stage 1
	t.reset(1)

	return nil
}

// Update
func (t *game) update(input *inputManager, tm float32) {

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
}

// Destroy
func (t *game) destroy() {

}

// Scene changed
func (t *game) onChange() {

}
