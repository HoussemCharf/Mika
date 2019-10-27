package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configuration file attributes.
type Configuration struct {
	Token  string
	Prefix string
	Debug  bool
}

// Grab is a function that return template of config.json.
func Grab() (*Configuration, error) {
	var config = new(Configuration)
	// grab config for discord bot here
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("problem grabbing config: %v", err)
	}

	// filling the current config
	config.Token = viper.GetString("bot.Token")
	config.Prefix = viper.GetString("bot.Prefix")
	config.Debug = viper.GetBool("general.Debug")
	return config, nil
}
