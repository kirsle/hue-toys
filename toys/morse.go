package toys

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	hue "github.com/collinux/gohue"
	"github.com/kirsle/hue-toys/registry"
)

// Type signal is a morse code signal.
type signal bool

const (
	// Morse code signals.
	dot  signal = true
	dash signal = false
)

// The Morse code alphabet.
var alphabet = map[rune][]signal{
	'A': []signal{dot, dash},
	'B': []signal{dash, dot, dot, dot},
	'C': []signal{dash, dot, dash, dot},
	'D': []signal{dash, dot, dot},
	'E': []signal{dot},
	'F': []signal{dot, dot, dash, dot},
	'G': []signal{dash, dash, dot},
	'H': []signal{dot, dot, dot, dot},
	'I': []signal{dot, dot},
	'J': []signal{dot, dash, dash, dash},
	'K': []signal{dash, dot, dash},
	'L': []signal{dot, dash, dot, dot},
	'M': []signal{dash, dash},
	'N': []signal{dash, dot},
	'O': []signal{dash, dash, dash},
	'P': []signal{dot, dash, dash, dot},
	'Q': []signal{dash, dash, dot, dash},
	'R': []signal{dot, dash, dot},
	'S': []signal{dot, dot, dot},
	'T': []signal{dash},
	'U': []signal{dot, dot, dash},
	'V': []signal{dot, dot, dot, dash},
	'W': []signal{dot, dash, dash},
	'X': []signal{dash, dot, dot, dash},
	'Y': []signal{dash, dot, dash, dash},
	'Z': []signal{dash, dash, dot, dot},
	'1': []signal{dot, dash, dash, dash, dash},
	'2': []signal{dot, dot, dash, dash, dash},
	'3': []signal{dot, dot, dot, dash, dash},
	'4': []signal{dot, dot, dot, dot, dash},
	'5': []signal{dot, dot, dot, dot, dot},
	'6': []signal{dash, dot, dot, dot, dot},
	'7': []signal{dash, dash, dot, dot, dot},
	'8': []signal{dash, dash, dash, dot, dot},
	'9': []signal{dash, dash, dash, dash, dot},
	'0': []signal{dash, dash, dash, dash, dash},
}

func init() {
	registry.Register(
		"morse",
		"Spell out a message in morse code.\n"+
			"Use -delay to control the duration of a dot.\n"+
			"The dash duration is 3X the dot duration.",
		func(c registry.RuntimeConfig) error {
			message := flag.Args()
			if len(message) == 0 {
				fmt.Println("Usage: hue-toys -toy morse <message>")
				os.Exit(1)
			}

			// Get the brightness values to use.
			var (
				bright      = float64(c.Brightness) / 100.0
				dim         = float64(c.Brightness) * 0.1 / 100.0
				dotDelay    = c.Delay * time.Millisecond
				dashDelay   = c.Delay * 3 * time.Millisecond
				brightState = hue.LightState{
					On:             true,
					TransitionTime: "1",
					Bri:            uint8(254 * bright),
				}
				dimState = hue.LightState{
					On:             true,
					TransitionTime: "1",
					Bri:            uint8(254 * dim),
				}
			)

			fmt.Printf("Brightness levels: %f / %f\n", bright, dim)

			// Dim all the lights to begin with.
			fmt.Printf("Setting initial (dim) brightness level to %f%%\n", dim)
			AllLights(c, func(light hue.Light) {
				light.SetState(dimState)
			})

			// Split up the message.
			words := strings.Fields(strings.Join(message, " "))
			for _, word := range words {
				word = strings.ToUpper(word)
				fmt.Printf("Encoding: %s   ", word)
				for _, char := range word {
					// Look up the Morse code signal for this character.
					signal, ok := alphabet[char]
					if !ok {
						log.Printf("Warning: character '%s' has no morse code symbol.\n", string(char))
						continue
					}

					// Make the flashes.
					for _, flash := range signal {
						AllLights(c, func(light hue.Light) {
							light.SetState(brightState)
						})

						if flash == dot {
							fmt.Printf(".")
							time.Sleep(dotDelay)
						} else {
							fmt.Printf("-")
							time.Sleep(dashDelay)
						}

						AllLights(c, func(light hue.Light) {
							light.SetState(dimState)
						})
					}
				}

				fmt.Println("")
				time.Sleep(dotDelay)
			}

			return nil
		},
	)
}
