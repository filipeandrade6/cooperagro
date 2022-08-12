package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DB DB `mapstructure:"database"`
}

type DB struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	TLS      string `mapstructure:"tls"`
}

func GetConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, fmt.Errorf("reading config file: %w", err)
	}

	var C Config
	if err := viper.Unmarshal(&C); err != nil {
		return Config{}, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return C, nil
}
