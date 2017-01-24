package toys

import (
	"log"
	"time"

	hue "github.com/collinux/gohue"
	"github.com/kirsle/hue-toys/registry"
)

// Rainbow puts all lights in sync with rainbow colors.
func Rainbow(c registry.RuntimeConfig) error {
	lights, err := c.Bridge.GetAllLights()
	if err != nil {
		return err
	}

	colors := []*[2]float32{
		hue.RED,
		hue.ORANGE,
		hue.YELLOW,
		hue.GREEN,
		hue.CYAN,
		hue.BLUE,
		hue.PINK,
	}

	// Turn on all lights to the right level.
	log.Printf("Turning on all lights at brightness %d\n", c.Brightness)
	for _, light := range lights {
		light.SetBrightness(c.Brightness)
		light.On()
	}

	for {
		for _, color := range colors {
			for _, light := range lights {
				light.SetColor(color)
			}
			time.Sleep(1000 * time.Millisecond)
			log.Printf("Set color to: %v\n", color)
		}
	}

	return nil
}

func init() {
	registry.Register("rainbow", "Puts all lights in sync with rainbow colors (animated).", Rainbow)
}
