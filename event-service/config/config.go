package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	RedisAddress  string `mapstructure:"REDIS_ADDRESS"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("../.env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file, %s", err)
		return
	}

	err = viper.Unmarshal(&config)
	return
}
