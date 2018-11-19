// Application
// (c) Jani Nykänen

package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Application class
type application struct {
	running      bool
	g            *graphics
	scenes       []scene
	currentScene *scene
	conf         config
	window       *sdl.Window
	rend         *sdl.Renderer
	input        *inputManager
	isFullscreen bool
}

// Initialize SDL2 content
func (app *application) initSDL() error {

	var err error

	// Initialize SDL
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {

		return err
	}

	// Open window using configuration parameters
	app.window, err = sdl.CreateWindow(app.conf.caption,
		sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		app.conf.winWidth, app.conf.winHeight, sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE)
	if err != nil {
		return err
	}

	// Toggle fullscreen if necessary
	app.isFullscreen = false
	if app.conf.fullscreen {
		app.toggleFullscreen()
	}

	// Create renderer
	app.rend, err = sdl.CreateRenderer(app.window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		return err
	}

	return nil
}

// Initialize
func (app *application) init(conf config) error {

	var err error

	// Store configuration
	app.conf = conf

	// Initialize SDL2 content
	err = app.initSDL()
	if err != nil {

		return err
	}

	// Hide mouse cursor
	sdl.ShowCursor(0)

	// Create input
	app.input = new(inputManager)

	// Create graphics
	app.g = new(graphics)
	err = app.g.init(app.rend)
	if err != nil {
		return err
	}

	// Create a slice for scenes
	app.scenes = make([]scene, 0)
	// Set current scene to nil
	app.currentScene = nil

	// Set running
	app.running = true

	return err
}

// Toggle fullscreen
func (app *application) toggleFullscreen() {

	if !app.isFullscreen {

		app.window.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
	} else {

		app.window.SetFullscreen(0)
	}
	app.isFullscreen = !app.isFullscreen
}

// Poll events
func (app *application) pollEvents() {

	// Go through events
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {

		switch t := event.(type) {

		// Quit
		case *sdl.QuitEvent:
			app.running = false
			break

		// Keyboard event
		case *sdl.KeyboardEvent:

			if t.Type == sdl.KEYDOWN {
				app.input.keyDown(int32(t.Keysym.Scancode))

			} else if t.Type == sdl.KEYUP {

				app.input.keyUp(int32(t.Keysym.Scancode))
			}
			break

		default:
			break
		}
	}
}

// Loop
func (app *application) loop() {

	maxUpdate := 5
	wait := uint32(1000 / app.conf.frameRate)
	waitCount := 0

	redraw := false
	timeSum := uint32(0)
	oldTime := sdl.GetTicks()

	for app.running {

		// Add time
		timeSum += sdl.GetTicks() - oldTime

		waitCount = 0
		for timeSum >= wait {
			// Update scene
			app.update(float32(wait))
			// Get base key events
			app.baseKeyEvents()
			// Update keys
			app.input.updateKeys()

			// Reduce the wait time from time sum
			timeSum -= wait

			// Check that we don't update too many times
			waitCount++
			if waitCount >= maxUpdate {
				break
			}

			redraw = true
		}

		// Redraw if necessary
		if redraw {
			app.draw()
			redraw = false
		}

		// Update frame
		app.rend.Present()

		// Poll events
		app.pollEvents()
	}
}

// Base key events
func (app *application) baseKeyEvents() {

	// Quit
	if app.input.getKey(sdl.SCANCODE_LCTRL) == stateDown &&
		app.input.getKey(sdl.SCANCODE_Q) == statePressed {

		app.running = false
	}

	// Fullscreen
	if (app.input.getKey(sdl.SCANCODE_LALT) == stateDown &&
		app.input.getKey(sdl.SCANCODE_RETURN) == statePressed) ||
		app.input.getKey(sdl.SCANCODE_F4) == statePressed {

		app.toggleFullscreen()
	}
}

// Update
func (app *application) update(delta float32) {

	tm := delta / 1000.0 / (1.0 / 60.0)

	// Update the current scene
	if app.currentScene != nil {

		(*app.currentScene).update(app.input, tm)
	}
}

// Draw
func (app *application) draw() {

	// Draw the current scene
	if app.currentScene != nil {

		(*app.currentScene).draw(app.g)
	}
}

// Destroy
func (app *application) destroy() {

	app.window.Destroy()
}

// Create
func (app *application) run() error {

	// Initialize scenes
	var err error
	var s scene
	for i := 0; i < len(app.scenes); i++ {

		s = app.scenes[i]
		err = s.init()
		if err != nil {

			return err
		}
	}

	// Loop
	app.loop()

	// Destroy scenes
	for i := 0; i < len(app.scenes); i++ {

		s = app.scenes[i]
		s.destroy()
	}
	// Destroy application
	app.destroy()

	return err
}

// Add a scene
func (app *application) addScene(s scene, makeCurrent bool) {

	app.scenes = append(app.scenes, s)
	if makeCurrent {
		app.currentScene = &s
	}
}
