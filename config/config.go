package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type EnvVar struct {
	PGHost     string `env:"PG_HOST" envDefault:"local"`
	PGUser     string `env:"PG_USER" envDefault:"http://localhost"`
	PGPassword string `env:"PG_PASSWORD" envDefault:"3000"`
	PGPort     string `env:"PG_PORT" envDefault:"neploy"`
	PGDatabase string `env:"PG_DATABASE" envDefault:"henrrius"`
	AppURL     string `env:"APP_URL" envDefault:"Reyshell"`
	JWTSecret  string `env:"JWT_SECRET" envDefault:"localhost"`
}

var Env EnvVar

func getEnvPath() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	return filepath.Join(basepath, "..", ".env")
}

func LoadEnv() {
	if err := godotenv.Load(getEnvPath()); err != nil {
		log.Println("No .env file found")
	}

	if err := env.Parse(&Env); err != nil {
		log.Fatalf("Failed to parse env: %v", err)
	}
}
