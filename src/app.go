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
	winID        uint32
	rend         *sdl.Renderer
	input        *inputManager
	isFullscreen bool
	canvas       *bitmap
	canvasPos    sdl.Rect
	assets       *assetPack
	trans        *transition
	evMan        *eventManager
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
	app.winID, _ = app.window.GetID()
	// Calculate canvas info
	app.resizeEvent(app.conf.winWidth, app.conf.winHeight)

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

	// Create canvas
	app.canvas, err = createEmptyBitmap(app.g,
		app.conf.canvasWidth,
		app.conf.canvasHeight,
		true)
	if err != nil {
		return err
	}

	// Create transition
	app.trans = createTransition()
	// Create event manager
	app.evMan = createEventManager(app)

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

// Resize event
func (app *application) resizeEvent(w, h int32) {

	// Check if horizontal aspect ratio
	// and update canvas position & size info
	if float32(w)/float32(h) >= float32(app.conf.canvasWidth)/float32(app.conf.canvasHeight) {

		app.canvasPos.H = h
		app.canvasPos.W = int32(float32(h) / float32(app.conf.canvasHeight) * float32(app.conf.canvasWidth))

		app.canvasPos.Y = 0
		app.canvasPos.X = w/2 - app.canvasPos.W/2

	} else {

		app.canvasPos.W = w
		app.canvasPos.H = int32(float32(w) / float32(app.conf.canvasWidth) * float32(app.conf.canvasHeight))

		app.canvasPos.X = 0
		app.canvasPos.Y = h/2 - app.canvasPos.H/2
	}
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

		// Window event
		case *sdl.WindowEvent:

			if t.WindowID == app.winID &&
				t.Event == sdl.WINDOWEVENT_RESIZED {

				app.resizeEvent(t.Data1, t.Data2)
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
	newTime := oldTime

	for app.running {

		// Add time
		oldTime = newTime
		newTime = sdl.GetTicks()
		timeSum += newTime - oldTime

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
			app.drawToCanvas()
			redraw = false
		}

		// Draw canvas (note: in every frame)
		app.draw()
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

	// Update transition
	if app.trans.active {

		app.trans.update(tm)
		return
	}

	// Update the current scene
	if app.currentScene != nil {

		(*app.currentScene).update(app.input, tm)
	}

}

// Draw to canvas
func (app *application) drawToCanvas() {

	app.g.setRenderTarget(app.canvas)

	// Draw the current scene
	if app.currentScene != nil {

		(*app.currentScene).draw(app.g)
	}

	// Draw transition
	app.trans.draw(app.g, int32(app.canvas.width), int32(app.canvas.height))

	app.g.setRenderTarget(nil)
}

// Draw
func (app *application) draw() {

	app.g.clearScreen(0, 0, 0)

	// Draw canvas content
	app.drawToCanvas()

	// Draw canvas
	app.g.drawScaledBitmap(app.canvas, app.canvasPos.X, app.canvasPos.Y,
		app.canvasPos.W, app.canvasPos.H)
}

// Destroy
func (app *application) destroy() {

	app.window.Destroy()
}

// Terminate
func (app *application) terminate() {

	app.running = false
}

// Run
func (app *application) run() error {

	// Initialize scenes
	var err error
	var s scene
	for i := 0; i < len(app.scenes); i++ {

		s = app.scenes[i]
		err = s.init(app.g, app.trans, app.evMan, app.assets)
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

// Load assets (before running!)
func (app *application) loadAssets(bmpList, bmpNames, mapList, mapNames []string) error {

	var err error
	app.assets, err = createAssetPack(app.g, bmpList, bmpNames, mapList, mapNames)

	return err
}

// Add a key configuration
func (app *application) bindKeyConfig(kconf *keyConfig) {

	app.input.bindKeyConfig(kconf)
}

// Add a scene
func (app *application) addScene(s scene, makeCurrent bool) {

	app.scenes = append(app.scenes, s)
	if makeCurrent {
		app.currentScene = &s
	}
}
