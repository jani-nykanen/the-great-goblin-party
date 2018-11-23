// Gremlin
// (c) Jani Nykänen

package main

// Gremlin type
type gremlin struct {
	x, y   int32
	vx, vy float32
	spr    sprite
	color  int32
}

// Animate
func (gr *gremlin) animate(tm float32) {

	standSpeed := float32(10.0)

	gr.spr.animate(0, 0, 3, standSpeed, tm)
}

// Update
func (gr *gremlin) update(input *inputManager, tm float32) {

	// Animate
	gr.animate(tm)
}

// Draw
func (gr *gremlin) draw(bmp *bitmap, g *graphics) {

	// Draw sprite
	gr.spr.draw(g, bmp, int32(gr.vx), int32(gr.vy), flipNone)
}

// Create a gremlin
func createGremlin(x, y, color int32) *gremlin {

	gr := new(gremlin)

	// Set position
	gr.x = x
	gr.y = y
	gr.vx = float32(x) * 16.0
	gr.vy = float32(y) * 16.0

	// Create sprite
	gr.spr = createSprite(16, 16)

	// Set color
	gr.color = color

	return gr
}
