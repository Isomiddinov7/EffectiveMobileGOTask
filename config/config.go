package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	Error = "error >>> "
	Info  = "info >>> "
	Log   = "log >>> "
)

type Config struct {
	PostgresHost          string
	PostgresUser          string
	PostgresDatabase      string
	PostgresPassword      string
	PostgresPort          int
	PostgresMaxConnection int32

	DefaultOffset string
	DefaultLimit  string

	ServiceHost     string
	ServiceHTTPPort string
}

func Load() Config {

	if err := godotenv.Load(".env"); err != nil {
		log.Println("not found env")
	}

	var cfg Config

	cfg.ServiceHost = cast.ToString(getValueOrDefault("SERVICE_HOST", "localhost"))
	cfg.ServiceHTTPPort = cast.ToString(getValueOrDefault("SERVICE_HTTP_PORT", ":8080"))

	cfg.PostgresHost = cast.ToString(getValueOrDefault("POSTGRES_HOST", "localhost"))
	cfg.PostgresUser = cast.ToString(getValueOrDefault("POSTGRES_USER", "bahodir"))
	cfg.PostgresDatabase = cast.ToString(getValueOrDefault("POSTGRES_DATABASE", "songs"))
	cfg.PostgresPassword = cast.ToString(getValueOrDefault("POSTGRES_PASSWORD", "1100"))
	cfg.PostgresPort = cast.ToInt(getValueOrDefault("POSTGRES_PORT", 5432))
	cfg.PostgresMaxConnection = cast.ToInt32(getValueOrDefault("POSTGRES_MAX_CONNECTION", 30))

	return cfg
}

func getValueOrDefault(key string, defaultValue interface{}) interface{} {

	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
