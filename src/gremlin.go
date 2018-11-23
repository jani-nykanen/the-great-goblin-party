// Gremlin
// (c) Jani Nykänen

package main

// Constants
const (
	gremlimMoveTime = 15
)

// Gremlin type
type gremlin struct {
	x, y      int32
	tx, ty    int32
	vx, vy    float32
	moveTimer float32
	stopped   bool
	moving    bool
	spr       sprite
	color     int32
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

			gr.spr.row--
			gr.stopped = false
		}

		return
	}

	// Set destination
	gr.tx = gr.x + dx
	gr.ty = gr.y + dy

	// Check collision
	if s.isTileSolid(int(gr.tx), int(gr.ty)) {

		gr.tx = gr.x
		gr.ty = gr.y

		// Change animation back
		if gr.stopped {

			gr.spr.row--
			gr.stopped = false
		}

	} else {

		// Move
		gr.moveTimer = gremlimMoveTime
		gr.moving = true

		// Update solid data
		s.updateSolid(int(gr.x), int(gr.y), false)
		s.updateSolid(int(gr.tx), int(gr.ty), true)

		if !gr.stopped {
			gr.spr.row++
		}
	}
}

// Move
func (gr *gremlin) move(s *stage, tm float32) {

	// If not moving
	if !gr.moving {

		// Update solid data
		s.updateSolid(int(gr.x), int(gr.y), true)

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
	}

}

// Animate
func (gr *gremlin) animate(tm float32) {

	standSpeed := float32(10.0)
	moveSpeed := float32(8.0)

	speed := standSpeed
	if gr.moving {
		speed = moveSpeed
	}

	gr.spr.animate(gr.spr.row, 0, 3, speed, tm)
}

// Update
func (gr *gremlin) update(input *inputManager, s *stage, tm float32) {

	// Control
	gr.control(input, s, tm)
	// Move
	gr.move(s, tm)
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

	// Set defaults
	gr.moveTimer = 0.0
	gr.moving = false

	return gr
}
