// Sample
// (c) Jani Nykänen

package main

import (
	"github.com/veandco/go-sdl2/mix"
)

// Sample type
type sample struct {
	chunk   *mix.Chunk
	channel int
	played  bool
}

// Play sample
func (s *sample) play(vol float32, loops int) {

	v := int(float32(mix.MAX_VOLUME))

	if !s.played {

		// Get channel, halt, set volume, play again
		s.channel, _ = s.chunk.Play(-1, 0)
		mix.HaltChannel(s.channel)

		// Set volume and play
		mix.Volume(s.channel, v)
		s.chunk.Play(-1, loops)

		s.played = true

	} else {

		// Stop and play
		mix.HaltChannel(s.channel)
		mix.Volume(s.channel, v)
		s.chunk.Play(-1, loops)
	}
}

// Stop sample
func (s *sample) stop() {

	mix.HaltChannel(s.channel)
}

// Load a sample
func loadSample(path string) (*sample, error) {

	var err error
	s := new(sample)
	s.played = false
	s.channel = 0

	// Load file
	s.chunk, err = mix.LoadWAV(path)
	if err != nil {

		return nil, err
	}

	return s, err
}
