// Intro scene
// (c) Jani Nykänen

package main

// Intro type
type intro struct {
	bmpIntro *bitmap
	trans    *transition
	evMan    *eventManager
	sprFace  sprite
	textSrc  int32
}

// Initialize
func (t *intro) init(g *graphics, trans *transition, evMan *eventManager, ass *assetPack) error {

	// Store references
	t.trans = trans
	t.evMan = evMan

	// Get bitmaps
	t.bmpIntro = ass.getBitmap("intro")

	// Create sprite
	t.sprFace = createSprite(64, 64)

	// Set defaults
	t.textSrc = 256

	return nil
}

// Update
func (t *intro) update(input *inputManager, tm float32) {

	animSpeed1 := 10
	animSpeed2 := 90
	animSpeed3 := 45

	// Determine animation speed
	speed := animSpeed1
	if t.sprFace.frame == 3 {
		speed = animSpeed2
	} else if t.sprFace.frame == 0 {
		speed = animSpeed3
	}

	// Animate
	t.sprFace.animate(0, 0, 6, float32(speed), tm)

	// Determine text source x
	t.textSrc = 256
	if t.sprFace.frame == 3 {
		t.textSrc = 128

	} else if t.sprFace.frame > 0 && t.sprFace.frame < 6 {
		t.textSrc = 0
	}

	// Check if the final frame
	if t.sprFace.frame == 6 {

		fn := func() {
			t.evMan.changeScene(0, "titlescreen")
		}

		t.trans.activate(fadeIn, 2.0, fn)
	}
}

// Draw
func (t *intro) draw(g *graphics) {

	faceY := int32(64)
	textY := int32(128)

	// Clear screen
	g.clearScreen(0, 0, 0)

	// Draw face
	t.sprFace.draw(g, t.bmpIntro, 128-32, faceY, flipNone)

	// Draw "game by" text
	g.drawBitmapRegion(t.bmpIntro, t.textSrc, 64, 128, 64, 128-64, textY, flipNone)

}

// Destroy
func (t *intro) destroy() {

}

// Scene changed
func (t *intro) onChange(param int) {

	// ...

}

// Get name
func (t *intro) getName() string {
	return "intro"
}
