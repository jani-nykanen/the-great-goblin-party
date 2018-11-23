// Utilities
// (c) Jani Nykänen

package main

// Int min
func minInt(a, b int) int {

	if a < b {
		return a
	}
	return b
}

// Int max
func maxInt(a, b int) int {

	if a > b {
		return a
	}
	return b
}

// Int abs
func absInt(a int) int {

	if a < 0 {
		return -a
	}
	return a
}
