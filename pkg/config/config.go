package config

import (
	"log"

	"github.com/spf13/viper"
)

// LoadConfig reads configuration from file or environment variables.
func Loadvariables(args ...interface{}) (variables *Variables, err error) {
	if len(args) > 0 {
		viper.AddConfigPath(args[0].(string))
		if len(args) > 1 {
			viper.SetConfigName(args[1].(string))
			viper.SetConfigType("env")
		}
		err = viper.ReadInConfig()
	}
	viper.AutomaticEnv()
	if err != nil {
		log.Fatalf("Can't load config file: %s\n", err)
	}
	err = viper.Unmarshal(&variables)
	return
}
