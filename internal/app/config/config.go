package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)


type Config struct {
	Env string `yaml:"env" env-required:"true"`
	HttpAddress string `yaml:"http_address" env-required:"true"`
	StoragePlaylist string `yaml:"storage_playlist" enf-required:"storage_playlist"`
}



func Read()(*Config, error){

	configPath := os.Getenv("Config")

	if configPath == "" {
		return nil, fmt.Errorf("config path is empty")
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("failed read config: %v", err)
	}
	return &cfg, nil

}