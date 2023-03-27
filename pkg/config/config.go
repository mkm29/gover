package config

import (
	"log"

	"github.com/spf13/viper"
)

// LoadConfig reads configuration from file or environment variables.
func Loadvariables(args ...string) (variables *Variables, err error) {
	if len(args) > 0 {
		viper.AddConfigPath(args[0])
		if len(args) > 1 {
			viper.SetConfigName(args[1])
			viper.SetConfigType("env")
		}
		err = viper.ReadInConfig()
	}
	viper.AutomaticEnv()
	if err != nil {
		log.Fatalf("Can't load config file: %s\n", err)
	}
	if err := viper.Unmarshal(&variables); err != nil {
		log.Fatalf("Can't unmarshal config: %s\n", err)
	}
	return
}
