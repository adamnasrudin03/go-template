package repository

import (
	"context"
	"encoding/json"
	"time"
)

func (r *logRepo) CreateCache(ctx context.Context, key string, data interface{}) {
	var (
		opName = "LogRepository-CreateCache"
		err    error
	)

	err = r.Cache.Set(key, data, time.Duration(r.Cfg.Redis.DefaultCacheTimeOut)*time.Minute)
	if err != nil {
		r.Logger.Errorf("%v error: %v ", opName, err)
		return
	}
}

func (r *logRepo) GetCache(ctx context.Context, key string, res interface{}) {
	var (
		opName = "LogRepository-GetCache"
		err    error
	)

	data, err := r.Cache.Get(key)
	if err != nil {
		r.Logger.Errorf("%v error: %v ", opName, err)
		return
	}

	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		r.Logger.Errorf("%v Unmarshal error: %v ", opName, err)
		return
	}
}

func (r *logRepo) DelCache(ctx context.Context, key string) error {
	var (
		opName = "LogRepository-DelCache"
		err    error
	)

	err = r.Cache.Del(key)
	if err != nil {
		r.Logger.Errorf("%v error: %v ", opName, err)
		return err
	}

	return nil
}
