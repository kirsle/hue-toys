// Package auth handles connecting and managing credentials for the Hue bridge.
package auth

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	hue "github.com/collinux/gohue"
	"github.com/kirsle/configdir"
)

// An app name for the user registration.
const DEVICE_TYPE = "hue-toys"

// The file we store the credentials in.
var CREDENTIALS string

func init() {
	// Make sure the configuration path exists.
	configdir.MakePath(configdir.LocalConfig())
	CREDENTIALS = configdir.LocalConfig("kirsle", "hue-toys", "credentials.json")
}

// Setup connects to the Hue bridge and registers the username if needed.
func Setup() (*hue.Bridge, error) {
	var (
		bridge   *hue.Bridge // The hue bridge when we get access to it.
		username string
		err      error
	)

	bridges, err := hue.FindBridges()
	if err != nil {
		return nil, err
	}
	bridge = &bridges[0]

	// Have any stored config?
	username, err = GetConfig()
	if err != nil {
		log.Printf("No stored configuration (%s)\n", err)
		log.Println("Press the 'Connect' button on top of your Hue bridge, then")
		log.Println("Press Enter to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')

		// Register the user.
		username, err = bridge.CreateUser(DEVICE_TYPE)
		fmt.Printf("%v, %v\n", bridge, err)
		if err != nil || username == "" {
			return nil, fmt.Errorf("Bridge authentication failed (no username); error: %v", err)
		}

		// Got the username, let's save it.
		SaveConfig(username)
	}

	// Log in.
	bridge.Login(username)

	log.Printf("Connected to Hue bridge at %s (username %s)\n", bridge.IPAddress, username)
	return bridge, nil
}

// Type config describes the credentials.json format.
type config struct {
	Username string `json:"username"`
}

// GetConfig retrieves the stored credentials.
func GetConfig() (string, error) {
	// See if the config exists.
	if _, err := os.Stat(CREDENTIALS); os.IsNotExist(err) {
		return "", fmt.Errorf("%s does not exist", CREDENTIALS)
	}

	// Open the config file.
	fh, err := os.Open(CREDENTIALS)
	if err != nil {
		return "", err
	}

	// Decode the JSON.
	cfg := config{}
	decoder := json.NewDecoder(fh)
	err = decoder.Decode(&cfg)
	if err != nil {
		return "", err
	}

	return cfg.Username, nil
}

// SaveConfig saves the username to disk.
func SaveConfig(username string) error {
	fh, err := os.Create(CREDENTIALS)
	if err != nil {
		return err
	}

	defer fh.Close()

	encoder := json.NewEncoder(fh)
	encoder.Encode(&config{
		Username: username,
	})

	return nil
}
