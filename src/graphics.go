// Graphics routines
// (c) Jani Nykänen

package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	flipNone = 0
	flipH    = 1
	flipV    = 2
	flipBoth = 3
)

// Color type
type color struct {
	r, g, b, a uint8
}

// Graphics class
type graphics struct {
	rend        *sdl.Renderer
	globalColor color
	src         sdl.Rect
	dst         sdl.Rect
	tx, ty      int32
}

// Initialize
func (g *graphics) init(rend *sdl.Renderer) error {

	// Store renderer
	g.rend = rend

	// Set defaults
	g.globalColor = color{
		r: 255, g: 255, b: 255, a: 255,
	}
	g.tx = 0
	g.ty = 0

	return nil
}

// Set translation
func (g *graphics) translate(x, y int32) {

	g.tx = x
	g.ty = y
}

// Clear screen
func (g *graphics) clearScreen(rc, gc, bc uint8) {

	g.rend.SetDrawColor(rc, gc, bc, 255)
	g.rend.Clear()
}

// Set render target
func (g *graphics) setRenderTarget(bmp *bitmap) {

	var tex *sdl.Texture

	if bmp == nil {
		tex = nil
	} else {
		tex = bmp.texture
	}

	g.rend.SetRenderTarget(tex)
}

// Set global color for rendering primitives
func (g *graphics) setGlobalColor(rc, gc, bc, ac uint8) {

	g.globalColor = color{r: rc, g: gc, b: bc, a: ac}
}

// Draw a bitmap
func (g *graphics) drawBitmap(bmp *bitmap, dx, dy int32) {

	dx += g.tx
	dy += g.ty

	g.dst = sdl.Rect{X: dx, Y: dy, W: int32(bmp.width), H: int32(bmp.height)}
	g.rend.Copy(bmp.texture, nil, &g.dst)
}

// Draw a scaled bitmap
func (g *graphics) drawScaledBitmap(bmp *bitmap, dx, dy, dw, dh int32) {

	dx += g.tx
	dy += g.ty

	g.dst = sdl.Rect{X: dx, Y: dy, W: dw, H: dh}
	g.rend.Copy(bmp.texture, nil, &g.dst)
}

// Draw a bitmap region
func (g *graphics) drawBitmapRegion(bmp *bitmap, sx, sy, sw, sh, dx, dy int32, flip int) {

	dx += g.tx
	dy += g.ty

	g.src = sdl.Rect{X: sx, Y: sy, W: sw, H: sh}
	g.dst = sdl.Rect{X: dx, Y: dy, W: sw, H: sh}

	f := sdl.RendererFlip(flip)

	g.rend.CopyEx(bmp.texture, &g.src, &g.dst, 0.0, nil, f)
}

// Draw text
func (g *graphics) drawText(font *bitmap, text string, dx, dy int32, xoff, yoff int32, center bool) {

	l := len(text)

	x := dx
	y := dy
	cw := int32(font.width / 16)
	ch := cw
	var c uint8
	var sx, sy int32

	// Center text, if required
	if center {

		dx -= int32((float32(l) + 1) / 2.0 * float32(cw+xoff))
		x = dx
	}

	// Draw every character
	for i := 0; i < l; i++ {

		c = text[i]
		// Check if newline
		if c == '\n' {

			x = dx
			y += yoff + ch
			continue
		}

		sx = int32(c) % 16
		sy = int32(c) / 16

		g.drawBitmapRegion(font, sx*cw, sy*ch, cw, ch, x, y, flipNone)

		x += cw + xoff
	}
}

// Draw a filled rectangle
func (g *graphics) fillRect(dx, dy, dw, dh int32) {

	dx += g.tx
	dy += g.ty

	c := g.globalColor
	g.rend.SetDrawColor(c.r, c.g, c.b, c.a)

	g.dst = sdl.Rect{X: dx, Y: dy, W: dw, H: dh}
	g.rend.FillRect(&g.dst)
}
