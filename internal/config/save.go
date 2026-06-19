package config

import "github.com/spf13/viper"

func Save(cfg *Config) error {
	viper.Set("tasks", cfg.Tasks)
	return viper.WriteConfig()
}
