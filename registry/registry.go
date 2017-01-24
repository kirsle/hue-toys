// Package registry provides functionality for the Hue toys to self-register.
package registry

import (
	"fmt"
	"log"
	"sync"
	"time"

	hue "github.com/collinux/gohue"
)

// Type Handler is the function prototype for a toy's entry point.
type Handler func(RuntimeConfig) error

// Type toy is a private internal map of toy data.
type toy struct {
	description string
	handler     Handler
}

var (
	lock sync.Mutex
	toys map[string]toy
)

func init() {
	toys = map[string]toy{}
}

// Register allows a toy to register its own function.
func Register(name, description string, fn Handler) {
	lock.Lock()
	defer lock.Unlock()
	toys[name] = toy{
		description: description,
		handler:     fn,
	}
}

// Type RuntimeConfig communicates options from the program to the toy.
type RuntimeConfig struct {
	Name       string // The name of the toy to run
	Bridge     *hue.Bridge
	Brightness int
	Delay      time.Duration
}

// Run executes a toy and returns its error, if any.
func Run(c RuntimeConfig) error {
	lock.Lock()
	defer lock.Unlock()

	// Make sure it exists.
	if _, ok := toys[c.Name]; !ok {
		return fmt.Errorf("%s: toy not found", c.Name)
	}

	log.Printf("Starting toy: %s\n", c.Name)
	return toys[c.Name].handler(c)
}

// PrintList prints all available toys to standard output.
func PrintList() {
	lock.Lock()
	defer lock.Unlock()

	fmt.Printf("Available Toys:\n\n")

	for name, info := range toys {
		fmt.Printf("* %s\n  %s\n", name, info.description)
	}
}
