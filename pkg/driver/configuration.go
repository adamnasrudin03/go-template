package driver

import (
	"fmt"
	"sync"
	"time"

	"github.com/adamnasrudin03/go-template/app/configs"
	"github.com/go-redis/redis"
	amqp "github.com/rabbitmq/amqp091-go"
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

		redisClient = NewRedis(redisConn)
	})

	return redisClient
}

func ConnectMQ(config *configs.Configs) (*amqp.Connection, *amqp.Channel) {
	addrs := fmt.Sprintf("amqp://%s:%s@%s:%d", config.RabbitMQ.Username, config.RabbitMQ.Password, config.RabbitMQ.Host, config.RabbitMQ.Port)

	conn, err := amqp.Dial(addrs)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to RabbitMQ: %s", err))
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(fmt.Sprintf("Failed to open a channel: %s", err))
	}
	return conn, ch
}

func CloseMQ(conn *amqp.Connection, channel *amqp.Channel) {

	defer conn.Close()    //rabbit mq close
	defer channel.Close() //rabbit mq channel close)
}
