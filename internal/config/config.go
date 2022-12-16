package config

import "github.com/spf13/viper"

type Config struct {
	Env struct {
		DatabaseUser     string `mapstructure:"DATABASE_USER"`
		DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
		DatabaseHost     string `mapstructure:"DATABASE_HOST"`
		DatabasePort     int    `mapstructure:"DATABASE_PORT"`
		DatabaseName     string `mapstructure:"DATABASE_NAME"`
	}
}

var config Config

func Load(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&config.Env); err != nil {
		panic(err)
	}
}

func Get() Config {
	return config
}
