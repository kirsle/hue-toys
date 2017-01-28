package toys

import (
	hue "github.com/collinux/gohue"
	"github.com/kirsle/hue-toys/registry"
)

func init() {
	registry.Register(
		"off",
		"Turns off all lights.",
		func(c registry.RuntimeConfig) error {
			AllLights(c, func(light hue.Light) {
				light.Off()
			})

			return nil
		},
	)
}
