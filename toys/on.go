package toys

import (
	hue "github.com/collinux/gohue"
	"github.com/kirsle/hue-toys/registry"
)

// On turns all lights on.
func On(c registry.RuntimeConfig) error {
	AllLights(c, func(light hue.Light) {
		light.SetBrightness(c.Brightness)
		light.SetColor(hue.WHITE)
		light.ColorLoop(false)
		light.On()
	})

	return nil
}

func init() {
	registry.Register("on", "Turns on all lights to white.", On)
}
