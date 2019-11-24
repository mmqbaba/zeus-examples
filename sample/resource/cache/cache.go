package cache

import (
	"context"

	"github.com/go-redis/redis"
	zeusctx "gitlab.dg.com/BackEnd/jichuchanpin/tif/zeus/context"
)

func GetUser(ctx context.Context, id string) (info string, err error) {
	rdc, err := zeusctx.ExtractRedis(ctx)
	if err != nil {
		return
	}
	result := rdc.Get(id)
	if err = result.Err(); err != nil && err != redis.Nil {
		return
	}
	err = nil
	info = result.Val()
	return
}

func SetUser(ctx context.Context, id string, info string) (err error) {
	rdc, err := zeusctx.ExtractRedis(ctx)
	if err != nil {
		return
	}
	err = rdc.Set(id, info, 0).Err()
	return
}
