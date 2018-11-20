// Bitmap
// (c) Jani Nykänen

package main

import (
	"image"
	_ "image/png"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

// Bitmap type
type bitmap struct {
	width   int
	height  int
	texture *sdl.Texture
}

// Create an empty bitmap
func createEmptyBitmap(g *graphics, w, h int32, target bool) (*bitmap, error) {

	bmp := new(bitmap)

	var access int
	if target {
		access = sdl.TEXTUREACCESS_TARGET

	} else {
		access = sdl.TEXTUREACCESS_STATIC
	}

	bmp.width = int(w)
	bmp.height = int(h)

	var err error
	bmp.texture, err = g.rend.CreateTexture(sdl.PIXELFORMAT_RGB332, access, w, h)

	return bmp, err
}

// Convert "raw" pixel data to uint8 rgba
func rgbaRawToBitmap(r, g, b, a uint32) (uint8, uint8, uint8, uint8) {

	return uint8(r / 257), uint8(g / 257), uint8(b / 257), uint8(a / 257)
}

// Load a bitmap
func loadBitmap(g *graphics, path string) (*bitmap, error) {

	bmp := new(bitmap)
	var err error

	// Open the file
	file, err := os.Open(path)
	if err != nil {

		return nil, err
	}

	// Decode
	data, _, err := image.Decode(file)
	if err != nil {

		return nil, err
	}

	bmp.width = data.Bounds().Max.X
	bmp.height = data.Bounds().Max.Y
	i := 0

	rmask := uint32(0x000000ff)
	gmask := uint32(0x0000ff00)
	bmask := uint32(0x00ff0000)
	amask := uint32(0xff000000)

	// Create surface
	surf, err := sdl.CreateRGBSurface(0, int32(bmp.width), int32(bmp.height), 32, rmask, gmask, bmask, amask)
	if err != nil {
		return nil, err
	}
	pdata := surf.Pixels()

	// Get pixels in uint8 format
	for y := 0; y < bmp.height; y++ {

		for x := 0; x < bmp.width; x++ {

			pdata[i], pdata[i+1], pdata[i+2], pdata[i+3] =
				rgbaRawToBitmap(data.At(x, y).RGBA())
			i += 4
		}
	}

	// Create a texture
	bmp.texture, err = g.rend.CreateTextureFromSurface(surf)
	if err != nil {
		return nil, err
	}

	return bmp, err
}
