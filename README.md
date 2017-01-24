# Philips Hue Toys

This is a simple Go program that uses the [Philips Hue](https://www.meethue.com/)
API to control the lights.

Actually, it's a collection of programs I call "toys" that do various specific
tasks. Each toy is a self-contained Go function that registers itself
automatically when the program is compiled/run.

## Installation

```bash
go get github.com/kirsle/hue-toys
```

## Usage

```
hue-toys [options] -toy NAME
```

Options include:

* `-list`, `-ls`

  List all available toys.

* `-toy <NAME>`, `-t <NAME>`

  Choose which toy to run.

* `-delay <MILLISECONDS>`, `-d <MILLISECONDS>`

  Specify how long to delay between light transitions, in milliseconds, if the
  toy supports it. The default is `1000`, or 1 second. All toys accept this
  parameter unless noted otherwise.

* `-brightness <PERCENT>`, `-b <PERCENT>`

  Specify which brightness level to use, as a percentage between 1 and 100.
  The default is 100. All the toys should accept this parameter.

## Toys

* `on`: turns on all the lights.

* `off`: turns off all the lights.

* `random`: turns on all the lights set to random colors for each.

* `disco`: chooses a random color for each light, and then puts it into a
  ColorLoop and exits. The ColorLoop is a feature of the Hue lights and doesn't
  accept a `-delay` parameter.

* `rainbow`: begins a rainbow animation with all lights in sync.

* `morse`: spells out a message in Morse code. This accepts additional command
  line arguments to specify the message to encode.

  Example:

  ```
  hue-toys -t morse hello world
  ```

## Adding Toys

To add a new toy, create a new Go source file in the `toys` directory with
this boilerplate code:

```go
package toys

import (
	"github.com/kirsle/hue-toys/registry"
)

// This function is the entry point to your toy and accepts a RuntimeConfig
// object, which communicates command line options such as Brightness and Delay
// and has a reference to the Hue bridge.
func MyToy(c registry.RuntimeConfig) error {
	lights, _ := c.Bridge.GetAllLights()
	for _, light := range lights {
		light.On()
	}

	return nil
}

// Register the toy with a name and description in the init() function.
// This makes it available to be called and lets its description show up
// when the user runs `hue-toys -list`
func init() {
	registry.Register("example", "An example toy.", MyToy)
}

```

## License

```
The MIT License (MIT)

Copyright (c) 2017 Noah Petherbridge

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
