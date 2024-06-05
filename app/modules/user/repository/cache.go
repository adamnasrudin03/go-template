package repository

import (
	"context"
	"encoding/json"
	"time"
)

func (r *userRepo) CreateCache(ctx context.Context, key string, data interface{}) {
	var (
		opName = "UserRepository-CreateCache"
		err    error
	)

	err = r.Cache.Set(key, data, time.Duration(config.Redis.DefaultCacheTimeOut)*time.Minute)
	if err != nil {
		r.Logger.Errorf("%v error: %v \n", opName, err)
		return
	}
}

func (r *userRepo) GetCache(ctx context.Context, key string, res interface{}) {
	var (
		opName = "UserRepository-GetCache"
		err    error
	)

	data, err := r.Cache.Get(key)
	if err != nil {
		r.Logger.Errorf("%v error: %v \n", opName, err)
		return
	}

	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		r.Logger.Errorf("%v Unmarshal error: %v \n", opName, err)
		return
	}
}

func (r *userRepo) DelCache(ctx context.Context, key string) error {
	var (
		opName = "UserRepository-DelCache"
		err    error
	)

	err = r.Cache.Del(key)
	if err != nil {
		r.Logger.Errorf("%v error: %v \n", opName, err)
		return err
	}

	return nil
}
