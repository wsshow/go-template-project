package config

import (
	"github.com/spf13/viper"
)

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.SetDefault("is_https", false)
	viper.SetDefault("port", 9095)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			if err := viper.WriteConfigAs("config.yaml"); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func IsHttps() bool {
	return viper.GetBool("is_https")
}

func DefaultPort() int {
	return viper.GetInt("port")
}
