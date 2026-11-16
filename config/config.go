package config

import "github.com/spf13/viper"

type Config struct {
	Token      string `mapstructure:"token"`
	WebhookUrl string `mapstructure:"webhook"`
}

func Init() (*Config, error) {
	var config Config
	if err := ParseEnv(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func ParseEnv(cfg *Config) error {
	viper.AutomaticEnv()
	if err := viper.BindEnv("token", "TOKEN", "webhook", "WEBHOOK"); err != nil {
		return err
	}

	cfg.Token = viper.GetString("token")
	cfg.WebhookUrl = viper.GetString("webhook")
	return nil

}
