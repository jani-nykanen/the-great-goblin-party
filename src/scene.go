// Scene type
// (c) Jani Nykänen

package main

// Scene interface
type scene interface {
	init(g *graphics) error
	update(input *inputManager, tm float32)
	draw(g *graphics)
	destroy()
	onChange()
}