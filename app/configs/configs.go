package configs

import (
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

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
				Name:         getEnv("APP_NAME", "go-template"),
				Env:          getEnv("APP_ENV", "dev"),
				Port:         getEnv("APP_PORT", "8000"),
				SecretKey:    getEnv("JWT_SECRET", "MySecretKey"),
				ExpiredToken: GetExpiredToken(),
				UseRabbitMQ:  UseRabbitMQ(),
				OtpLength:    GetOtpLength(),
				OtpExpired:   GetOtpExpired(),
			},
			DB: DbConfig{
				Host:        getEnv("DB_HOST", "127.0.0.1"),
				Port:        getEnv("DB_PORT", "5432"),
				DbName:      getEnv("DB_NAME", "movie_festival_db"),
				Username:    getEnv("DB_USER", "postgres"),
				Password:    getEnv("DB_PASS", ""),
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
				Password: getEnv("RABBIT_PASSWORD", ""),
			},
			Email: EmailConfig{
				Host:         getEnv("MAIL_HOST", "smtp.gmail.com"),
				Port:         GetEmailPort(),
				AuthEmail:    getEnv("MAIL_AUTH_EMAIL", ""),
				AuthPassword: getEnv("MAIL_AUTH_PASSWORD", ""),
				Sender:       getEnv("MAIL_SENDER", ""),
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
func GetRedisDefaultCacheTimeOut() time.Duration {
	intVar, err := strconv.Atoi(getEnv("CACHE_DEFAULT_TIMEOUT", "5"))
	if err != nil {
		return 5 * time.Minute
	}

	return time.Duration(intVar) * time.Minute
}

func GetRabbitPort() int {
	intVar, err := strconv.Atoi(getEnv("RABBIT_PORT", "5672"))
	if err != nil {
		return 5672
	}

	return intVar
}

func UseRabbitMQ() bool {
	res, err := strconv.ParseBool(getEnv("USE_RABBIT", "false"))
	if err != nil {
		return true
	}

	return res
}

func GetEmailPort() int {
	intVar, err := strconv.Atoi(getEnv("MAIL_PORT", "587"))
	if err != nil {
		return 587
	}

	return intVar
}

func GetOtpLength() int {
	intVar, err := strconv.Atoi(getEnv("OTP_LENGTH", "6"))
	if err != nil {
		return 6
	}

	return intVar
}

func GetOtpExpired() time.Duration {
	intVar, err := strconv.Atoi(getEnv("OTP_EXPIRED", "1"))
	if err != nil {
		return 1 * time.Minute
	}

	return time.Duration(intVar) * time.Minute
}
