package driver

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/go-redis/redis"
)

var (
	redisClientOnce sync.Once
	redisClient     RedisClient
)

func Redis(config *configs.Configs) RedisClient {
	redisClientOnce.Do(func() {

		addrs := fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)
		redisConn := redis.NewClient(&redis.Options{
			Addr:         addrs,
			Password:     config.Redis.Password,
			DB:           config.Redis.Database,
			PoolSize:     config.Redis.PoolSize,
			PoolTimeout:  time.Duration(config.Redis.PoolTimeout) * time.Second,
			MinIdleConns: config.Redis.MinIdleConn,
		})

		_, err := redisConn.Ping().Result()
		if err != nil {
			log.Fatal(err)
		}

		redisClient = NewRedis(redisConn)
	})

	return redisClient
}
