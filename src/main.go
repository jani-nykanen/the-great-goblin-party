package main

import (
	"log"
)

// Main file
// (c) Jani Nykänen

// Main
func main() {

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
	}
	// Add scenes
	g := new(game)
	app.addScene(g, true)

	err = app.run()
	if err != nil {
		log.Fatal(err)
	}
}
