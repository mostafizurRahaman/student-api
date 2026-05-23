package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// ? Define another struct for http server:

type HttpServer struct {
	Address string `yaml:"address" env:"ADDRESS"  env-required:"true" env-default:"localhost:8080"`
	Timeout string `yaml:"timeout" env:"TIMEOUT" env-required:"true" env-default:"timeout"`
}

// ? Define a struct:
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HttpServer  `yaml:"http_server"`
}

func MustLoad() *Config {

	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {

		flags := flag.String("config", "", "Please add config file path")

		flag.Parse()

		configPath = *flags

		if configPath == "" {
			log.Fatal("Failed to load configuration file.")
		}

	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Failed to load the env file: %s", err.Error())
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Failed to parse env: %s", err.Error())
	}

	return &cfg

}
