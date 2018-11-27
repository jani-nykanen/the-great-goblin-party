// An asset pack
// (c) Jani Nykänen

package main

// Asset pack type
type assetPack struct {
	bitmaps     []*bitmap
	tilemaps    []*tilemap
	samples     []*sample
	bmpNames    []string
	mapNames    []string
	sampleNames []string
	bmpCount    int
	mapCount    int
	sampleCount int
}

// Get a bitmap
func (ass *assetPack) getBitmap(name string) *bitmap {

	for i := 0; i < ass.bmpCount; i++ {

		if ass.bmpNames[i] == name {
			return ass.bitmaps[i]
		}
	}

	return nil
}

// Get a tilemap
func (ass *assetPack) getTilemap(name string) *tilemap {

	for i := 0; i < ass.mapCount; i++ {

		if ass.mapNames[i] == name {
			return ass.tilemaps[i]
		}
	}

	return nil
}

// Get a sample
func (ass *assetPack) getSample(name string) *sample {

	for i := 0; i < ass.sampleCount; i++ {

		if ass.sampleNames[i] == name {
			return ass.samples[i]
		}
	}

	return nil
}

// Create an asset pack & load the files
func createAssetPack(g *graphics,
	bmpList, bmpNames,
	mapList, mapNames,
	sampleList, sampleNames []string) (*assetPack, error) {

	ass := new(assetPack)
	var err error

	// Create slices
	ass.bitmaps = make([]*bitmap, 0)
	ass.tilemaps = make([]*tilemap, 0)
	ass.samples = make([]*sample, 0)
	ass.bmpNames = make([]string, 0)
	ass.mapNames = make([]string, 0)
	ass.sampleNames = make([]string, 0)

	ass.bmpCount = 0
	ass.mapCount = 0
	ass.sampleCount = 0

	// Load bitmaps
	i := 0
	var bmp *bitmap
	for i = 0; i < minInt(len(bmpNames), len(bmpList)); i++ {

		bmp, err = loadBitmap(g, bmpList[i])
		if err != nil {

			return nil, err
		}
		ass.bmpCount++
		ass.bitmaps = append(ass.bitmaps, bmp)
		ass.bmpNames = append(ass.bmpNames, bmpNames[i])
	}

	// Load tilemaps
	var m *tilemap
	for i = 0; i < minInt(len(mapNames), len(mapList)); i++ {

		m, err = loadTilemap(mapList[i])
		if err != nil {

			return nil, err
		}
		ass.mapCount++
		ass.tilemaps = append(ass.tilemaps, m)
		ass.mapNames = append(ass.mapNames, mapNames[i])
	}

	// Load samples
	var s *sample
	for i = 0; i < minInt(len(sampleNames), len(sampleList)); i++ {

		s, err = loadSample(sampleList[i])
		if err != nil {

			return nil, err
		}
		ass.sampleCount++
		ass.samples = append(ass.samples, s)
		ass.sampleNames = append(ass.sampleNames, sampleNames[i])
	}

	return ass, err
}
