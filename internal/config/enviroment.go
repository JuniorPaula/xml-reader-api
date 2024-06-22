package config

import (
	"log"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *Config

type Config struct {
	Port      string `mapstructure:"PORT"`
	JWTSecret string `mapstructure:"JWT_SECRET"`
	JWTExp    int    `mapstructure:"JWT_EXP"`
	TokenAuth *jwtauth.JWTAuth
}

func LoadEnv(path string) (*Config, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, nil
}
