package config

import (
	"fmt"
	"os"
	"strconv"
)

var global *config

type config struct {
	DatabaseName    string
	MongoDbHost     string
	MongoDbUsername string
	MongoDbPassword string
	MongoDbPort     uint64
}

func Load() error {
	port, err := strconv.ParseUint(os.Getenv("MONGODB_PORT"), 10, 0)
	if err != nil {
		return fmt.Errorf("error parsing port %s as int: %v", os.Getenv("MONGODB_PORT"), err)
	}

	global = &config{
		DatabaseName:    os.Getenv("DATABASE_NAME"),
		MongoDbHost:     os.Getenv("MONGODB_HOST"),
		MongoDbUsername: os.Getenv("MONGODB_USERNAME"),
		MongoDbPassword: os.Getenv("MONGODB_PASSWORD"),
		MongoDbPort:     port,
	}

	return nil
}

func Global() *config {
	if global == nil {
		global = &config{}
	}
	return global
}
