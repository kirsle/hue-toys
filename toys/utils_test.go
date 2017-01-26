package toys_test

import (
	"testing"

	"github.com/kirsle/hue-toys/toys"
)

func TestRandomColors(t *testing.T) {
	// Should be able to select more random colors than we have colors.
	maxColors := len(toys.HueColors)

	// We'll do two passes of exhausting all the valid colors. This ensures that
	// no overflow causes index errors, and we get unique colors each round.
	for a := 0; a < 2; a++ {
		// Keep track of unique colors this round.
		unique := map[*[2]float32]bool{}
		for i := 0; i < maxColors; i++ {
			next := toys.RandomColor()
			if _, ok := unique[next]; ok {
				t.Errorf("Got a duplicate random color in the first pass! (%d)", i)
			}
			unique[next] = true
		}
	}
}
