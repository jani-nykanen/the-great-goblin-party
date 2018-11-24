// Transitions
// (c) Jani Nykänen

package main

const (
	fadeIn   = 0
	fadeOut  = 1
	fadeTime = 60
)

// Transition type
type transition struct {
	timer  float32
	active bool
	mode   int
	speed  float32
	cb     cbfun
}

// Activate
func (tr *transition) activate(mode int, speed float32, cb cbfun) {

	tr.active = true
	tr.mode = mode
	tr.speed = speed
	tr.timer = fadeTime
	tr.cb = cb
}

// Update
func (tr *transition) update(tm float32) {

	if !tr.active {
		return
	}

	// Update timer
	tr.timer -= tr.speed * tm
	if tr.timer <= 0.0 {

		tr.timer = 0.0
		if tr.mode == fadeIn {

			// Call callback function
			if tr.cb != nil {
				tr.cb()
			}
			tr.timer = fadeTime
			tr.mode = fadeOut

		} else {
			tr.active = false
		}
	}
}

// Draw
func (tr *transition) draw(g *graphics, cw, ch int32) {

	if !tr.active {
		return
	}

	g.translate(0, 0)
	g.setGlobalColor(0, 0, 0, 255)

	t := tr.timer / fadeTime
	if tr.mode == fadeIn {
		t = 1.0 - t
	}
	w := int32(t * float32(cw/2))
	h := int32(t * float32(ch/2))

	// Split to 8x8 blocks
	w = w / 8
	w *= 8
	h = h / 8
	h *= 8

	// Draw black thing
	g.fillRect(0, 0, cw, h)
	g.fillRect(0, ch-h, cw, h)
	g.fillRect(0, 0, w, ch)
	g.fillRect(cw-w, 0, w, ch)
}

// Create
func createTransition() *transition {

	tr := new(transition)

	// Set defaults
	tr.timer = 0.0
	tr.active = false
	tr.mode = fadeIn
	tr.cb = nil

	return tr
}
