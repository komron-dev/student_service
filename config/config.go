package config

import (
	"os"

	"github.com/spf13/cast"
)

type Config struct {
	Environment string

	PostgresHost     string
	PostgresPort     int
	PostgresDatabase string
	PostgresUser     string
	PostgresPassword string

	TaskServiceHost string
	TaskServicePort int

	LogLevel string
	RPCPort  string
}

func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "students"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "komron"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "pass_kom"))

	c.TaskServiceHost = cast.ToString(getOrReturnDefault("STUDENT_TASK_SERVICE_HOST", "localhost"))
	c.TaskServicePort = cast.ToInt(getOrReturnDefault("STUDENT_TASK_SERVICE_PORT", 9005))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":9005"))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
