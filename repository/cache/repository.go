package cache

import (
	"context"
	"github.com/avtara/travair-api/businesses/cache"
	"github.com/go-redis/redis/v8"

	"time"
)

type repoCache struct {
	Conn *redis.Client
}

func NewRepoCache(rc *redis.Client) cache.Repository {
	return &repoCache{
		Conn: rc,
	}
}

func (rc *repoCache) Get(ctx context.Context, key string) (string, error) {
	val, err := rc.Conn.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (rc *repoCache) Set(ctx context.Context, key, value string, timeOut time.Duration) (string, error) {
	set, err := rc.Conn.SetEX(ctx, key, value, timeOut*time.Second).Result()
	if err != nil {
		return "", err
	}
	return set, nil
}
