package toys

import (
	hue "github.com/collinux/gohue"
	"github.com/kirsle/hue-toys/registry"
)

// Disco puts all lights into ColorLoop mode.
func Disco(c registry.RuntimeConfig) error {
	AllLights(c, func(light hue.Light) {
		light.On()
		light.SetBrightness(c.Brightness)
		light.SetColor(RandomColor())
		light.ColorLoop(true)
	})

	return nil
}

func init() {
	registry.Register("disco", "Turns all lights onto a color loop.", Disco)
}
