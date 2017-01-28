package toys

import (
	hue "github.com/collinux/gohue"
	"github.com/kirsle/hue-toys/registry"
)

func init() {
	registry.Register(
		"noop",
		"Don't do anything (useful if you just want to set the brightness with -b)",
		func(c registry.RuntimeConfig) error {
			AllLights(c, func(light hue.Light) {
				light.SetBrightness(c.Brightness)
			})
			return nil
		},
	)
}
