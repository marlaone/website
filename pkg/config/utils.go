package config

import "github.com/spf13/viper"

func IsDebug() bool {
	return viper.GetString(KeyAppEnv) == "debug"
}

func IsProduction() bool {
	return viper.GetString(KeyAppEnv) == "production"
}
