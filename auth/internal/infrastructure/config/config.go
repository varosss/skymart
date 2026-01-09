package config

import (
	"fmt"
	"os"
	"time"
)

type Database struct {
	DSN string
}

type Redis struct {
	DSN string
}

type JWT struct {
	Issuer     string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

type HttpServer struct {
	HostURL string
	Port    string
}

type GrpcServer struct {
	Port string
}

type UserService struct {
	GRPCAddr string
}

type Security struct {
	PublicKeyPath  string
	PrivateKeyPath string
}

type Config struct {
	HttpServer  HttpServer
	GrpcServer  GrpcServer
	Database    Database
	Security    Security
	UserService UserService
	JWT         JWT
	Redis       Redis
}

func Load() (*Config, error) {
	accessTTL, err := time.ParseDuration(os.Getenv("JWT_ACCESS_TTL"))
	if err != nil {
		return nil, err
	}

	refreshTTL, err := time.ParseDuration(os.Getenv("JWT_REFRESH_TTL"))
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Security: Security{
			PublicKeyPath:  os.Getenv("SECURITY_PUBLIC_KEY_PATH"),
			PrivateKeyPath: os.Getenv("SECURITY_PRIVATE_KEY_PATH"),
		},
		HttpServer: HttpServer{
			Port:    getEnv("HTTP_SERVER_PORT", "80"),
			HostURL: os.Getenv("HOST_URL"),
		},
		GrpcServer: GrpcServer{
			Port: getEnv("GRPC_SERVER_PORT", "9000"),
		},
		Database: Database{
			DSN: os.Getenv("POSTGRES_URL"),
		},
		UserService: UserService{
			GRPCAddr: os.Getenv("GRPC_USER_SERVICE_ADDR"),
		},
		JWT: JWT{
			Issuer:     os.Getenv("JWT_ISSUER"),
			AccessTTL:  accessTTL,
			RefreshTTL: refreshTTL,
		},
		Redis: Redis{
			DSN: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
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
