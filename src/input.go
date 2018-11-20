// inputManager manager
// (c) Jani Nykänen

package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Constants
const (
	stateUp       = 0
	stateDown     = 1
	statePressed  = 2
	stateReleased = 3
	keyMax        = 256
)

// Button
type button struct {
	key  sdl.Scancode
	name string
}

// Key configuration
type keyConfig struct {
	buttons []button
}

// Create a key configuration
func createKeyConfig() *keyConfig {

	k := new(keyConfig)
	k.buttons = make([]button, 0)

	return k
}

// Add a button
func (k *keyConfig) addButton(name string, key sdl.Scancode) {

	k.buttons = append(k.buttons, button{name: name, key: key})
}

// Input manager class
type inputManager struct {
	keys  [keyMax]int
	kconf *keyConfig
}

// Initialize
func (t *inputManager) init() {

	// Clear keyboard array
	for i := 0; i < keyMax; i++ {
		t.keys[i] = stateUp
	}
}

// Bind keyconfiguration
func (t *inputManager) bindKeyConfig(kconf *keyConfig) {

	t.kconf = kconf
}

// Key down
func (t *inputManager) keyDown(key int32) {

	if key < 0 || key >= keyMax || t.keys[key] == stateDown {
		return
	}

	t.keys[key] = statePressed
}

// Key up
func (t *inputManager) keyUp(key int32) {

	if key < 0 || key >= keyMax || t.keys[key] == stateUp {
		return
	}

	t.keys[key] = stateReleased
}

// Update keystates
func (t *inputManager) updateKeys() {

	for i := 0; i < keyMax; i++ {

		if t.keys[i] == statePressed {

			t.keys[i] = stateDown

		} else if t.keys[i] == stateReleased {

			t.keys[i] = stateUp
		}
	}
}

// Get key state
func (t *inputManager) getKey(key sdl.Scancode) int {

	i := int32(key)
	if i < 0 || i >= keyMax {
		return stateUp
	}

	return t.keys[i]
}

// Get button state
func (t *inputManager) getButton(name string) int {

	// Find a corresponding button
	for i := 0; i < len(t.kconf.buttons); i++ {

		if name == t.kconf.buttons[i].name {

			return t.getKey(t.kconf.buttons[i].key)
		}
	}

	return stateUp
}
