// Place-holder tilemap routines.
// Works only with Tiled maps with one layer
// (c) Jani Nykänen

package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// Tilemap type
type tilemap struct {
	data      []int
	width     int
	height    int
	moveLimit int
}

// Get string parameter value
func (t tilemap) getStringParamValue(src, param string) string {

	pos := strings.Index(src, param) + len(param) + 2
	end := -1
	ret := ""

	for i := pos; i < len(src); i++ {

		if src[i] == '"' {

			end = i
			break
		}
	}

	if end != -1 {
		ret = src[pos:end]
	}

	return ret
}

// Just like above, but with integers
func (t tilemap) getIntParamValue(src, param string) int {

	// Convert to int (and ignore the error!)
	i, _ := strconv.Atoi(t.getStringParamValue(src, param))
	return i
}

// Get CSV content
func (t tilemap) getCSV(src, param string) string {

	ret := ""
	pos := strings.Index(src, param) + len(param) + 1
	end := -1

	for i := pos; i < len(src); i++ {

		if src[i] == '<' {

			end = i
			break
		}
	}

	if end != -1 {
		ret = src[pos:end]
	}

	ret = strings.Replace(strings.Replace(ret, "\n", "", -1), "\r", "", -1)

	return ret
}

// Parse CSV to integer array
func (t tilemap) parseCSV(src string) []int {

	words := strings.Split(src, ",")
	ret := make([]int, len(words))
	for i := 0; i < len(ret); i++ {

		ret[i], _ = strconv.Atoi(words[i])
	}

	return ret
}

// Load a tilemap from a file (with max 1 layer)
func loadTilemap(path string) (*tilemap, error) {

	t := new(tilemap)

	// Load file content to a string
	bytes, err := ioutil.ReadFile(path)
	if err != nil {

		return nil, err
	}
	content := string(bytes)

	// Get size
	t.width = t.getIntParamValue(content, "width")
	t.height = t.getIntParamValue(content, "height")

	// Get CSV (=layer) data
	csv := t.getCSV(content, "encoding=\"csv\"")
	t.data = t.parseCSV(csv)

	// Get move limit
	t.moveLimit = t.getIntParamValue(content, "property name=\"moves\" value")

	return t, nil
}
