package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var Viper *config

type config struct {
	*viper.Viper
}

func init() {
	v := viper.New()
	v.AddConfigPath(".")
	v.AddConfigPath("../")
	v.SetDefault("env", "dev")
	v.BindEnv("env")
	v.SetConfigName("config-" + v.GetString("env"))
	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	Viper = &config{v}
}

func GetEnvironmentSpecificConfig(env string) *viper.Viper {
	v := viper.New()
	v.AddConfigPath(".")
	v.AddConfigPath("../")
	v.SetConfigName("config-" + env)
	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return v
}
