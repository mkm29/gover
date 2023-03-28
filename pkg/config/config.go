package config

import (
	"log"
	"reflect"

	"github.com/spf13/viper"
)

var cfg *Config

func init() {
	cfg = &Config{
		requiredVars: map[string]string{
			"CI_DEFAULT_BRANCH":                   "DefaultBranch",
			"CI_MERGE_REQUEST_TARGET_BRANCH_NAME": "MergeRequestTargetBranchName",
			"CI_PIPELINE_IID":                     "PipelineIid",
		},
		Debug: false,
	}
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(args ...string) (config *Config, err error) {
	v := viper.New()
	if len(args) == 2 {
		path, fname := args[0], args[1]
		v.AddConfigPath(path)
		v.SetConfigName(fname)
		v.SetConfigType("env")
		err = v.ReadInConfig()
		if err != nil {
			return nil, err
		}
	}
	v.AutomaticEnv()
	if len(args) == 0 {
		for ev, _ := range cfg.requiredVars {
			v.BindEnv(ev)
		}
	}
	// set defaults
	setDefaults(v)
	var variables Variables

	if err := v.Unmarshal(&variables); err != nil {
		return nil, err
	}
	// update cfg with variables
	cfg.Variables = &variables
	if cfg.Debug {
		log.Printf("Unmarshaled: %+v\n", cfg.Variables)
	}
	return cfg, nil
}

func setDefaults(v *viper.Viper) {
	v.SetDefault("CI_BUILDS_DIR", "/builds")
}

func (c *Config) CheckVariables() (bool, []string) {
	// check if necessary variables are set
	var missing []string
	metaValue := reflect.ValueOf(c.Variables).Elem()
	for _, name := range cfg.requiredVars {
		field := metaValue.FieldByName(name)
		if field == (reflect.Value{}) {
			missing = append(missing, name)
		}
	}
	if len(missing) > 0 {
		return false, missing
	}
	return true, nil
}
