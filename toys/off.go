package toys

import (
	hue "github.com/collinux/gohue"
	"github.com/kirsle/hue-toys/registry"
)

// Off turns all lights off.
func Off(c registry.RuntimeConfig) error {
	AllLights(c, func(light hue.Light) {
		light.Off()
	})

	return nil
}

func init() {
	registry.Register("off", "Turns off all lights.", Off)
}
