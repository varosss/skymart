package config

import (
	"os"
	"strconv"
)

type Database struct {
	DSN string
}

type GrpcServer struct {
	Port string
}

type Security struct {
	HashCost int
}

type Config struct {
	GrpcServer GrpcServer
	Database   Database
	Security   Security
}

func Load() (*Config, error) {
	cost, err := strconv.Atoi(os.Getenv("SECURITY_HASH_COST"))
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		GrpcServer: GrpcServer{
			Port: getEnv("GRPC_SERVER_PORT", "9000"),
		},
		Database: Database{
			DSN: os.Getenv("POSTGRES_URL"),
		},
		Security: Security{
			HashCost: cost,
		},
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
