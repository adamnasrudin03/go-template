package driver

import (
	"log"
	"time"

	"github.com/adamnasrudin03/go-template/pkg/helpers"
)

func (c *redisCtx) Del(key string) error {
	err := c.redisClient.Del(key).Err()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (c *redisCtx) Get(key string) (string, error) {
	data, err := c.redisClient.Get(key).Result()
	if err != nil {
		log.Print(err)
		return "", err
	}
	return data, nil
}

func (c *redisCtx) Set(key string, value interface{}, expDur time.Duration) error {
	payload, err := helpers.SafeJsonMarshal(value)
	if err != nil {
		log.Print(err)
		return err
	}

	err = c.redisClient.Set(key, payload, expDur).Err()
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (c *redisCtx) Keys(key string) ([]string, error) {
	data, err := c.redisClient.Keys(key).Result()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return data, nil
}
