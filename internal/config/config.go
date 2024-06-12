package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	APIPort     int    `envconfig:"API_PORT" default:"7777"`
	PathRoot    string `envconfig:"PATH_ROOT"`
	MongoHost   string `envconfig:"MONGO_HOST" default:"localhost:27017"`
	MongoDBName string `envconfig:"MONGO_DB"`
}

func Read() Config {
	conf := Config{} //nolint:exhaustruct

	if err := godotenv.Load(".env"); err != nil && !errors.Is(err, os.ErrNotExist) {
		panic(err)
	}

	if err := envconfig.Process("", &conf); err != nil {
		panic(err)
	}

	return conf
}

func MongoURI(conf Config) string {
	uri := fmt.Sprintf("mongodb://%s", conf.MongoHost)

	return uri
}
