// Game object
// (c) Jani Nykänen

package main

// Object base interface
type object interface {
	init(ass *assetPack)
	update(input *inputManager, tm float32)
	draw(g *graphics)
}
