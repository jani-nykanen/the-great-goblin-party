// Save manager
// (c) Jani Nykänen

package main

import "io/ioutil"

// Save manager type
type saveManager struct {
	// Nothing at all...
	output []byte
}

// Save to a file
func (sm *saveManager) saveToFile(data []byte, path string) error {

	var err error
	err = ioutil.WriteFile(path, data, 0644)
	return err
}

// Load from a file
func (sm *saveManager) loadFromFile(path string) error {

	var err error
	sm.output, err = ioutil.ReadFile(path)
	return err
}

// Create save manager
func createSaveManager() *saveManager {

	sm := new(saveManager)
	return sm
}
