package config

import (
    "log"

    "github.com/spf13/viper"
)

var Config *viper.Viper

func LoadConfig() {
    Config = viper.New()
    Config.SetConfigName("config")
    Config.AddConfigPath(".")
    Config.SetConfigType("yaml")

    if err := Config.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file: %v", err)
    }
}
