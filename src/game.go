// Game scene
// (c) Jani Nykänen

package main

// Contants
const (
	themeVol = 0.70
)

// Game type
type game struct {
	ass         *assetPack
	gameStage   *stage
	trans       *transition
	evMan       *eventManager
	audio       *audioManager
	pauseScreen *pause
	info        *infoBox
	sTheme      *sample
	sPause      *sample
	sVictory    *sample
	sDefeat     *sample
	sReset      *sample
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
		// Replay music
		t.audio.playMusic(t.sTheme, themeVol)
	}
	// Activate transition
	t.trans.activate(fadeIn, 2.0, fn)
}

// Show info box
func (t *game) showInfoBox(victory bool) {

	if t.info.active || t.trans.active {
		return
	}

	fn1 := func() {
		t.readyReset()
	}
	fn2 := func() {
		t.quit(1)
	}

	t.audio.stopMusic()
	if victory {
		t.info.activate("STAGE CLEAR!", fn2)
		t.audio.playSample(t.sVictory, 0.60)

	} else {
		t.info.activate("OUT OF MOVES!", fn1)
		t.audio.playSample(t.sDefeat, 0.80)
	}
}

// Quit
func (t *game) quit(state int) {

	// Set transition callback
	fn := func() {
		t.evMan.changeScene(state, "stagemenu")
	}
	// Activate transition
	t.trans.activate(fadeIn, 2.0, fn)
}

// Initialize
func (t *game) init(g *graphics, trans *transition, evMan *eventManager,
	audio *audioManager, ass *assetPack) error {

	// Store references for future use
	t.ass = ass
	t.trans = trans
	t.evMan = evMan
	t.audio = audio

	// Create pause screen
	t.pauseScreen = createPause(t, ass)
	// Create info box
	t.info = createInfoBox(ass)

	// Get samples
	t.sTheme = ass.getSample("theme")
	t.sPause = ass.getSample("pause")
	t.sVictory = ass.getSample("victory")
	t.sDefeat = ass.getSample("defeat")
	t.sReset = ass.getSample("reset")

	return nil
}

// Update
func (t *game) update(input *inputManager, tm float32) {

	if !t.info.active {

		// Check if pause enabled
		if t.pauseScreen.active {
			t.pauseScreen.update(input, t.audio)
			return
		}
		// Otherwise check if pause is to be enabled
		if input.getButton("start") == statePressed ||
			input.getButton("cancel") == statePressed {

			t.audio.playSample(t.sPause, 0.30)

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
		t.audio.playSample(t.sReset, 0.50)
		return
	}

	// Update stage
	t.gameStage.update(input, t.audio, tm)
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
func (t *game) onChange(param int) {

	// Stop music
	t.audio.stopMusic()
	// Play theme
	t.audio.playMusic(t.sTheme, themeVol)

	// Deactive pause menu & info box
	t.info.active = false
	t.pauseScreen.active = false

	// Reset
	t.reset(param)
}

// Get name
func (t *game) getName() string {
	return "game"
}
