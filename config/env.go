package config

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	ServerPort string `mapstructure:"SERVER_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPass     string `mapstructure:"DB_PASS"`
	DBName     string `mapstructure:"DB_NAME"`
	JwtKey     string `mapstructure:"JWT_KEY"`
}

func NewEnv() *Env {
	env := Env{}
	viper.AutomaticEnv()

	err := viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}
