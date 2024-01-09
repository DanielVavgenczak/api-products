package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigEnv struct {
	JWTSecret string `mapstructure:""`
}

func LoadEnv() bool {
	_,err := load()
	return err == nil
}

func load() (ConfigEnv, error) {
	err := godotenv.Load()
	if err != nil {
		return ConfigEnv{}, err
	}

	teste := os.Getenv("TOKEN_CONFIG")
	
	config := ConfigEnv{
		JWTSecret: teste,
	}
	
	return config, nil
}