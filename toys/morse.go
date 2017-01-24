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

// Type Signal is a morse code signal.
type Signal bool

const (
	// Morse code signals.
	DOT  Signal = true
	DASH Signal = false
)

// The Morse code alphabet.
var Alphabet = map[rune][]Signal{
	'A': []Signal{DOT, DASH},
	'B': []Signal{DASH, DOT, DOT, DOT},
	'C': []Signal{DASH, DOT, DASH, DOT},
	'D': []Signal{DASH, DOT, DOT},
	'E': []Signal{DOT},
	'F': []Signal{DOT, DOT, DASH, DOT},
	'G': []Signal{DASH, DASH, DOT},
	'H': []Signal{DOT, DOT, DOT, DOT},
	'I': []Signal{DOT, DOT},
	'J': []Signal{DOT, DASH, DASH, DASH},
	'K': []Signal{DASH, DOT, DASH},
	'L': []Signal{DOT, DASH, DOT, DOT},
	'M': []Signal{DASH, DASH},
	'N': []Signal{DASH, DOT},
	'O': []Signal{DASH, DASH, DASH},
	'P': []Signal{DOT, DASH, DASH, DOT},
	'Q': []Signal{DASH, DASH, DOT, DASH},
	'R': []Signal{DOT, DASH, DOT},
	'S': []Signal{DOT, DOT, DOT},
	'T': []Signal{DASH},
	'U': []Signal{DOT, DOT, DASH},
	'V': []Signal{DOT, DOT, DOT, DASH},
	'W': []Signal{DOT, DASH, DASH},
	'X': []Signal{DASH, DOT, DOT, DASH},
	'Y': []Signal{DASH, DOT, DASH, DASH},
	'Z': []Signal{DASH, DASH, DOT, DOT},
	'1': []Signal{DOT, DASH, DASH, DASH, DASH},
	'2': []Signal{DOT, DOT, DASH, DASH, DASH},
	'3': []Signal{DOT, DOT, DOT, DASH, DASH},
	'4': []Signal{DOT, DOT, DOT, DOT, DASH},
	'5': []Signal{DOT, DOT, DOT, DOT, DOT},
	'6': []Signal{DASH, DOT, DOT, DOT, DOT},
	'7': []Signal{DASH, DASH, DOT, DOT, DOT},
	'8': []Signal{DASH, DASH, DASH, DOT, DOT},
	'9': []Signal{DASH, DASH, DASH, DASH, DOT},
	'0': []Signal{DASH, DASH, DASH, DASH, DASH},
}

// Morse sends a message in morse code.
func Morse(c registry.RuntimeConfig) error {
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
			signal, ok := Alphabet[char]
			if !ok {
				log.Printf("Warning: character '%s' has no morse code symbol.\n", string(char))
				continue
			}

			// Make the flashes.
			for _, flash := range signal {
				AllLights(c, func(light hue.Light) {
					light.SetState(brightState)
				})

				if flash == DOT {
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
}

func init() {
	registry.Register(
		"morse",
		"Spell out a message in morse code.\n"+
			"Use -delay to control the duration of a dot.\n"+
			"The dash duration is 3X the dot duration.",
		Morse,
	)
}
