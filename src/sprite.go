// Animated sprite
// (c) Jani Nykänen

package main

// Sprite type
type sprite struct {
	frame  int32
	row    int32
	count  float32
	width  int32
	height int32
}

// Constructor
func createSprite(w, h int32) sprite {

	return sprite{0, 0, 0.0, w, h}
}

// Animate a sprite
func (spr *sprite) animate(row, start, end int32, speed, tm float32) {

	if start == end {

		spr.count = 0
		spr.frame = start
		spr.row = row
		return
	}

	// Swap to the correct row
	if spr.row != row {

		spr.count = 0
		if end > start {
			spr.frame = start
		} else {
			spr.frame = end
		}
		spr.row = row
	}

	if start < end && spr.frame < start {

		spr.frame = start

	} else if end < start && spr.frame < end {

		spr.frame = end
	}

	// Animate
	spr.count += 1.0 * tm
	if spr.count > speed {

		if start < end {

			spr.frame++
			if spr.frame > end {
				spr.frame = start
			}

		} else {

			spr.frame--
			if spr.frame < end {

				spr.frame = start
			}
		}

		spr.count -= speed
	}
}

// Draw sprite frame
func (spr *sprite) drawFrame(g *graphics, bmp *bitmap, frame, row int32, dx, dy int32, flip int) {

	g.drawBitmapRegion(bmp, frame*spr.width, row*spr.height,
		spr.width, spr.height, dx, dy, flip)
}

// Draw sprite
func (spr *sprite) draw(g *graphics, bmp *bitmap, dx, dy int32, flip int) {

	g.drawBitmapRegion(bmp, spr.frame*spr.width, spr.row*spr.height,
		spr.width, spr.height, dx, dy, flip)
}
