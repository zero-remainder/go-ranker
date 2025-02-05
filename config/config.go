package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	AppName string `mapstructure:"APP_NAME"`
	Port    int    `mapstructure:"PORT"`
	Debug   bool   `mapstructure:"DEBUG"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
	}
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}
