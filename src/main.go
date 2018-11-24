package main

import (
	"log"

	"github.com/veandco/go-sdl2/sdl"
)

// Main file
// (c) Jani Nykänen

// Main
func main() {

	// Asset lists
	// (yeah, we do not bother loading them from a file
	// this time/yet)
	bmpPaths := []string{
		"assets/bitmaps/font.png",
		"assets/bitmaps/goat.png",
		"assets/bitmaps/borders.png",
		"assets/bitmaps/wall.png",
		"assets/bitmaps/gremlin.png",
		"",
	}
	bmpNames := []string{
		"font",
		"goat",
		"borders",
		"wall",
		"gremlin",
	}
	mapPaths := []string{
		"assets/maps/1.tmx",
	}
	mapNames := []string{
		"1",
	}

	// Key configuration
	kconf := createKeyConfig()
	kconf.addButton("right", sdl.SCANCODE_RIGHT)
	kconf.addButton("up", sdl.SCANCODE_UP)
	kconf.addButton("left", sdl.SCANCODE_LEFT)
	kconf.addButton("down", sdl.SCANCODE_DOWN)
	kconf.addButton("start", sdl.SCANCODE_RETURN)
	kconf.addButton("cancel", sdl.SCANCODE_ESCAPE)
	kconf.addButton("restart", sdl.SCANCODE_R)

	// Create default configuration
	conf := config{
		caption:      "The Great Gremlin Party",
		winWidth:     512,
		winHeight:    480,
		canvasWidth:  256,
		canvasHeight: 240,
		frameRate:    30,
		fullscreen:   false,
	}

	// Initialize application
	app := new(application)
	err := app.init(conf)
	if err != nil {
		log.Fatal(err)
		return
	}
	// Add key configuration
	app.bindKeyConfig(kconf)
	// Add scenes
	g := new(game)
	app.addScene(g, true)

	// Load assets
	err = app.loadAssets(bmpPaths, bmpNames, mapPaths, mapNames)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Run application
	err = app.run()
	if err != nil {
		log.Fatal(err)
		return
	}
}
