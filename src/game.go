// Game scene
// (c) Jani Nykänen

package main

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

// Game type
type game struct {
	pos     float32
	bmpGoat *bitmap
	bmpFont *bitmap
	sprGoat sprite
}

// Initialize
func (t *game) init(g *graphics) error {

	fmt.Println("Let's init the game!")

	t.pos = 0

	// Load test bitmaps
	var err error
	t.bmpGoat, err = loadBitmap(g, "assets/bitmaps/goat.png")
	if err != nil {

		return err
	}
	t.bmpFont, err = loadBitmap(g, "assets/bitmaps/font.png")
	if err != nil {

		return err
	}

	// Create goat sprite
	t.sprGoat = createSprite(32, 32)

	return nil
}

// Update
func (t *game) update(input *inputManager, tm float32) {

	// Test
	if input.getKey(sdl.SCANCODE_A) == statePressed {

		fmt.Println("beep boop")
	}

	// Update pos
	t.pos += 2.0 * tm
	t.pos = float32(math.Mod(float64(t.pos), 240))

	// Animate goat
	t.sprGoat.animate(0, 0, 7, 4.0, tm)
}

// Draw
func (t *game) draw(g *graphics) {

	g.clearScreen(85, 85, 85)

	g.setGlobalColor(255, 0, 0, 255)
	g.fillRect(int32(t.pos), int32(t.pos), 32, 32)
	t.sprGoat.draw(g, t.bmpGoat, int32(t.pos), 8, flipNone)
	t.sprGoat.draw(g, t.bmpGoat, int32(t.pos), int32(t.pos), flipH)

	// Draw goat
	t.sprGoat.draw(g, t.bmpGoat, 8, 8, flipBoth)

	// Hello world!
	g.drawText(t.bmpFont, "Hello world!", 128, 8, -1, 0, true)
}

// Destroy
func (t *game) destroy() {

	fmt.Println("Let's destroy the game!")
}

// Scene changed
func (t *game) onChange() {

}
