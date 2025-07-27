package config

import (
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct {
	Host     string `mapstructure:"DB_HOST"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
}

var once sync.Once
var config *Config

var configError error

// Init populates Config struct with values from config file
// located at filepath and environment variables.
func Init() (*Config, error) {
	once.Do(func() {
		if err := parseConfigFile(); err != nil {
			config = nil
			configError = err
			return
		}

		var cfg Config
		if err := unmarshal(&cfg); err != nil {
			config = nil
			configError = err
			return
		}

		config = &cfg
		configError = nil
	})

	return config, configError
}

func parseConfigFile() error {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, err := os.Stat("/.dockerenv"); err != nil {
			/*
				if running in a Dockerfile container, a missing configuration might prevent
				correct behaviour (depending on core functionality and lack of environment variable usage)
			*/
			return err
		}

		for _, s := range os.Environ() {
			a := strings.Split(s, "=")
			viper.BindEnv(a[0])
		}
	}

	return viper.MergeInConfig()
}

func unmarshal(cfg *Config) error {
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return err
	}

	return viper.Unmarshal(&cfg.DB)
}
