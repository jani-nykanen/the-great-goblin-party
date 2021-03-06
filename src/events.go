// Event manager
// Handles termination event and possible
// some other events. We don't want to
// directly pass application reference to
// other objects.
// (c) Jani Nykänen

package main

// Event manager type
type eventManager struct {
	app *application
}

// Terminate application
func (ev *eventManager) terminate() {

	ev.app.terminate()
}

// Change scene
func (ev *eventManager) changeScene(param int, name string) {

	ev.app.changeScene(param, name)
}

// Create an event manager
func createEventManager(app *application) *eventManager {

	ev := new(eventManager)
	ev.app = app

	return ev
}
