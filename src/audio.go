// Audio manager
// (c) Jani Nykänen

package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/mix"
)

// Audio manager type
type audioManager struct {
	vol        float32
	loopSample *sample
	enabled    bool
}

// Play a sample once
func (a *audioManager) playSample(s *sample, vol float32) {

	if !a.enabled {
		return
	}

	s.play(vol*a.vol, 0)
}

// Play music (aka: loop sample
func (a *audioManager) playMusic(s *sample, vol float32) {

	if !a.enabled || a.loopSample != nil {
		return
	}

	a.loopSample = s

	s.play(vol*a.vol, -1)
}

// Stop music
func (a *audioManager) stopMusic() {

	if !a.enabled {
		return
	}

	if a.loopSample != nil {

		a.loopSample.stop()
		a.loopSample = nil
	}
}

// Create audio manager
func createAudioManager(vol float32) *audioManager {

	a := new(audioManager)
	a.vol = vol

	// Open & init audio
	mix.Init(0)

	// Open audio
	a.enabled = true
	err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 512)
	if err != nil {

		fmt.Println("Failed to open audio with error:")
		fmt.Println(err)
		a.enabled = false
	}

	return a

}
