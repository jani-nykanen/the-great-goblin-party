// Graphics routines
// (c) Jani Nykänen

package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Graphics class
type graphics struct {
	rend *sdl.Renderer
}

// Initialize
func (g *graphics) init(rend *sdl.Renderer) error {

	// Store renderer
	g.rend = rend

	return nil
}

// Clear screen
func (g *graphics) clearScreen(rc, gc, bc uint8) {

	g.rend.SetDrawColor(rc, gc, bc, 255)
	g.rend.Clear()
}
