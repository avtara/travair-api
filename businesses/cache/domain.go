package cache

import (
	"context"
	"time"
)

type Repository interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string, time time.Duration) (string, error)
}