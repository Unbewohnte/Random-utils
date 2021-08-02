package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"twitch-hooks/discordhooks"
	"twitch-hooks/twitchhooks"
	"twitch-hooks/vkhooks"
)

const configFilename string = "config.cfg"

type keys struct {
	Twitch  twitchhooks.Keys
	Discord discordhooks.Webhook
	VK      vkhooks.ApiKey
}

type messages struct {
	DiscordMessage discordhooks.Message
	VKmessage      vkhooks.Message
}

type Config struct {
	TwitchName string
	Keys       keys
	ForceSend  bool `json:"force-send"`
	Messages   messages
}

// Checks if config file exists in the same directory
func ConfigExists() bool {
	_, err := os.Stat(configFilename)
	if err != nil {
		return false
	}
	return true
}

// Creates a new config file in current directory.
func CreateConfig() error {
	// create a config file in the same directory
	configF, err := os.Create(configFilename)
	if err != nil {
		return fmt.Errorf("could not create a config file: %s", err)
	}

	// write default config fields
	defaults, err := json.MarshalIndent(&Config{}, "", "    ")
	if err != nil {
		return fmt.Errorf("could not marshal default config fields: %s", err)
	}
	_, err = configF.Write(defaults)
	if err != nil {
		return fmt.Errorf("could not write defaults to config: %s", err)
	}

	return nil
}

// Opens and reads config file, returns `Config` struct.
// If ReadConfig cannot unmarshal config file - it creates a new one with
// all default fields
func ReadConfig() (*Config, error) {
	// get config`s contents
	configContents, err := os.ReadFile(configFilename)
	if err != nil {
		return nil, fmt.Errorf("could not read config: %s", err)
	}

	var config Config
	err = json.Unmarshal(configContents, &config)
	if err != nil {
		_ = CreateConfig()
		return nil, fmt.Errorf("could not unmarshal config: %s\nCreatead a new one", err)
	}

	// remove uneccessary spaces
	config.Keys.Discord.WebhookUrl = strings.TrimSpace(config.Keys.Discord.WebhookUrl)
	config.Keys.Twitch.ClientID = strings.TrimSpace(config.Keys.Twitch.ClientID)
	config.Keys.Twitch.ClientSecret = strings.TrimSpace(config.Keys.Twitch.ClientSecret)
	config.Keys.VK.Key = strings.TrimSpace(config.Keys.VK.Key)

	// validate inputs
	if config.Keys.Discord.WebhookUrl == "" &&
		config.Keys.Twitch.ClientID == "" &&
		config.Keys.Twitch.ClientSecret == "" &&
		config.Keys.VK.Key == "" {

		return nil, fmt.Errorf("does not use any keys")
	}
	if len(config.TwitchName) < 2 {
		return nil, fmt.Errorf("twitch name is too short")
	}

	return &config, nil
}
