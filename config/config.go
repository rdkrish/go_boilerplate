// Config package contains all the configurations for your project
package config

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string, script bool) {
	var err error
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(strings.ToLower(env))
	if env == "test" || script {
		v.AddConfigPath("../config/")
	} else {
		v.AddConfigPath("config/")
	}
	err = v.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
	config = v
}

// GetConfig function to expose the config object
func GetConfig() *viper.Viper {
	return config
}
