package config

import (
	"os"
	"strconv"
	"sync"

	"github.com/goakshit/gandalf/internal/types"
)

// Singleton
var conf *types.Config
var once sync.Once

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	env := getEnv(key, strconv.Itoa(defaultValue))
	result, err := strconv.Atoi(env)
	if err != nil {
		panic(err)
	}
	return result
}

func getPostgresConf() types.PostgresConfig {
	return types.PostgresConfig{
		DatabaseName: getEnv("POSTGRES_DB_NAME", "billingdb"),
		User:         getEnv("POSTGRES_USER", "dev"),
		Passwd:       getEnv("POSTGRES_PASSWD", "passwd"),
		Host:         getEnv("POSTGRES_HOST", "localhost"),
		Port:         getEnvAsInt("POSTGRES_PORT", 5432),
		SSLMode:      getEnv("POSTGRES_SSLMODE", "disable"),
		Timezone:     getEnv("POSTGRES_TZ", "UTC"),
	}
}

// New - Initialize Configuration
func New() *types.Config {
	once.Do(func() {
		conf = &types.Config{
			Database: getPostgresConf(),
		}
	})

	return conf
}
