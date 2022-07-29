package config

import "github.com/spf13/viper"

func SetViperDefaults() {
	viper.SetDefault(KeyAppEnv, "production")
	viper.SetDefault(KeyAppContents, "./contents")
	viper.SetDefault(KeyAppViews, "./web/views")
	viper.SetDefault(KeyHttpProtocol, "http")
	viper.SetDefault(KeyHttpHost, "localhost")
	viper.SetDefault(KeyHttpPort, 1819)
}
