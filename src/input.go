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

// Input manager class
type inputManager struct {
	keys [keyMax]int
}

// Initialize
func (t *inputManager) init() {

	// Clear keyboard array
	for i := 0; i < keyMax; i++ {
		t.keys[i] = stateUp
	}
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
