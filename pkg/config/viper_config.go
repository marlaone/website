package config

import "github.com/spf13/viper"

func SetViperDefaults() {
	viper.SetDefault("app.env", "production")
	viper.SetDefault("app.contents", "./contents")
	viper.SetDefault("app.views", "./web/views")
	viper.SetDefault("http.protocol", "http")
	viper.SetDefault("http.host", "localhost")
	viper.SetDefault("http.port", 1819)
}
