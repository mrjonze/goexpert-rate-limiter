package config

import (
	"github.com/spf13/viper"
)

type EnvConfig struct {
	TokenName string `mapstructure:"TOKEN_NAME"`
}

func LoadConfig() (*EnvConfig, error) {
	var config *EnvConfig
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return config, nil
}
