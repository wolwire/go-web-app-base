package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfig() {
	fmt.Println("Loading config ..........")
	viper.SetConfigName("config.local")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	fmt.Println("Config Loaded ...........")
}
