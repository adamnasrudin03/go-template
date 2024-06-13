package repository

import (
	"context"
)

func (r *cacheRepo) DelCache(ctx context.Context, key string) error {
	var (
		opName = "CacheRepository-DelCache"
		err    error
	)

	err = r.Cache.Del(key)
	if err != nil {
		r.Logger.Errorf("%v error: %v ", opName, err)
		return err
	}

	return nil
}
