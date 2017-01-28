package toys

import (
	"math/rand"
	"sort"
	"time"

	hue "github.com/collinux/gohue"
	"github.com/kirsle/hue-toys/registry"
)

// AllLights runs a function against all lights in the bridge.
func AllLights(c registry.RuntimeConfig, fn func(hue.Light)) error {
	lights, err := c.Bridge.GetAllLights()
	if err != nil {
		return err
	}

	// The lights given by the Go API are given in random order; for
	// predictability sort them by their indexes. For lookup later, group
	// the lights by index too.
	byIndex := map[int][]hue.Light{}
	indices := []int{}
	for _, light := range lights {
		if _, ok := byIndex[light.Index]; !ok {
			byIndex[light.Index] = []hue.Light{}
		}
		indices = append(indices, light.Index)
		byIndex[light.Index] = append(byIndex[light.Index], light)
	}
	sort.Ints(indices)

	// Run their functions in order.
	for _, index := range indices {
		for _, light := range byIndex[index] {
			fn(light)
		}
	}

	return nil
}

var randomColorIndex int
var HueColors = []*[2]float32{
	hue.RED,
	hue.YELLOW,
	hue.ORANGE,
	hue.GREEN,
	hue.CYAN,
	hue.BLUE,
	hue.PURPLE,
	hue.PINK,
}

// Common helper functions.
func init() {
	// Seed the random number generator.
	rand.Seed(time.Now().Unix())

	// Shuffle the array of valid colors.
	for i := range HueColors {
		j := rand.Intn(i + 1)
		HueColors[i], HueColors[j] = HueColors[j], HueColors[i]
	}
}

// RandomColor returns a random named color, with no duplicates until the list
// of colors has been exhausted.
func RandomColor() *[2]float32 {
	color := HueColors[randomColorIndex]
	randomColorIndex++
	if randomColorIndex == len(HueColors) {
		randomColorIndex = 0
	}
	return color
}
