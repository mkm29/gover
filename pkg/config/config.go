package config

import (
	"log"
	"reflect"

	"github.com/spf13/viper"
)

// LoadConfig reads configuration from file or environment variables.
func Loadvariables(args ...string) (variables *Variables, err error) {
	v := viper.New()
	if len(args) == 2 {
		log.Println("Reading variables from file")
		path, fname := args[0], args[1]
		v.AddConfigPath(path)
		v.SetConfigName(fname)
		v.SetConfigType("env")
		if err := v.ReadInConfig(); err != nil {
			return nil, err
		}
	}
	v.AutomaticEnv()
	if len(args) == 0 {
		log.Println("Reading variables from environment")
		// iterate over all keys in Variables
		// and set them to viper
		keys := reflect.ValueOf(Variables{})
		typeOfS := keys.Type()

		for i := 0; i < keys.NumField(); i++ {
			// log.Printf("Setting env var: %s\n", typeOfS.Field(i).Name)
			v.BindEnv(typeOfS.Field(i).Name)
		}
	}
	if err := viper.Unmarshal(&variables); err != nil {
		log.Fatalf("Can't unmarshal config: %s\n", err)
	}
	return
}
