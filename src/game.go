// Game scene
// (c) Jani Nykänen

package main

// Game type
type game struct {
	ass       *assetPack
	bmpFont   *bitmap
	gameStage *stage
}

// Reset
func (t *game) reset(sIndex int) {

	t.gameStage = createStage(sIndex, t.ass)
}

// Initialize
func (t *game) init(g *graphics, ass *assetPack) error {

	// Store assets for future use
	t.ass = ass
	// Get font
	t.bmpFont = ass.getBitmap("font")

	// Start with stage 1
	t.reset(1)

	return nil
}

// Update
func (t *game) update(input *inputManager, tm float32) {

	// Update stage
	t.gameStage.update(input, tm)
}

// Draw
func (t *game) draw(g *graphics) {

	g.clearScreen(0, 85, 170)

	// Hello world!
	g.drawText(t.bmpFont, "Hello world!", 128, 8, -7, 0, true)

	// Draw stage
	t.gameStage.draw(g)
}

// Destroy
func (t *game) destroy() {

}

// Scene changed
func (t *game) onChange() {

}
