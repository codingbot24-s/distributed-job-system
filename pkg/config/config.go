package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ApiServerPort string
	Redis         string
	Postgres      string
	LogLevel      string
}

func LoadConfig() (Config, error) {
	viper.AutomaticEnv()

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return Config{}, fmt.Errorf("error reading config file: %w", err)
		}
	}

	c := Config{
		ApiServerPort: viper.GetString("API_SERVER_PORT"),
		Redis:         viper.GetString("REDIS"),
		Postgres:      viper.GetString("POSTGRES"),
		LogLevel:      viper.GetString("LOG_LEVEL"),
	}

	
	return c, nil
}


type WorkerConfig struct {
	WorkerId string
	// number of goroutines 
	Concurrency int
	// how frequently worker check for new job
	PollInterval int
	QueueName string
	ShutdownTimeout int
}
