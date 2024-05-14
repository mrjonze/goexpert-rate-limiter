package config

import (
	"github.com/spf13/viper"
)

type EnvConfig struct {
	TokenName         string `mapstructure:"TOKEN_NAME"`
	RequestLimitToken int    `mapstructure:"REQUEST_LIMIT_TOKEN"`
	RequestLimitIp    int    `mapstructure:"REQUEST_LIMIT_IP"`
	BlockTimeToken    int    `mapstructure:"BLOCK_TIME_TOKEN"`
	BlockTimeIp       int    `mapstructure:"BLOCK_TIME_IP"`
	DatabaseUrl       string `mapstructure:"DATABASE_URL"`
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
