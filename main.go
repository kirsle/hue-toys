package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/kirsle/hue-toys/auth"
	"github.com/kirsle/hue-toys/registry"
	_ "github.com/kirsle/hue-toys/toys"
)

const (
	// App username.
	DEVICE_TYPE = "hue-toys"
)

var (
	// Command line flags.
	list       bool
	toy        string
	brightness int
	delay      int
)

func init() {
	listHelp := "List available Hue toys"
	flag.BoolVar(&list, "list", false, listHelp)
	flag.BoolVar(&list, "ls", false, listHelp+" (alias)")

	toyHelp := "Specify a Hue toy to run"
	flag.StringVar(&toy, "toy", "", toyHelp)
	flag.StringVar(&toy, "t", "", toyHelp+" (alias)")

	brightHelp := "Specify a global light brightness for the toys to use (1-100)"
	flag.IntVar(&brightness, "brightness", 100, brightHelp)
	flag.IntVar(&brightness, "b", 100, brightHelp+" (alias)")

	delayHelp := "Specify a global delay (in milliseconds) for light changes"
	flag.IntVar(&delay, "delay", 1000, delayHelp)
	flag.IntVar(&delay, "d", 1000, delayHelp+" (alias)")
}

func main() {
	flag.Parse()

	// Are they listing available toys with -list?
	if list {
		registry.PrintList()
		os.Exit(0)
	}

	// Did they specify a toy to run with -toy?
	if toy == "" {
		fmt.Println("Usage: hue-toys -toy <name>")
		fmt.Println("Use `hue-toys -help` for more usage information.")
		os.Exit(1)
	}

	// Authenticate with the bridge.
	bridge, err := auth.Setup()
	handle(err)

	err = registry.Run(registry.RuntimeConfig{
		Name:       toy,
		Bridge:     bridge,
		Brightness: brightness,
		Delay:      time.Duration(delay),
	})
	handle(err)
}

// Handle errors by panicking.
func handle(err error) {
	if err != nil {
		panic(err)
	}
}
