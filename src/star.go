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
func (st *star) animate(tm float32) {

	animSpeed := float32(8.0)

	st.spr.animate(2, 0, 3, animSpeed, tm)
}

// Update
func (st *star) update(input *inputManager, s *stage, tm float32) {

	// Update solid
	s.updateSolid(int(st.x), int(st.y), 1)

	// Animate
	st.animate(tm)
}

// Draw
func (st *star) draw(bmp *bitmap, g *graphics) {

	// Draw sprite
	st.spr.draw(g, bmp, int32(st.x)*16, int32(st.y)*16, flipNone)
}

// Create a star
func createStar(x, y, color int32) *star {

	s := new(star)

	// Set position
	s.x = x
	s.y = y

	// Create sprite
	s.spr = createSprite(16, 16)
	s.spr.row = color*5 + 2

	// Set color
	s.color = color

	return s
}
