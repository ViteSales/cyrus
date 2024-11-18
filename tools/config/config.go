package config

import (
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("cyrus")
	viper.AddConfigPath("./cyrus")
	viper.SetConfigType("toml")
	viper.ReadInConfig()
}
