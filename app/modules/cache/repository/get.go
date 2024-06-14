package repository

import (
	"context"
	"encoding/json"
)

func (r *cacheRepo) GetCache(ctx context.Context, key string, res interface{}) bool {
	var (
		opName = "CacheRepository-GetCache"
		err    error
	)

	data, err := r.Cache.Get(key)
	if err != nil {
		r.Logger.Errorf("%v error: %v ", opName, err)
		return false
	}

	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		r.Logger.Errorf("%v Unmarshal error: %v ", opName, err)
		return false
	}

	return true
}
