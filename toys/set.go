package toys

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	hue "github.com/collinux/gohue"
	"github.com/kirsle/hue-toys/registry"
)

// colorMap maps human readable color names into Hue colors.
var colorMap = map[string]*[2]float32{
	"red":    hue.RED,
	"yellow": hue.YELLOW,
	"orange": hue.ORANGE,
	"green":  hue.GREEN,
	"cyan":   hue.CYAN,
	"blue":   hue.BLUE,
	"purple": hue.PURPLE,
	"pink":   hue.PINK,
	"white":  hue.WHITE,
}
var sortedColors []string

func init() {
	// Sort the color keys for consistent output.
	sortedColors = []string{}
	for k, _ := range colorMap {
		sortedColors = append(sortedColors, k)
	}
	sort.Strings(sortedColors)

	// Prints the valid colors.
	printColorList := func() {
		fmt.Println("The valid color options are:")
		for _, k := range sortedColors {
			fmt.Printf("* %s\n", k)
		}
	}

	registry.Register(
		"set",
		"Set the lights according to a series of colors (e.g. "+
			"hue-toys -t set red blue green yellow)",
		func(c registry.RuntimeConfig) error {
			args := flag.Args()
			if len(args) == 0 {
				fmt.Print(
					"Usage: hue-toys -t set <colors...>\n" +
						"Example: hue-toys -t set red green blue yellow\n\n" +
						"This toy sets each light in sequence to the set of colors you give it.\n" +
						"If there are more lights than colors, they'll wrap around and reuse\n" +
						"the same list given.\n\n",
				)
				printColorList()
				os.Exit(1)
			}

			// Color and parse the colors.
			colors := []*[2]float32{}
			for _, arg := range args {
				arg = strings.ToLower(arg)
				if _, ok := colorMap[arg]; !ok {
					fmt.Printf("Error: %s is not an acceptable color.", arg)
					printColorList()
					os.Exit(1)
				}
				colors = append(colors, colorMap[arg])
			}

			// Index counter to step through the color list in order.
			var index int

			// Set the lights to each color in series.
			AllLights(c, func(light hue.Light) {
				// Get the color for this light.
				nextColor := colors[index]

				// Increment and wrap the color index.
				index++
				if index == len(colors) {
					index = 0
				}

				// Set this light's color.
				light.SetColor(nextColor)
			})

			return nil
		},
	)
}
