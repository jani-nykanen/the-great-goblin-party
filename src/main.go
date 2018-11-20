package main

import (
	"log"
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
	}
	bmpNames := []string{
		"font",
		"goat",
	}
	mapPaths := ([]string)(nil)
	mapNames := ([]string)(nil)

	// Create default configuration
	conf := config{
		caption:      "Go Experiment",
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
