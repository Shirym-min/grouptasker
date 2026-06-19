package config

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type Config struct {
	Tasks map[string][]string `mapstructure:"tasks"`
}

func Load() (*Config, error) {
	home, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	configDir := filepath.Join(home, "gpx")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, err
	}
	configPath := filepath.Join(home, "gpx/gpx.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		err = os.WriteFile(
			configPath,
			[]byte("tasks: {}\n"),
			0644,
		)
		if err != nil {
			return nil, err
		}
	}
	viper.SetConfigFile(configPath)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
