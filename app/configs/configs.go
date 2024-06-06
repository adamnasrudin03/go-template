package configs

import (
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

var (
	lock    = &sync.Mutex{}
	configs *Configs
)

func GetInstance() *Configs {
	if configs == nil {
		lock.Lock()

		if err := godotenv.Load(); err != nil {
			log.Println("Failed to load env file")
		}

		configs = &Configs{
			App: AppConfig{
				Name:         getEnv("APP_NAME", "movie-festival"),
				Env:          getEnv("APP_ENV", "dev"),
				Port:         getEnv("APP_PORT", "8000"),
				SecretKey:    getEnv("JWT_SECRET", "MySecretKey"),
				ExpiredToken: GetExpiredToken(),
				UseRabbitMQ:  UseRabbitMQ(),
			},
			DB: DbConfig{
				Host:        getEnv("DB_HOST", "127.0.0.1"),
				Port:        getEnv("DB_PORT", "5432"),
				DbName:      getEnv("DB_NAME", "movie_festival_db"),
				Username:    getEnv("DB_USER", "postgres"),
				Password:    getEnv("DB_PASS", "postgres"),
				DbIsMigrate: getEnv("DB_IS_MIGRATE", "true") == "true",
			},
			Redis: RedisConfig{
				Host:                getEnv("REDIS_HOST", "127.0.0.1"),
				Port:                GetRedisPort(),
				Password:            getEnv("REDIS_PASSWORD", ""),
				Database:            GetRedisDatabase(),
				Master:              getEnv("REDIS_MASTER", "master"),
				PoolSize:            GetRedisPoolSize(),
				PoolTimeout:         GetRedisPoolTimeout(),
				MinIdleConn:         GetRedisMinIdleConn(),
				DefaultCacheTimeOut: GetRedisDefaultCacheTimeOut(),
			},
			RabbitMQ: RabbitMQConfig{
				Host:     getEnv("RABBIT_HOST", "127.0.0.1"),
				Port:     GetRabbitPort(),
				Username: getEnv("RABBIT_USERNAME", "GUEST"),
				Password: getEnv("RABBIT_PASSWORD", "GUEST"),
			},
		}
		lock.Unlock()
	}

	return configs
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return strings.TrimSpace(value)
	}
	return strings.TrimSpace(fallback)
}

func BackEndUrl() string {
	backEndUrl := ``
	switch os.Getenv(`ENVIRONMENT`) {
	case `dev`:
		backEndUrl = os.Getenv(`BACK_END_DEV_URL`)
	case `stg`:
		backEndUrl = os.Getenv(`BACK_END_STG_URL`)
	case `prd`:
		backEndUrl = os.Getenv(`BACK_END_PRD_URL`)
	}
	return backEndUrl
}

func ServiceName() string {
	return os.Getenv("SERVICE_NAME")
}

func GetExpiredToken() int {
	intVar, err := strconv.Atoi(getEnv("EXPIRED_TOKEN", "7"))
	if err != nil {
		return 1
	}

	return intVar
}

func GetRedisPort() int {
	intVar, err := strconv.Atoi(getEnv("REDIS_HOST", "6379"))
	if err != nil {
		return 6379
	}

	return intVar
}
func GetRedisDatabase() int {
	intVar, err := strconv.Atoi(getEnv("REDIS_DATABASE", "0"))
	if err != nil {
		return 0
	}

	return intVar
}

func GetRedisPoolSize() int {
	intVar, err := strconv.Atoi(getEnv("REDIS_POOL_SIZE", "128"))
	if err != nil {
		return 128
	}

	return intVar
}

func GetRedisPoolTimeout() int {
	intVar, err := strconv.Atoi(getEnv("REDIS_POOL_TIMEOUT", "10"))
	if err != nil {
		return 10
	}

	return intVar
}

func GetRedisMinIdleConn() int {
	intVar, err := strconv.Atoi(getEnv("REDIS_MIN_IDLE_CONN", "4"))
	if err != nil {
		return 4
	}

	return intVar
}
func GetRedisDefaultCacheTimeOut() int {
	intVar, err := strconv.Atoi(getEnv("CACHE_DEFAULT_TIMEOUT", "5"))
	if err != nil {
		return 5
	}

	return intVar
}

func GetRabbitPort() int {
	intVar, err := strconv.Atoi(getEnv("RABBIT_PORT", "5672"))
	if err != nil {
		return 5672
	}

	return intVar
}

func UseRabbitMQ() bool {
	res, err := strconv.ParseBool(getEnv("USE_RABBIT", "true"))
	if err != nil {
		return true
	}

	return res
}
