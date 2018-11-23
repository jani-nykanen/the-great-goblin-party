// Star
// (c) Jani Nykänen

package main

// Star type
type star struct {
	x, y   int32
	vx, vy float32
	spr    sprite
	color  int32
}

// Animate
func (s *star) animate(tm float32) {

	animSpeed := float32(8.0)

	s.spr.animate(2, 0, 3, animSpeed, tm)
}

// Update
func (s *star) update(input *inputManager, tm float32) {

	// Animate
	s.animate(tm)
}

// Draw
func (s *star) draw(bmp *bitmap, g *graphics) {

	// Draw sprite
	s.spr.draw(g, bmp, int32(s.x)*16, int32(s.y)*16, flipNone)
}

// Create a star
func createStar(x, y, color int32) *star {

	s := new(star)

	// Set position
	s.x = x
	s.y = y

	// Create sprite
	s.spr = createSprite(16, 16)

	// Set color
	s.color = color

	return s
}
