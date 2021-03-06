package toys

import (
	hue "github.com/collinux/gohue"
	"github.com/kirsle/hue-toys/registry"
)

func init() {
	registry.Register(
		"random",
		"Pick a random color for each light.",
		func(c registry.RuntimeConfig) error {
			AllLights(c, func(light hue.Light) {
				light.On()
				light.SetBrightness(c.Brightness)
				light.SetColor(RandomColor())
			})

			return nil
		},
	)
}
