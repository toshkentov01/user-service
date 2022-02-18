package config

import (
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/spf13/cast"
)

// Config struct
type Config struct {
	Environment      string
	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string
	LogLevel         string
	RPCPort          string
}

// load for loading a config
func load() *Config {
	return &Config{
		Environment:      cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop")),
		PostgresHost:     cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost")),
		PostgresPort:     cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432)),
		PostgresDatabase: cast.ToString(getOrReturnDefault("POSTGRES_DB", "")),
		PostgresUser:     cast.ToString(getOrReturnDefault("POSTGRES_USER", "")),
		PostgresPassword: cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "")),
		LogLevel:         cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug")),
		RPCPort:          cast.ToString(getOrReturnDefault("RPC_PORT", ":9000")),
	}
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}

var (
	instance *Config
	once     sync.Once
)

//Get ...
func Get() *Config {
	once.Do(func() {
		instance = load()
	})

	return instance
}
