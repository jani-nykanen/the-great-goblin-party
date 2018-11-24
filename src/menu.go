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
}

// Update
func (m *menu) update(input *inputManager) {

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

		// Call callback function, if defined
		if m.buttons[m.cursorPos].cb != nil {

			m.buttons[m.cursorPos].cb()
		}
	}

	// Check if cursor position changed
	if m.cursorPos != oldPos {
		// Play sound
	}
}

// Draw menu
func (m *menu) drawMenu(g *graphics, bmpFont *bitmap, dx, dy, yoff int32) {

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
		g.drawText(bmpFont, str, dx, dy+yoff*i, xoff, 0, false)
	}
}

// Create a menu
func createMenu(text []string, cbs []cbfun) *menu {

	m := new(menu)

	// Create buttons
	m.length = minInt(len(text), len(cbs))
	m.buttons = make([]menuButton, m.length)
	for i := 0; i < m.length; i++ {
		m.buttons[i] = menuButton{
			text[i], cbs[i],
		}
	}

	// Set cursor position
	m.cursorPos = 0

	return m
}
