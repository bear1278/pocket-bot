package config

import "github.com/spf13/viper"

type Config struct {
	Token string `mapstructure:"token"`
}

func Init() (*Config, error) {
	var config Config
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	config.Token = viper.GetString("token")
	return &config, nil
}
