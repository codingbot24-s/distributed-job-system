package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// load env here

type Config struct {
	ApiServerPort string
	Redis         string
	Postgres      string
	LogLevel      string
}

func (c *Config) LoadConfig() error {
	viper.AddConfigPath("../../.env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("error reading config %w", err)
	}

	return nil
}
