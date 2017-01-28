package toys

import (
	hue "github.com/collinux/gohue"
	"github.com/kirsle/hue-toys/registry"
)

func init() {
	registry.Register(
		"on",
		"Turns on all lights to white.",
		func(c registry.RuntimeConfig) error {
			AllLights(c, func(light hue.Light) {
				light.SetBrightness(c.Brightness)
				light.SetColor(hue.WHITE)
				light.ColorLoop(false)
				light.On()
			})

			return nil
		},
	)
}
