package toys

import (
	hue "github.com/collinux/gohue"
	"github.com/kirsle/hue-toys/registry"
)

// Random picks random colors for each light.
func Random(c registry.RuntimeConfig) error {
	AllLights(c, func(light hue.Light) {
		light.On()
		light.SetBrightness(c.Brightness)
		light.SetColor(RandomColor())
	})

	return nil
}

func init() {
	registry.Register("random", "Pick a random color for each light.", Random)
}
