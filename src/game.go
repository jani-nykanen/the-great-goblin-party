// Game scene
// (c) Jani Nykänen

package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Game type
type game struct {
}

// Initialize
func (t *game) init() error {

	fmt.Println("Let's init the game!")

	return nil
}

// Update
func (t *game) update(input *inputManager, tm float32) {

	// Test
	if input.getKey(sdl.SCANCODE_A) == statePressed {

		fmt.Println("beep boop")
	}
}

// Draw
func (t *game) draw(g *graphics) {

	g.clearScreen(170, 170, 170)
}

// Destroy
func (t *game) destroy() {

	fmt.Println("Let's destroy the game!")
}

// Scene changed
func (t *game) onChange() {

}
