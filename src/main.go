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
		"assets/bitmaps/borders.png",
		"assets/bitmaps/wall.png",
		"assets/bitmaps/gremlin.png",
		"assets/bitmaps/numbers.png",
		"assets/bitmaps/button.png",
		"",
	}
	bmpNames := []string{
		"font",
		"borders",
		"wall",
		"gremlin",
		"numbers",
		"button",
	}
	mapPaths := []string{
		"assets/maps/1.tmx",
		"assets/maps/2.tmx",
		"assets/maps/3.tmx",
		"assets/maps/4.tmx",
		"assets/maps/5.tmx",
		"assets/maps/6.tmx",
		"assets/maps/7.tmx",
		"assets/maps/8.tmx",
		"assets/maps/9.tmx",
		"assets/maps/10.tmx",
		"assets/maps/11.tmx",
		"assets/maps/12.tmx",
		"",
	}
	mapNames := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"10",
		"11",
		"12",
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
	app.addScene(new(game), false)
	app.addScene(new(stageMenu), true)

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
