package toys

import (
	"math/rand"
	"time"

	hue "github.com/collinux/gohue"
	"github.com/kirsle/hue-toys/registry"
)

// Common helper functions.
func init() {
	rand.Seed(time.Now().Unix())
}

// AllLights runs a function against all lights in the bridge.
func AllLights(c registry.RuntimeConfig, fn func(hue.Light)) error {
	lights, err := c.Bridge.GetAllLights()
	if err != nil {
		return err
	}

	for _, light := range lights {
		fn(light)
	}

	return nil
}

// RandomColor returns a random named color.
func RandomColor() *[2]float32 {
	options := []*[2]float32{
		hue.RED,
		hue.YELLOW,
		hue.ORANGE,
		hue.GREEN,
		hue.CYAN,
		hue.BLUE,
		hue.PURPLE,
		hue.PINK,
	}
	return options[rand.Intn(len(options))]
}
