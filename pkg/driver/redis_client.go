package driver

import (
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type RedisClient interface {
	HGet(key, field string) ([]byte, error)
	HSet(key, field string, value interface{}, expDur time.Duration) error
	HDel(key, field string) error
	Del(key string) error
	Set(key string, value interface{}, expDur time.Duration) error
	Get(key string) (string, error)
	SetNX(key string, value string, duration time.Duration) (bool, error)
	Keys(key string) ([]string, error)
}

type redisCtx struct {
	redisClient redis.Cmdable
}

func NewRedis(redisClient redis.Cmdable) RedisClient {
	return &redisCtx{
		redisClient: redisClient,
	}
}

// HGet get
func (c *redisCtx) HGet(key, field string) ([]byte, error) {
	data, err := c.redisClient.HGet(key, field).Result()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return []byte(data), nil
}

// HSet set
func (c *redisCtx) HSet(key, field string, value interface{}, expDur time.Duration) error {
	payload, err := json.Marshal(value)
	if err != nil {
		log.Print(err)
		return err
	}

	err = c.redisClient.HSet(key, field, payload).Err()
	if err != nil {
		log.Print(err)
		return err
	}

	if expDur > 0 {
		err = c.redisClient.Expire(key, expDur).Err()
		if err != nil {
			log.Print(err)
			return err
		}
	}

	return nil
}

// HDel delete
func (c *redisCtx) HDel(key, field string) error {
	err := c.redisClient.HDel(key, field).Err()
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

// SetNX set
func (c *redisCtx) SetNX(key string, value string, exp time.Duration) (bool, error) {
	return c.redisClient.SetNX(key, value, exp).Result()
}
