// Gremlin
// (c) Jani Nykänen

package main

// Constants
const (
	gremlimMoveTime = 15
)

// Gremlin type
type gremlin struct {
	x, y          int32
	tx, ty        int32
	vx, vy        float32
	moveTimer     float32
	startedMoving bool
	collisionSet  bool
	stopped       bool
	moving        bool
	spr           sprite
	color         int32
	exist         bool
	dying         bool
	sleeping      bool
	startRow      int32
}

// Find a free tile
func (gr *gremlin) findFreeTile(dx, dy int32, s *stage) bool {

	x := gr.x
	y := gr.y
	var tileID int

	for {

		x += dx
		y += dy

		tileID = s.isTileSolid(int(x), int(y))
		if tileID == 0 {
			return true

		} else if tileID == 1 {
			break
		}
	}

	return false
}

// Control
func (gr *gremlin) control(input *inputManager, s *stage, tm float32) {

	// If something moving, do not control
	if s.anyMoving {
		return
	}

	// Get moving direction
	dx := int32(0)
	dy := int32(0)
	if input.getButton("right") == stateDown {
		dx = 1

	} else if input.getButton("up") == stateDown {
		dy = -1

	} else if input.getButton("left") == stateDown {
		dx = -1

	} else if input.getButton("down") == stateDown {
		dy = 1
	}

	// If no direction chosen, stop
	if dx == 0 && dy == 0 {

		// Change animation back
		if gr.stopped {

			gr.spr.row = gr.startRow
			gr.stopped = false
		}

		return
	}

	// Set destination
	gr.tx = gr.x + dx
	gr.ty = gr.y + dy

	// Check if free
	if !gr.findFreeTile(dx, dy, s) {

		gr.tx = gr.x
		gr.ty = gr.y

		// Change animation row
		gr.spr.row = gr.startRow

		return
	}

	// Move
	gr.moveTimer = gremlimMoveTime
	gr.moving = true
	gr.startedMoving = true
	gr.collisionSet = false
	gr.spr.row = gr.startRow + 1

	// Update solid data
	s.updateSolid(int(gr.x), int(gr.y), 0)
	// s.updateSolid(int(gr.tx), int(gr.ty), 2)
}

// Move
func (gr *gremlin) move(s *stage, tm float32) {

	// If not moving
	if !gr.moving {

		// Update virtual position when
		// standing
		gr.vx = float32(gr.x) * 16.0
		gr.vy = float32(gr.y) * 16.0

		return

	}

	// Compute virtual position when
	// moving
	t := gr.moveTimer / gremlimMoveTime
	gr.vx = float32(gr.x*16)*t + (1-t)*float32(gr.tx*16)
	gr.vy = float32(gr.y*16)*t + (1-t)*float32(gr.ty*16)

	// Update move timer
	gr.moveTimer -= 1.0 * tm
	if gr.moveTimer <= 0.0 {

		gr.moveTimer = 0.0
		gr.moving = false
		gr.stopped = true

		// Set to the new position
		gr.x = gr.tx
		gr.y = gr.ty

		// Update solid data
		s.updateSolid(int(gr.x), int(gr.y), 2)

	}

}

// Animate
func (gr *gremlin) animate(tm float32) {

	standSpeed := float32(10.0)
	moveSpeed := float32(8.0)

	// If boulder
	if gr.color == 3 {
		gr.spr.row = 3 * 5
		gr.spr.frame = 0
		if gr.moving {
			gr.spr.frame = 1
		}
		return
	}

	speed := standSpeed
	if gr.moving {
		speed = moveSpeed
	}

	gr.spr.animate(gr.spr.row, 0, 3, speed, tm)
}

// Die
func (gr *gremlin) die(s *stage, tm float32) {

	animSpeed := float32(8)

	// Update virtual position
	gr.vx = float32(gr.x) * 16.0
	gr.vy = float32(gr.y) * 16.0

	// Animate
	gr.spr.animate(gr.startRow+3, 0, 4, animSpeed, tm)
	if gr.spr.frame == 4 {
		gr.exist = false
		s.addStar(gr.x, gr.y, gr.color)
	}
}

// Sleep
func (gr *gremlin) sleep(s *stage, tm float32) {

	sleepSpeed := float32(60.0)

	// Animate
	gr.spr.animate(gr.startRow+4, 0, 3, sleepSpeed, tm)

	// Set solid
	s.updateSolid(int(gr.x), int(gr.y), 1)
}

// Is active
func (gr *gremlin) isActive() bool {

	return gr.exist && (gr.dying || gr.moving)
}

// Update
func (gr *gremlin) update(input *inputManager, s *stage, tm float32) {

	if !gr.exist {
		return
	}

	gr.startedMoving = false

	// Die
	if gr.dying {
		gr.die(s, tm)
		return
	}

	// Sleep
	if gr.sleeping {
		gr.sleep(s, tm)
		return
	}

	// Control
	gr.control(input, s, tm)
	// Move
	gr.move(s, tm)
	// Animate
	gr.animate(tm)
}

// Check star collision
func (gr *gremlin) getStarCollision(st *star, s *stage) {

	if gr.color == 3 || gr.moving || gr.dying || !gr.exist || st.color != gr.color {
		return
	}

	// Check if near a star
	m := absInt(int(st.x-gr.x)) + absInt(int(st.y-gr.y))
	if m == 1 {
		gr.dying = true
		gr.spr.frame = 0
		gr.spr.row = 3
	}
}

// Draw
func (gr *gremlin) draw(bmp *bitmap, g *graphics) {

	if !gr.exist {
		return
	}

	// Draw sprite
	gr.spr.draw(g, bmp, int32(gr.vx), int32(gr.vy), flipNone)
}

// Create a gremlin
func createGremlin(x, y, color int32, s *stage, sleeping bool) *gremlin {

	gr := new(gremlin)

	// Set position
	gr.x = x
	gr.y = y
	gr.vx = float32(x) * 16.0
	gr.vy = float32(y) * 16.0
	solid := 2
	if sleeping {
		solid = 1
	}
	s.updateSolid(int(x), int(y), solid)
	// Is sleeping
	gr.sleeping = sleeping

	// Create sprite
	gr.spr = createSprite(16, 16)
	gr.spr.row = color * 5
	gr.spr.frame = 0
	if sleeping {
		gr.spr.row += 4
	}

	// Set color
	gr.color = color
	gr.startRow = color * 5

	// Set defaults
	gr.moveTimer = 0.0
	gr.moving = false
	gr.exist = true
	gr.dying = false

	return gr
}
