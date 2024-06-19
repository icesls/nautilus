// +----------------------------------------------------------------------
// | nautilus [ redis implements cache ]
// +----------------------------------------------------------------------
// | Copyright (c) 2013~2024 https://www.secdos.com All rights reserved.
// +----------------------------------------------------------------------
// | Licensed ( http://www.apache.org/licenses/LICENSE-2.0 )
// +----------------------------------------------------------------------
// | Author: ︶ㄣ逍遥楓 <admin@secdos.com>
// +----------------------------------------------------------------------

package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisCache struct {
	client *redis.Client
	prefix string
}

// NewRedisCache 初始配置信息
func NewRedisCache(address, username, password, prefix string, db int) *RedisCache {
	rc := &RedisCache{}
	rc.client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})
	rc.prefix = prefix + ":cache:"

	return rc
}

func (r *RedisCache) Set(key string, value string, expire time.Duration) error {
	return r.client.Set(context.Background(), r.prefix+key, value, expire).Err()
}

func (r *RedisCache) Get(key string) string {
	return r.client.Get(context.Background(), r.prefix+key).Val()
}

func (r *RedisCache) Has(key string) bool {
	_, err := r.client.Exists(context.Background(), r.prefix+key).Result()
	if err != nil {
		return false
	}

	return true
}

func (r *RedisCache) Del(key string) error {
	return r.client.Del(context.Background(), r.prefix+key).Err()
}
