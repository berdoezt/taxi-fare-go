package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	FareRules map[string]FareRule `mapstructure:"fare_rules"`
	App       AppConfig           `mapstructure:"app"`
}

type AppConfig struct {
	Debug    bool   `mapstructure:"debug"`
	LogLevel string `mapstructure:"log_level"`
}

type FareRule struct {
	Price             float64 `mapstructure:"price"`
	Distance          float64 `mapstructure:"distance"`
	DistanceThreshold float64 `mapstructure:"distance_threshold"`
}

func LoadConfig() (Config, error) {
	cfg := Config{}

	viper.AddConfigPath(".")
	viper.SetConfigName("application")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		return cfg, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
