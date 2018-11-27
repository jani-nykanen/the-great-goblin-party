// A simple menu
// (c) Jani Nykänen

package main

// Button type
type menuButton struct {
	text string
	cb   cbfun
}

// Menu type
type menu struct {
	buttons   []menuButton
	cursorPos int
	length    int
	bmpFont   *bitmap
	sSelect   *sample
	sAccept   *sample
}

// Update
func (m *menu) update(input *inputManager, audio *audioManager) {

	// Update cursor pos
	oldPos := m.cursorPos
	if input.getButton("down") == statePressed {
		m.cursorPos++

	} else if input.getButton("up") == statePressed {

		m.cursorPos--
	}

	// Restrict
	if m.cursorPos < 0 {
		m.cursorPos = 0

	} else if m.cursorPos >= m.length {

		m.cursorPos = m.length - 1
	}

	// Check if enter pressed
	if input.getButton("start") == statePressed {

		audio.playSample(m.sAccept, 0.30)

		// Call callback function, if defined
		if m.buttons[m.cursorPos].cb != nil {

			m.buttons[m.cursorPos].cb()
		}
	}

	// Check if cursor position changed
	if m.cursorPos != oldPos {

		audio.playSample(m.sSelect, 0.30)
	}
}

// Draw menu
func (m *menu) drawMenu(g *graphics, dx, dy, yoff int32) {

	xoff := int32(-6)

	var str string
	for i := int32(0); i < int32(m.length); i++ {

		// Make string to be drawn
		str = ""
		if i == int32(m.cursorPos) {
			str += "&"
		}
		str += m.buttons[i].text

		// Draw "button"
		g.drawText(m.bmpFont, str, dx, dy+yoff*i, xoff, 0, false)
	}
}

// Create a menu
func createMenu(text []string, cbs []cbfun, ass *assetPack) *menu {

	m := new(menu)

	// Create buttons
	m.length = minInt(len(text), len(cbs))
	m.buttons = make([]menuButton, m.length)
	for i := 0; i < m.length; i++ {
		m.buttons[i] = menuButton{
			text[i], cbs[i],
		}
	}

	// Get assets
	m.bmpFont = ass.getBitmap("font")
	m.sAccept = ass.getSample("accept")
	m.sSelect = ass.getSample("select")

	// Set cursor position
	m.cursorPos = 0

	return m
}
