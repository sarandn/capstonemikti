package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

var Config *viper.Viper

func LoadConfig() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err.Error())
	}
	// Config = viper.New()
	// Config.SetConfigName("config")
	// Config.AddConfigPath(".")
	// Config.SetConfigType("yaml")

	// if err := Config.ReadInConfig(); err != nil {
	// 	log.Fatalf("Error reading config file: %v", err)
	// }
}
