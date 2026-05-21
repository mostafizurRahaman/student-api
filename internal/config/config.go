package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Address string `yaml:"address" env:"ADDRESS" env-required:"true" default:"localhost:8080"`
	timeout uint16 `yaml:"timeout" env:"TIMEOUT" env-required:"true" env-default:"5000"`
}

// ! We can add anotation into struct
// ! This are called struct tags : `yaml:""`
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true" env-default:"storage/storage.db"`
	HttpServer  `yaml:"http_server"`
}

func MustLoad() {

	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" { 
		flags := flag.String("config", "", "path to the configuration file")

		flag.Parse()

		configPath = *flags

		if configPath == ""{ 
			log.Fatal("Cofig path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) { 
		log.Fatalf("Cofig path doesn't exist: %s", configPath)
	}	

	var cfg Config

 		err:=	cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("can not read config file : %s", err.Error())
	}

	return &cfg

}