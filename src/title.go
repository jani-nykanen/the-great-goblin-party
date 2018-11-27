// Title screen scene
// (c) Jani Nykänen

package main

// Constants
const (
	startTimeMax    = 30
	menuMusicVolume = 0.5
)

// Title screen type
type titleScreen struct {
	bmpFont    *bitmap
	bmpLogo    *bitmap
	bmpBorders *bitmap
	sMenu      *sample
	trans      *transition
	evMan      *eventManager
	audio      *audioManager
	phase      int
	startTimer float32
	tmenu      *menu
}

// Initialize
func (ts *titleScreen) init(g *graphics, trans *transition, evMan *eventManager,
	audio *audioManager, ass *assetPack) error {

	// Store references
	ts.trans = trans
	ts.evMan = evMan
	ts.audio = audio

	// Get bitmaps
	ts.bmpFont = ass.getBitmap("font")
	ts.bmpLogo = ass.getBitmap("logo")
	ts.bmpBorders = ass.getBitmap("borders")
	// Get samples
	ts.sMenu = ass.getSample("menu")

	// Set defaults
	ts.phase = 0
	ts.startTimer = 0.0

	// Create menu
	fn1 := func() {
		fun := func() {
			ts.evMan.changeScene(0, "stagemenu")
		}
		ts.trans.activate(fadeIn, 2.0, fun)
	}
	fn2 := func() {
		ts.trans.activate(fadeIn, 2.0, ts.evMan.terminate)
	}
	ts.tmenu = createMenu([]string{"Start game", "Quit"},
		[]cbfun{fn1, fn2})

	return nil
}

// Update
func (ts *titleScreen) update(input *inputManager, tm float32) {

	// Check escape
	if input.getButton("cancel") == statePressed {

		// Quit
		ts.trans.activate(fadeIn, 2.0, ts.evMan.terminate)
	}

	// Update start timer, if phase is zero aka. "press enter"
	if ts.phase == 0 {

		ts.startTimer += 1.0 * tm
		if ts.startTimer >= startTimeMax {
			ts.startTimer -= startTimeMax
		}

		// Check enter
		if input.getButton("start") == statePressed {
			ts.phase++
		}

	} else {

		// Update menu
		ts.tmenu.update(input)
	}
}

// Draw
func (ts *titleScreen) draw(g *graphics) {

	xoff := int32(-7)
	logoY := int32(32)
	copyrightY := int32(32)

	pauseY := int32(160)
	menuX := int32(128 - 64)
	menuY := int32(144)

	// Clear screen
	g.clearScreen(30, 160, 248)

	// Draw borders
	drawBorders(g, ts.bmpBorders, 8, 8, 16-1, 15-1)

	// Draw logo
	g.drawBitmap(ts.bmpLogo, 128-int32(ts.bmpLogo.width/2), logoY, flipNone)

	if ts.phase == 0 {

		// Draw "Press enter" text
		if ts.startTimer >= startTimeMax/2 {
			g.drawText(ts.bmpFont, "PRESS ENTER", 128, pauseY, xoff, 0, true)
		}

	} else {

		// Draw menu
		ts.tmenu.drawMenu(g, ts.bmpFont,
			menuX,
			menuY, 20)
	}

	// Draw copyright
	g.drawText(ts.bmpFont, "(c)2018 Jani Nyk%nen", 128, 240-copyrightY, xoff, 0, true)
}

// Destroy
func (ts *titleScreen) destroy() {

}

// Scene changed
func (ts *titleScreen) onChange(param int) {

	// Play music
	ts.audio.playMusic(ts.sMenu, menuMusicVolume)
}

// Get name
func (ts *titleScreen) getName() string {
	return "titlescreen"
}
