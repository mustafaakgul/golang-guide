package config

import (
	"os"
	"strconv"
)

type Global struct {
	DBHost     string
	DBPort     int
	DBName     string
	DBUser     string
	DBPassword string
}

func SetConfig(prefix string) Global {
	port, _ := strconv.Atoi(os.Getenv(prefix + "PGPORT"))
	if port == 0 {
		port = 5432
	}

	return Global{
		DBHost:     os.Getenv(prefix + "PGHOST"),
		DBPort:     port,
		DBName:     os.Getenv(prefix + "PGDATABASE"),
		DBUser:     os.Getenv(prefix + "PGUSER"),
		DBPassword: os.Getenv(prefix + "PGUSER"),
	}
}
